// Package log is the logging library used by IPFS
// (https://github.com/ipfs/go-ipfs). It uses a modified version of
// https://godoc.org/github.com/whyrusleeping/go-logging .
package log

import (
	"bytes"
	"context"
	"encoding/json"
	"path"
	"runtime"
	"time"

	writer "github.com/ipfs/go-log/writer"
	"github.com/whyrusleeping/go-logging"

	opentrace "github.com/opentracing/opentracing-go"
	otExt "github.com/opentracing/opentracing-go/ext"
)

var log = Logger("eventlog")

// StandardLogger provides API compatibility with standard printf loggers
// eg. go-logging
type StandardLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	// Deprecated use Warn
	Warning(args ...interface{})
	// Deprecated use Warnf
	Warningf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

// EventLogger extends the StandardLogger interface to allow for log items
// containing structured metadata
type EventLogger interface {
	StandardLogger

	// Event merges structured data from the provided inputs into a single
	// machine-readable log event.
	//
	// If the context contains metadata, a copy of this is used as the base
	// metadata accumulator.
	//
	// If one or more loggable objects are provided, these are deep-merged into base blob.
	//
	// Next, the event name is added to the blob under the key "event". If
	// the key "event" already exists, it will be over-written.
	//
	// Finally the timestamp and package name are added to the accumulator and
	// the metadata is logged.
	// DEPRECATED
	// Deprecated: Stop using go-log for event logging
	Event(ctx context.Context, event string, m ...Loggable)

	// DEPRECATED
	// Deprecated: Stop using go-log for event logging
	EventBegin(ctx context.Context, event string, m ...Loggable) *EventInProgress

	// Start starts an opentracing span with `name`, using
	// any Span found within `ctx` as a ChildOfRef. If no such parent could be
	// found, Start creates a root (parentless) Span.
	//
	// The return value is a context.Context object built around the
	// returned Span.
	//
	// Example usage:
	//
	//    SomeFunction(ctx context.Context, ...) {
	//        ctx := log.Start(ctx, "SomeFunction")
	//        defer log.Finish(ctx)
	//        ...
	//    }
	// Deprecated: Stop using go-log for event logging
	Start(ctx context.Context, name string) context.Context

	// StartFromParentState starts an opentracing span with `name`, using
	// any Span found within `ctx` as a ChildOfRef. If no such parent could be
	// found, StartSpanFromParentState creates a root (parentless) Span.
	//
	// StartFromParentState will attempt to deserialize a SpanContext from `parent`,
	// using any Span found within to continue the trace
	//
	// The return value is a context.Context object built around the
	// returned Span.
	//
	// An error is returned when `parent` cannot be deserialized to a SpanContext
	//
	// Example usage:
	//
	//    SomeFunction(ctx context.Context, bParent []byte) {
	//        ctx := log.StartFromParentState(ctx, "SomeFunction", bParent)
	//        defer log.Finish(ctx)
	//        ...
	//    }
	// Deprecated: Stop using go-log for event logging
	StartFromParentState(ctx context.Context, name string, parent []byte) (context.Context, error)

	// Finish completes the span associated with `ctx`.
	//
	// Finish() must be the last call made to any span instance, and to do
	// otherwise leads to undefined behavior.
	// Finish will do its best to notify (log) when used in correctly
	//		.e.g called twice, or called on a spanless `ctx`
	// Deprecated: Stop using go-log for event logging
	Finish(ctx context.Context)

	// FinishWithErr completes the span associated with `ctx` and also calls
	// SetErr if `err` is non-nil
	//
	// FinishWithErr() must be the last call made to any span instance, and to do
	// otherwise leads to undefined behavior.
	// FinishWithErr will do its best to notify (log) when used in correctly
	//		.e.g called twice, or called on a spanless `ctx`
	// Deprecated: Stop using go-log for event logging
	FinishWithErr(ctx context.Context, err error)

	// SetErr tags the span associated with `ctx` to reflect an error occured, and
	// logs the value `err` under key `error`.
	// Deprecated: Stop using go-log for event logging
	SetErr(ctx context.Context, err error)

	// LogKV records key:value logging data about an event stored in `ctx`
	// Eexample:
	//    log.LogKV(
	//        "error", "resolve failure",
	//        "type", "cache timeout",
	//        "waited.millis", 1500)
	// Deprecated: Stop using go-log for event logging
	LogKV(ctx context.Context, alternatingKeyValues ...interface{})

	// SetTag tags key `k` and value `v` on the span associated with `ctx`
	// Deprecated: Stop using go-log for event logging
	SetTag(ctx context.Context, key string, value interface{})

	// SetTags tags keys from the `tags` maps on the span associated with `ctx`
	// Example:
	//    log.SetTags(ctx, map[string]{
	//		"type": bizStruct,
	//      "request": req,
	//		})
	// Deprecated: Stop using go-log for event logging
	SetTags(ctx context.Context, tags map[string]interface{})

	// SerializeContext takes the SpanContext instance stored in `ctx` and Seralizes
	// it to bytes. An error is returned if the `ctx` cannot be serialized to
	// a bytes array
	// Deprecated: Stop using go-log for event logging
	SerializeContext(ctx context.Context) ([]byte, error)
}

// Logger retrieves an event logger by name
func Logger(system string) EventLogger {

	// TODO if we would like to adjust log levels at run-time. Store this event
	// logger in a map (just like the util.Logger impl)
	if len(system) == 0 {
		setuplog := getLogger("setup-logger")
		setuplog.Warning("Missing name parameter")
		system = "undefined"
	}

	logger := getLogger(system)

	return &eventLogger{system: system, Logger: *logger}
}

// eventLogger implements the EventLogger and wraps a go-logging Logger
type eventLogger struct {
	logging.Logger

	system string
	// TODO add log-level
}

func (el *eventLogger) Warn(args ...interface{}) {
	el.Warning(args...)
}
func (el *eventLogger) Warnf(format string, args ...interface{}) {
	el.Warningf(format, args...)
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) Start(ctx context.Context, operationName string) context.Context {
	span, ctx := opentrace.StartSpanFromContext(ctx, operationName)
	span.SetTag("system", el.system)
	return ctx
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) StartFromParentState(ctx context.Context, operationName string, parent []byte) (context.Context, error) {
	sc, err := deserializeContext(parent)
	if err != nil {
		return nil, err
	}

	//TODO RPCServerOption is probably not the best tag, as this is likely from a peer
	span, ctx := opentrace.StartSpanFromContext(ctx, operationName, otExt.RPCServerOption(sc))
	span.SetTag("system", el.system)
	return ctx, nil
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) SerializeContext(ctx context.Context) ([]byte, error) {
	gTracer := opentrace.GlobalTracer()
	b := make([]byte, 0)
	carrier := bytes.NewBuffer(b)
	span := opentrace.SpanFromContext(ctx)
	if err := gTracer.Inject(span.Context(), opentrace.Binary, carrier); err != nil {
		return nil, err
	}
	return carrier.Bytes(), nil
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) LogKV(ctx context.Context, alternatingKeyValues ...interface{}) {
	span := opentrace.SpanFromContext(ctx)
	if span == nil {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("LogKV with no Span in context called on %s:%d", path.Base(file), line)
		return
	}
	span.LogKV(alternatingKeyValues...)
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) SetTag(ctx context.Context, k string, v interface{}) {
	span := opentrace.SpanFromContext(ctx)
	if span == nil {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("SetTag with no Span in context called on %s:%d", path.Base(file), line)
		return
	}
	span.SetTag(k, v)
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) SetTags(ctx context.Context, tags map[string]interface{}) {
	span := opentrace.SpanFromContext(ctx)
	if span == nil {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("SetTags with no Span in context called on %s:%d", path.Base(file), line)
		return
	}
	for k, v := range tags {
		span.SetTag(k, v)
	}
}

func (el *eventLogger) setErr(ctx context.Context, err error, skip int) {
	span := opentrace.SpanFromContext(ctx)
	if span == nil {
		_, file, line, _ := runtime.Caller(skip)
		log.Errorf("SetErr with no Span in context called on %s:%d", path.Base(file), line)
		return
	}
	if err == nil {
		return
	}

	otExt.Error.Set(span, true)
	span.LogKV("error", err.Error())
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) SetErr(ctx context.Context, err error) {
	el.setErr(ctx, err, 1)
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) Finish(ctx context.Context) {
	span := opentrace.SpanFromContext(ctx)
	if span == nil {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("Finish with no Span in context called on %s:%d", path.Base(file), line)
		return
	}
	span.Finish()
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) FinishWithErr(ctx context.Context, err error) {
	el.setErr(ctx, err, 2)
	el.Finish(ctx)
}

func deserializeContext(bCtx []byte) (opentrace.SpanContext, error) {
	gTracer := opentrace.GlobalTracer()
	carrier := bytes.NewReader(bCtx)
	spanContext, err := gTracer.Extract(opentrace.Binary, carrier)
	if err != nil {
		log.Warning("Failed to deserialize context %s", err)
		return nil, err
	}
	return spanContext, nil
}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) EventBegin(ctx context.Context, event string, metadata ...Loggable) *EventInProgress {
	ctx = el.Start(ctx, event)

	for _, m := range metadata {
		for l, v := range m.Loggable() {
			el.LogKV(ctx, l, v)
		}
	}

	eip := &EventInProgress{}
	eip.doneFunc = func(additional []Loggable) {
		// anything added during the operation
		// e.g. deprecated methods event.Append(...) or event.SetError(...)
		for _, m := range eip.loggables {
			for l, v := range m.Loggable() {
				el.LogKV(ctx, l, v)
			}
		}
		el.Finish(ctx)
	}
	return eip
}

type activeEventKeyType struct{}

var activeEventKey = activeEventKeyType{}

// Deprecated: Stop using go-log for event logging
func (el *eventLogger) Event(ctx context.Context, event string, metadata ...Loggable) {

	// short circuit if theres nothing to write to
	if !writer.WriterGroup.Active() {
		return
	}

	// Collect loggables for later logging
	var loggables []Loggable

	// get any existing metadata from the context
	existing, err := MetadataFromContext(ctx)
	if err != nil {
		existing = Metadata{}
	}
	loggables = append(loggables, existing)

	for _, datum := range metadata {
		loggables = append(loggables, datum)
	}

	e := entry{
		loggables: loggables,
		system:    el.system,
		event:     event,
	}

	accum := Metadata{}
	for _, loggable := range e.loggables {
		accum = DeepMerge(accum, loggable.Loggable())
	}

	// apply final attributes to reserved keys
	// TODO accum["level"] = level
	accum["event"] = e.event
	accum["system"] = e.system
	accum["time"] = FormatRFC3339(time.Now())

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(accum)
	if err != nil {
		el.Errorf("ERROR FORMATTING EVENT ENTRY: %s", err)
		return
	}

	writer.WriterGroup.Write(buf.Bytes())
}

// DEPRECATED
// EventInProgress represent and event which is happening
// Deprecated: Stop using go-log for event logging
type EventInProgress struct {
	loggables []Loggable
	doneFunc  func([]Loggable)
}

// DEPRECATED use `LogKV` or `SetTag`
// Append adds loggables to be included in the call to Done
func (eip *EventInProgress) Append(l Loggable) {
	eip.loggables = append(eip.loggables, l)
}

// DEPRECATED use `SetError(ctx, error)`
// SetError includes the provided error
func (eip *EventInProgress) SetError(err error) {
	eip.loggables = append(eip.loggables, LoggableMap{
		"error": err.Error(),
	})
}

// Done creates a new Event entry that includes the duration and appended
// loggables.
// Deprecated: Stop using go-log for event logging
func (eip *EventInProgress) Done() {
	eip.doneFunc(eip.loggables) // create final event with extra data
}

// DEPRECATED use `FinishWithErr`
// DoneWithErr creates a new Event entry that includes the duration and appended
// loggables. DoneWithErr accepts an error, if err is non-nil, it is set on
// the EventInProgress. Otherwise the logic is the same as the `Done()` method
func (eip *EventInProgress) DoneWithErr(err error) {
	if err != nil {
		eip.SetError(err)
	}
	eip.doneFunc(eip.loggables)
}

// Close is an alias for done
// Deprecated: Stop using go-log for event logging
func (eip *EventInProgress) Close() error {
	eip.Done()
	return nil
}

// FormatRFC3339 returns the given time in UTC with RFC3999Nano format.
func FormatRFC3339(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

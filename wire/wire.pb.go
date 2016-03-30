// Code generated by protoc-gen-gogo.
// source: wire.proto
// DO NOT EDIT!

/*
	Package wire is a generated protocol buffer package.

	It is generated from these files:
		wire.proto

	It has these top-level messages:
		TracerState
*/
package wire

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type TracerState struct {
	TraceId      int64             `protobuf:"fixed64,1,opt,name=trace_id,proto3" json:"trace_id,omitempty"`
	SpanId       int64             `protobuf:"fixed64,2,opt,name=span_id,proto3" json:"span_id,omitempty"`
	Sampled      bool              `protobuf:"varint,3,opt,name=sampled,proto3" json:"sampled,omitempty"`
	BaggageItems map[string]string `protobuf:"bytes,4,rep,name=baggage_items" json:"baggage_items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *TracerState) Reset()                    { *m = TracerState{} }
func (m *TracerState) String() string            { return proto.CompactTextString(m) }
func (*TracerState) ProtoMessage()               {}
func (*TracerState) Descriptor() ([]byte, []int) { return fileDescriptorWire, []int{0} }

func (m *TracerState) GetBaggageItems() map[string]string {
	if m != nil {
		return m.BaggageItems
	}
	return nil
}

func init() {
	proto.RegisterType((*TracerState)(nil), "basictracer_go.wire.TracerState")
}
func (m *TracerState) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TracerState) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TraceId != 0 {
		data[i] = 0x9
		i++
		i = encodeFixed64Wire(data, i, uint64(m.TraceId))
	}
	if m.SpanId != 0 {
		data[i] = 0x11
		i++
		i = encodeFixed64Wire(data, i, uint64(m.SpanId))
	}
	if m.Sampled {
		data[i] = 0x18
		i++
		if m.Sampled {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.BaggageItems) > 0 {
		for k, _ := range m.BaggageItems {
			data[i] = 0x22
			i++
			v := m.BaggageItems[k]
			mapSize := 1 + len(k) + sovWire(uint64(len(k))) + 1 + len(v) + sovWire(uint64(len(v)))
			i = encodeVarintWire(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintWire(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintWire(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	return i, nil
}

func encodeFixed64Wire(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Wire(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWire(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *TracerState) Size() (n int) {
	var l int
	_ = l
	if m.TraceId != 0 {
		n += 9
	}
	if m.SpanId != 0 {
		n += 9
	}
	if m.Sampled {
		n += 2
	}
	if len(m.BaggageItems) > 0 {
		for k, v := range m.BaggageItems {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWire(uint64(len(k))) + 1 + len(v) + sovWire(uint64(len(v)))
			n += mapEntrySize + 1 + sovWire(uint64(mapEntrySize))
		}
	}
	return n
}

func sovWire(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWire(x uint64) (n int) {
	return sovWire(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TracerState) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TracerState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TracerState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			m.TraceId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.TraceId = int64(data[iNdEx-8])
			m.TraceId |= int64(data[iNdEx-7]) << 8
			m.TraceId |= int64(data[iNdEx-6]) << 16
			m.TraceId |= int64(data[iNdEx-5]) << 24
			m.TraceId |= int64(data[iNdEx-4]) << 32
			m.TraceId |= int64(data[iNdEx-3]) << 40
			m.TraceId |= int64(data[iNdEx-2]) << 48
			m.TraceId |= int64(data[iNdEx-1]) << 56
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.SpanId = int64(data[iNdEx-8])
			m.SpanId |= int64(data[iNdEx-7]) << 8
			m.SpanId |= int64(data[iNdEx-6]) << 16
			m.SpanId |= int64(data[iNdEx-5]) << 24
			m.SpanId |= int64(data[iNdEx-4]) << 32
			m.SpanId |= int64(data[iNdEx-3]) << 40
			m.SpanId |= int64(data[iNdEx-2]) << 48
			m.SpanId |= int64(data[iNdEx-1]) << 56
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sampled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Sampled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaggageItems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthWire
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthWire
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.BaggageItems == nil {
				m.BaggageItems = make(map[string]string)
			}
			m.BaggageItems[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWire
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWire(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWire
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWire
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWire
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWire
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWire
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWire(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWire = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWire   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorWire = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xcf, 0x2c, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4e, 0x4a, 0x2c, 0xce, 0x4c, 0x2e, 0x29, 0x4a,
	0x4c, 0x4e, 0x2d, 0x8a, 0x4f, 0xcf, 0xd7, 0x03, 0x49, 0x29, 0x1d, 0x64, 0xe4, 0xe2, 0x0e, 0x01,
	0x0b, 0x05, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x09, 0x70, 0x71, 0x80, 0x55, 0xc4, 0x67, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x08, 0x08, 0xf1, 0x73, 0xb1, 0x17, 0x17, 0x24, 0xe6, 0x81, 0x04, 0x98,
	0xe0, 0x02, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0x29, 0x12, 0xcc, 0x40, 0x01, 0x0e, 0x21, 0x4f, 0x2e,
	0xde, 0xa4, 0xc4, 0xf4, 0xf4, 0xc4, 0x74, 0xa0, 0xae, 0x92, 0xd4, 0xdc, 0x62, 0x09, 0x16, 0x05,
	0x66, 0x0d, 0x6e, 0x23, 0x23, 0x3d, 0x2c, 0x16, 0xea, 0x21, 0x59, 0xa6, 0xe7, 0x04, 0xd1, 0xe5,
	0x09, 0xd2, 0xe4, 0x9a, 0x57, 0x52, 0x54, 0x29, 0x65, 0xcc, 0x25, 0x88, 0x21, 0x28, 0xc4, 0xcd,
	0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0x76, 0x0e, 0xa7, 0x10, 0x2f, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69,
	0x2a, 0xd8, 0x31, 0x9c, 0x56, 0x4c, 0x16, 0x8c, 0x4e, 0x62, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x00,
	0xe2, 0x07, 0x40, 0x3c, 0xe1, 0xb1, 0x1c, 0x43, 0x14, 0x0b, 0xc8, 0xaa, 0x24, 0x36, 0xb0, 0xbf,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x09, 0x69, 0x48, 0x08, 0x05, 0x01, 0x00, 0x00,
}
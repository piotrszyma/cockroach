// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/log/eventpb/events.proto

package eventpb

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// CommonEventDetails contains the fields common to all events.
type CommonEventDetails struct {
	// The timestamp of the event. Expressed as nanoseconds since
	// the Unix epoch.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:",omitempty"`
	// The type of the event.
	EventType string `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:",omitempty" redact:"nonsensitive"`
}

func (m *CommonEventDetails) Reset()         { *m = CommonEventDetails{} }
func (m *CommonEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonEventDetails) ProtoMessage()    {}
func (*CommonEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_c8e68b80e3722df3, []int{0}
}
func (m *CommonEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *CommonEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonEventDetails.Merge(dst, src)
}
func (m *CommonEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonEventDetails proto.InternalMessageInfo

// CommonSQLEventDetails contains the fields common to all
// SQL events.
type CommonSQLEventDetails struct {
	// A normalized copy of the SQL statement that triggered the event.
	Statement string `protobuf:"bytes,1,opt,name=statement,proto3" json:",omitempty"`
	// The user account that triggered the event.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:",omitempty"`
	// The primary object descriptor affected by the operation. Set to zero for operations
	// that don't affect descriptors.
	DescriptorID uint32 `protobuf:"varint,3,opt,name=descriptor_id,json=descriptorId,proto3" json:",omitempty"`
}

func (m *CommonSQLEventDetails) Reset()         { *m = CommonSQLEventDetails{} }
func (m *CommonSQLEventDetails) String() string { return proto.CompactTextString(m) }
func (*CommonSQLEventDetails) ProtoMessage()    {}
func (*CommonSQLEventDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_c8e68b80e3722df3, []int{1}
}
func (m *CommonSQLEventDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonSQLEventDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *CommonSQLEventDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonSQLEventDetails.Merge(dst, src)
}
func (m *CommonSQLEventDetails) XXX_Size() int {
	return m.Size()
}
func (m *CommonSQLEventDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonSQLEventDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CommonSQLEventDetails proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CommonEventDetails)(nil), "cockroach.util.log.eventpb.CommonEventDetails")
	proto.RegisterType((*CommonSQLEventDetails)(nil), "cockroach.util.log.eventpb.CommonSQLEventDetails")
}
func (m *CommonEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonEventDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.EventType) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EventType)))
		i += copy(dAtA[i:], m.EventType)
	}
	return i, nil
}

func (m *CommonSQLEventDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonSQLEventDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Statement) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Statement)))
		i += copy(dAtA[i:], m.Statement)
	}
	if len(m.User) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	if m.DescriptorID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.DescriptorID))
	}
	return i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CommonEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovEvents(uint64(m.Timestamp))
	}
	l = len(m.EventType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *CommonSQLEventDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Statement)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.DescriptorID != 0 {
		n += 1 + sovEvents(uint64(m.DescriptorID))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommonEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommonEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func (m *CommonSQLEventDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommonSQLEventDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonSQLEventDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Statement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Statement = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DescriptorID", wireType)
			}
			m.DescriptorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DescriptorID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvents
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipEvents(dAtA[start:])
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
	ErrInvalidLengthEvents = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("util/log/eventpb/events.proto", fileDescriptor_events_c8e68b80e3722df3)
}

var fileDescriptor_events_c8e68b80e3722df3 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x40, 0x33, 0xcf, 0xc7, 0x7b, 0x64, 0xd0, 0xb7, 0x08, 0x4f, 0x10, 0xa1, 0x13, 0xc9, 0xa6,
	0x16, 0x24, 0x59, 0x74, 0xd7, 0x65, 0xb4, 0x05, 0xa1, 0x9b, 0xda, 0xae, 0xba, 0x91, 0x98, 0xdc,
	0xa6, 0x43, 0x93, 0xdc, 0x90, 0xb9, 0x0a, 0xfe, 0x42, 0x57, 0xfd, 0x92, 0x7e, 0x87, 0x4b, 0x97,
	0xae, 0xa4, 0x8d, 0xbb, 0x2e, 0xfb, 0x05, 0xc5, 0x89, 0x28, 0xb5, 0xab, 0x99, 0xe1, 0x9e, 0x33,
	0x1c, 0x66, 0xf8, 0xc9, 0x94, 0x64, 0xe2, 0x25, 0x18, 0x7b, 0x30, 0x83, 0x8c, 0xf2, 0x49, 0xb5,
	0x2a, 0x37, 0x2f, 0x90, 0xd0, 0x6a, 0x87, 0x18, 0x3e, 0x15, 0x18, 0x84, 0x8f, 0xee, 0x16, 0x74,
	0x13, 0x8c, 0xdd, 0x1d, 0xd8, 0xfe, 0x1f, 0x63, 0x8c, 0x1a, 0xf3, 0xb6, 0xbb, 0xca, 0x68, 0xdb,
	0x31, 0x62, 0x9c, 0x80, 0xa7, 0x4f, 0x93, 0xe9, 0x83, 0x47, 0x32, 0x05, 0x45, 0x41, 0x9a, 0x57,
	0x80, 0xf3, 0xcc, 0xb8, 0xd5, 0xc7, 0x34, 0xc5, 0xec, 0x72, 0x7b, 0xd1, 0x00, 0x28, 0x90, 0x89,
	0xb2, 0x7a, 0xdc, 0xdc, 0x93, 0x2d, 0xd6, 0x61, 0xdd, 0x9a, 0xff, 0xef, 0x63, 0x6d, 0xf3, 0x1e,
	0xa6, 0x92, 0x20, 0xcd, 0x69, 0x3e, 0x3a, 0x00, 0xd6, 0x15, 0xe7, 0x3a, 0x63, 0x4c, 0xf3, 0x1c,
	0x5a, 0xbf, 0x3a, 0xac, 0x6b, 0xfa, 0xa7, 0xdf, 0xf1, 0xcf, 0xb5, 0xdd, 0x2c, 0x20, 0x0a, 0x42,
	0xba, 0x70, 0x32, 0xcc, 0x14, 0x64, 0x4a, 0x92, 0x9c, 0x81, 0x33, 0x32, 0xb5, 0x7a, 0x37, 0xcf,
	0xc1, 0x79, 0x65, 0xbc, 0x59, 0xc5, 0xdc, 0xde, 0x5c, 0x1f, 0xf7, 0x28, 0x0a, 0x08, 0x52, 0xc8,
	0x48, 0xf7, 0x98, 0x3f, 0x7b, 0xf6, 0x80, 0xe5, 0xf0, 0xdf, 0x53, 0x05, 0xc5, 0xae, 0xe4, 0x18,
	0xd4, 0x33, 0xab, 0xcf, 0x1b, 0x11, 0xa8, 0xb0, 0x90, 0x39, 0x61, 0x31, 0x96, 0x51, 0xab, 0xd6,
	0x61, 0xdd, 0x86, 0x2f, 0xca, 0xb5, 0x5d, 0x1f, 0xec, 0x07, 0xc3, 0xc1, 0x91, 0x5c, 0x3f, 0x48,
	0xc3, 0xc8, 0x3f, 0x5b, 0xbc, 0x0b, 0x63, 0x51, 0x0a, 0xb6, 0x2c, 0x05, 0x5b, 0x95, 0x82, 0xbd,
	0x95, 0x82, 0xbd, 0x6c, 0x84, 0xb1, 0xdc, 0x08, 0x63, 0xb5, 0x11, 0xc6, 0xfd, 0xdf, 0xdd, 0xff,
	0x4c, 0xfe, 0xe8, 0xf7, 0x3e, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x21, 0xf3, 0xc8, 0xa4, 0xe3,
	0x01, 0x00, 0x00,
}
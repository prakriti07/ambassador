// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/listener/proxy_protocol/v2/proxy_protocol.proto

package envoy_config_filter_listener_proxy_protocol_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProxyProtocol struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyProtocol) Reset()         { *m = ProxyProtocol{} }
func (m *ProxyProtocol) String() string { return proto.CompactTextString(m) }
func (*ProxyProtocol) ProtoMessage()    {}
func (*ProxyProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_614eea5d50c8eea0, []int{0}
}
func (m *ProxyProtocol) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProxyProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProxyProtocol.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProxyProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyProtocol.Merge(m, src)
}
func (m *ProxyProtocol) XXX_Size() int {
	return m.Size()
}
func (m *ProxyProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyProtocol proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProxyProtocol)(nil), "envoy.config.filter.listener.proxy_protocol.v2.ProxyProtocol")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/proxy_protocol/v2/proxy_protocol.proto", fileDescriptor_614eea5d50c8eea0)
}

var fileDescriptor_614eea5d50c8eea0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0x8c, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x33, 0x42, 0x13, 0xd1, 0x03, 0x33, 0x84,
	0xf4, 0xc0, 0x86, 0xe8, 0x41, 0x0c, 0xd1, 0x83, 0x18, 0xa2, 0x07, 0x33, 0x44, 0x0f, 0x4d, 0x4b,
	0x99, 0x91, 0x94, 0x5c, 0x69, 0x4a, 0x41, 0xa2, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49,
	0x66, 0x7e, 0x5e, 0xb1, 0x7e, 0x6e, 0x66, 0x7a, 0x51, 0x62, 0x49, 0x2a, 0xc4, 0x3c, 0x29, 0x59,
	0x0c, 0xf9, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x88, 0xb4, 0x12, 0x3f, 0x17, 0x6f, 0x00, 0xc8,
	0xcc, 0x00, 0xa8, 0x91, 0x4e, 0x8b, 0x19, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x4f, 0x33, 0xfe, 0xf5, 0xb3, 0x9a, 0x0a, 0x19, 0x43, 0x1c, 0x95, 0x5a,
	0x51, 0x92, 0x9a, 0x57, 0x0c, 0x32, 0x04, 0xea, 0xb0, 0x62, 0xdc, 0x2e, 0x33, 0xde, 0xd5, 0x70,
	0xe2, 0x22, 0x1b, 0x93, 0x00, 0x03, 0x97, 0x4d, 0x66, 0x3e, 0xc4, 0x53, 0x60, 0x25, 0x24, 0xfa,
	0xcf, 0x49, 0x08, 0xc5, 0x79, 0x60, 0x3a, 0x80, 0x31, 0x89, 0x0d, 0xac, 0xc4, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0x73, 0x93, 0xe2, 0x73, 0x01, 0x00, 0x00,
}

func (m *ProxyProtocol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProxyProtocol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProxyProtocol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintProxyProtocol(dAtA []byte, offset int, v uint64) int {
	offset -= sovProxyProtocol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProxyProtocol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovProxyProtocol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProxyProtocol(x uint64) (n int) {
	return sovProxyProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProxyProtocol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProxyProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProxyProtocol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProxyProtocol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProxyProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProxyProtocol
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthProxyProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProxyProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProxyProtocol
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
					return 0, ErrIntOverflowProxyProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProxyProtocol
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
			if length < 0 {
				return 0, ErrInvalidLengthProxyProtocol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProxyProtocol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProxyProtocol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProxyProtocol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProxyProtocol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProxyProtocol = fmt.Errorf("proto: unexpected end of group")
)

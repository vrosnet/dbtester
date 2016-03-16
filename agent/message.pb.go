// Code generated by protoc-gen-gogo.
// source: message.proto
// DO NOT EDIT!

/*
	Package agent is a generated protocol buffer package.

	It is generated from these files:
		message.proto

	It has these top-level messages:
		Request
		Response
*/
package agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	grpc "github.com/coreos/etcd/Godeps/_workspace/src/google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Request_Operation int32

const (
	Request_Start   Request_Operation = 0
	Request_Stop    Request_Operation = 1
	Request_Restart Request_Operation = 2
)

var Request_Operation_name = map[int32]string{
	0: "Start",
	1: "Stop",
	2: "Restart",
}
var Request_Operation_value = map[string]int32{
	"Start":   0,
	"Stop":    1,
	"Restart": 2,
}

func (x Request_Operation) String() string {
	return proto.EnumName(Request_Operation_name, int32(x))
}
func (Request_Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 0} }

type Request_Database int32

const (
	Request_etcd      Request_Database = 0
	Request_ZooKeeper Request_Database = 1
)

var Request_Database_name = map[int32]string{
	0: "etcd",
	1: "ZooKeeper",
}
var Request_Database_value = map[string]int32{
	"etcd":      0,
	"ZooKeeper": 1,
}

func (x Request_Database) String() string {
	return proto.EnumName(Request_Database_name, int32(x))
}
func (Request_Database) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 1} }

type Request struct {
	Operation Request_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=agent.Request_Operation" json:"operation,omitempty"`
	Database  Request_Database  `protobuf:"varint,2,opt,name=database,proto3,enum=agent.Request_Database" json:"database,omitempty"`
	PeerIPs   string            `protobuf:"bytes,3,opt,name=peerIPs,proto3" json:"peerIPs,omitempty"`
	// EtcdServerIndex is the index in peerIPs.
	EtcdServerIndex uint32 `protobuf:"varint,4,opt,name=etcdServerIndex,proto3" json:"etcdServerIndex,omitempty"`
	ZookeeperMyID   uint32 `protobuf:"varint,5,opt,name=zookeeperMyID,proto3" json:"zookeeperMyID,omitempty"`
	// ZookeeperPreAllocSize avoid seeks ZooKeeper allocates space
	// in the transaction log file in blocks of PreAllocSize kilobytes.
	// Default value is 65536 * 1024.
	ZookeeperPreAllocSize int64 `protobuf:"varint,6,opt,name=zookeeperPreAllocSize,proto3" json:"zookeeperPreAllocSize,omitempty"`
	// ZookeeperMaxClientCnxns limits the number of concurrent connections
	// (at the socket level) that a single client, identified by IP address.
	ZookeeperMaxClientCnxns int64 `protobuf:"varint,7,opt,name=zookeeperMaxClientCnxns,proto3" json:"zookeeperMaxClientCnxns,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

type Response struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "agent.Request")
	proto.RegisterType((*Response)(nil), "agent.Response")
	proto.RegisterEnum("agent.Request_Operation", Request_Operation_name, Request_Operation_value)
	proto.RegisterEnum("agent.Request_Database", Request_Database_name, Request_Database_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Transporter service

type TransporterClient interface {
	Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type transporterClient struct {
	cc *grpc.ClientConn
}

func NewTransporterClient(cc *grpc.ClientConn) TransporterClient {
	return &transporterClient{cc}
}

func (c *transporterClient) Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/agent.Transporter/Transfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Transporter service

type TransporterServer interface {
	Transfer(context.Context, *Request) (*Response, error)
}

func RegisterTransporterServer(s *grpc.Server, srv TransporterServer) {
	s.RegisterService(&_Transporter_serviceDesc, srv)
}

func _Transporter_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(TransporterServer).Transfer(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Transporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Transporter",
	HandlerType: (*TransporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Transporter_Transfer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Operation != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintMessage(data, i, uint64(m.Operation))
	}
	if m.Database != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintMessage(data, i, uint64(m.Database))
	}
	if len(m.PeerIPs) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintMessage(data, i, uint64(len(m.PeerIPs)))
		i += copy(data[i:], m.PeerIPs)
	}
	if m.EtcdServerIndex != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintMessage(data, i, uint64(m.EtcdServerIndex))
	}
	if m.ZookeeperMyID != 0 {
		data[i] = 0x28
		i++
		i = encodeVarintMessage(data, i, uint64(m.ZookeeperMyID))
	}
	if m.ZookeeperPreAllocSize != 0 {
		data[i] = 0x30
		i++
		i = encodeVarintMessage(data, i, uint64(m.ZookeeperPreAllocSize))
	}
	if m.ZookeeperMaxClientCnxns != 0 {
		data[i] = 0x38
		i++
		i = encodeVarintMessage(data, i, uint64(m.ZookeeperMaxClientCnxns))
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		data[i] = 0x8
		i++
		if m.Success {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Message(data []byte, offset int, v uint64) int {
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
func encodeFixed32Message(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMessage(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Operation != 0 {
		n += 1 + sovMessage(uint64(m.Operation))
	}
	if m.Database != 0 {
		n += 1 + sovMessage(uint64(m.Database))
	}
	l = len(m.PeerIPs)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.EtcdServerIndex != 0 {
		n += 1 + sovMessage(uint64(m.EtcdServerIndex))
	}
	if m.ZookeeperMyID != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperMyID))
	}
	if m.ZookeeperPreAllocSize != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperPreAllocSize))
	}
	if m.ZookeeperMaxClientCnxns != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperMaxClientCnxns))
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Operation |= (Request_Operation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			m.Database = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Database |= (Request_Database(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerIPs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerIPs = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdServerIndex", wireType)
			}
			m.EtcdServerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.EtcdServerIndex |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperMyID", wireType)
			}
			m.ZookeeperMyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ZookeeperMyID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperPreAllocSize", wireType)
			}
			m.ZookeeperPreAllocSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ZookeeperPreAllocSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperMaxClientCnxns", wireType)
			}
			m.ZookeeperMaxClientCnxns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ZookeeperMaxClientCnxns |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Response) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(data[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorMessage = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0x6e, 0x2f, 0x97, 0x4b, 0x7b, 0x6e, 0x7a, 0x2f, 0x99, 0x68, 0xee, 0x84, 0x05, 0x31, 0x95,
	0x05, 0x1b, 0x4a, 0x02, 0xc6, 0xb8, 0x70, 0xa3, 0xb0, 0x21, 0xc6, 0x48, 0x5a, 0x57, 0xee, 0xa6,
	0xe5, 0x50, 0x1b, 0xa1, 0x53, 0x67, 0xa6, 0x06, 0x79, 0x12, 0x1f, 0xc0, 0x87, 0x61, 0xe9, 0x23,
	0xf8, 0xf3, 0x22, 0x0e, 0x03, 0x2d, 0x81, 0xe8, 0x62, 0x92, 0xf3, 0xfd, 0x9d, 0xf4, 0x9c, 0x53,
	0xf0, 0xd6, 0x28, 0x25, 0x4b, 0x31, 0x28, 0x04, 0x57, 0x9c, 0x34, 0x75, 0x99, 0xab, 0xce, 0x20,
	0xcd, 0xd4, 0xc7, 0x32, 0x0e, 0x12, 0xbe, 0x1e, 0xa6, 0x3c, 0xe5, 0x43, 0xa3, 0xc6, 0xe5, 0xd2,
	0x20, 0x03, 0x4c, 0x75, 0x48, 0xf9, 0xdf, 0x1b, 0xd0, 0x0a, 0xf1, 0x73, 0x89, 0x52, 0x91, 0xe7,
	0xe0, 0xf2, 0x02, 0x05, 0x53, 0x19, 0xcf, 0xa9, 0xfd, 0xc4, 0xee, 0xdf, 0x8d, 0x68, 0x60, 0xba,
	0x06, 0x47, 0x4b, 0xf0, 0xae, 0xd2, 0xc3, 0x93, 0x95, 0x8c, 0xc1, 0x59, 0x30, 0xc5, 0x62, 0x26,
	0x91, 0x5e, 0x99, 0xd8, 0xc3, 0x45, 0x6c, 0x7a, 0x94, 0xc3, 0xda, 0x48, 0x28, 0xb4, 0x0a, 0x44,
	0x31, 0x9b, 0x4b, 0xda, 0xd0, 0x19, 0x37, 0xac, 0x20, 0xe9, 0xc3, 0x3d, 0xaa, 0x64, 0x11, 0xa1,
	0xf8, 0xa2, 0x89, 0x7c, 0x81, 0x1b, 0x7a, 0xad, 0x1d, 0x5e, 0x78, 0x49, 0x93, 0x1e, 0x78, 0x5b,
	0xce, 0x3f, 0x21, 0xea, 0x4f, 0x79, 0xfb, 0x75, 0x36, 0xa5, 0x4d, 0xe3, 0x3b, 0x27, 0xc9, 0x33,
	0x78, 0x5c, 0x13, 0x73, 0x81, 0xaf, 0x56, 0x2b, 0x9e, 0x44, 0xd9, 0x16, 0xe9, 0x8d, 0x76, 0x37,
	0xc2, 0x7f, 0x8b, 0xe4, 0x05, 0x3c, 0x9c, 0xda, 0xb0, 0xcd, 0x64, 0x95, 0xe9, 0x81, 0x26, 0xf9,
	0x26, 0x97, 0xb4, 0x65, 0x72, 0xff, 0x93, 0xfd, 0x01, 0xb8, 0xf5, 0x9a, 0x88, 0x0b, 0xcd, 0x48,
	0x31, 0xa1, 0xda, 0x16, 0x71, 0xe0, 0x3a, 0x52, 0xbc, 0x68, 0xdb, 0xe4, 0x76, 0xbf, 0x73, 0x69,
	0xe8, 0x2b, 0xff, 0x29, 0x38, 0xd5, 0x7a, 0xf6, 0x96, 0xfd, 0x8c, 0xda, 0xec, 0x81, 0xfb, 0x81,
	0xf3, 0x37, 0xa6, 0x7f, 0xdb, 0xf6, 0x7b, 0xe0, 0xe8, 0x44, 0xc1, 0xf3, 0xc3, 0xe6, 0x64, 0x99,
	0x24, 0xfa, 0xf8, 0xe6, 0x48, 0x4e, 0x58, 0xc1, 0xd1, 0x4b, 0xb8, 0x7d, 0x2f, 0x58, 0xae, 0x7d,
	0x42, 0xa1, 0x20, 0x03, 0x70, 0x0c, 0x5c, 0xea, 0xfa, 0xee, 0xfc, 0x22, 0x9d, 0xfb, 0x1a, 0x1f,
	0xba, 0xfa, 0xd6, 0xeb, 0x47, 0xbb, 0x5f, 0x5d, 0x6b, 0xf7, 0xbb, 0x6b, 0xff, 0xd0, 0xef, 0xa7,
	0x7e, 0xdf, 0xfe, 0x74, 0xad, 0xf8, 0xc6, 0xfc, 0x27, 0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x62, 0x4a, 0xb3, 0x66, 0x6e, 0x02, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/message.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Employee struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BadgeNumber          int32    `protobuf:"varint,2,opt,name=badgeNumber,proto3" json:"badgeNumber,omitempty"`
	Firstname            string   `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Employee) Reset()         { *m = Employee{} }
func (m *Employee) String() string { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()    {}
func (*Employee) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{0}
}

func (m *Employee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Employee.Unmarshal(m, b)
}
func (m *Employee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Employee.Marshal(b, m, deterministic)
}
func (m *Employee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Employee.Merge(m, src)
}
func (m *Employee) XXX_Size() int {
	return xxx_messageInfo_Employee.Size(m)
}
func (m *Employee) XXX_DiscardUnknown() {
	xxx_messageInfo_Employee.DiscardUnknown(m)
}

var xxx_messageInfo_Employee proto.InternalMessageInfo

func (m *Employee) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetBadgeNumber() int32 {
	if m != nil {
		return m.BadgeNumber
	}
	return 0
}

func (m *Employee) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Employee) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

type EmployeeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmployeeRequest) Reset()         { *m = EmployeeRequest{} }
func (m *EmployeeRequest) String() string { return proto.CompactTextString(m) }
func (*EmployeeRequest) ProtoMessage()    {}
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{1}
}

func (m *EmployeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeRequest.Unmarshal(m, b)
}
func (m *EmployeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeRequest.Marshal(b, m, deterministic)
}
func (m *EmployeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeRequest.Merge(m, src)
}
func (m *EmployeeRequest) XXX_Size() int {
	return xxx_messageInfo_EmployeeRequest.Size(m)
}
func (m *EmployeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeRequest proto.InternalMessageInfo

type EmployeeResponse struct {
	Employee             *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EmployeeResponse) Reset()         { *m = EmployeeResponse{} }
func (m *EmployeeResponse) String() string { return proto.CompactTextString(m) }
func (*EmployeeResponse) ProtoMessage()    {}
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8447775385e7eb85, []int{2}
}

func (m *EmployeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmployeeResponse.Unmarshal(m, b)
}
func (m *EmployeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmployeeResponse.Marshal(b, m, deterministic)
}
func (m *EmployeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmployeeResponse.Merge(m, src)
}
func (m *EmployeeResponse) XXX_Size() int {
	return xxx_messageInfo_EmployeeResponse.Size(m)
}
func (m *EmployeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmployeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmployeeResponse proto.InternalMessageInfo

func (m *EmployeeResponse) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

func init() {
	proto.RegisterType((*Employee)(nil), "Employee")
	proto.RegisterType((*EmployeeRequest)(nil), "EmployeeRequest")
	proto.RegisterType((*EmployeeResponse)(nil), "EmployeeResponse")
}

func init() { proto.RegisterFile("pb/message.proto", fileDescriptor_8447775385e7eb85) }

var fileDescriptor_8447775385e7eb85 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x40, 0x69, 0x5c, 0xa5, 0x9d, 0x05, 0xb7, 0x9b, 0x53, 0x59, 0x3c, 0x94, 0x82, 0xb0, 0xa7,
	0x08, 0x15, 0x04, 0x8f, 0x0a, 0x22, 0x5e, 0x3c, 0xd4, 0x9b, 0xb7, 0xc4, 0x8e, 0x25, 0x90, 0x34,
	0x31, 0x49, 0x0b, 0xfe, 0x7b, 0x21, 0xf4, 0x43, 0x7a, 0xcc, 0x7b, 0x61, 0xde, 0x30, 0x90, 0x5b,
	0x71, 0xa7, 0xd1, 0x7b, 0xde, 0x21, 0xb3, 0xce, 0x04, 0x53, 0x8d, 0x90, 0xbe, 0x68, 0xab, 0xcc,
	0x2f, 0x22, 0xbd, 0x06, 0x22, 0xdb, 0x22, 0x29, 0x93, 0xf3, 0x65, 0x43, 0x64, 0x4b, 0x4b, 0xd8,
	0x0b, 0xde, 0x76, 0xf8, 0x3e, 0x68, 0x81, 0xae, 0x20, 0x51, 0xfc, 0x47, 0xf4, 0x06, 0xb2, 0x6f,
	0xe9, 0x7c, 0xe8, 0xb9, 0xc6, 0xe2, 0xa2, 0x4c, 0xce, 0x59, 0xb3, 0x02, 0x7a, 0x82, 0x54, 0xf1,
	0x49, 0xee, 0xa2, 0x5c, 0xde, 0xd5, 0x11, 0x0e, 0x73, 0xb7, 0xc1, 0x9f, 0x01, 0x7d, 0xa8, 0x1e,
	0x21, 0x5f, 0x91, 0xb7, 0xa6, 0xf7, 0x48, 0x6f, 0x21, 0xc5, 0x89, 0xc5, 0xc5, 0xf6, 0x75, 0xc6,
	0x96, 0x4f, 0x8b, 0xaa, 0xdf, 0xd6, 0x69, 0x1f, 0xe8, 0x46, 0xf9, 0x85, 0xf4, 0x01, 0x0e, 0xaf,
	0x18, 0x9e, 0x94, 0x9a, 0x85, 0xa7, 0x39, 0xdb, 0x24, 0x4f, 0x47, 0xb6, 0x2d, 0x3e, 0xef, 0x3e,
	0x89, 0x15, 0xe2, 0x2a, 0x5e, 0xe7, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x43, 0x20, 0xaa, 0xef,
	0x31, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	GetAllEmployees(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
}

type employeeServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmployeeServiceClient(cc *grpc.ClientConn) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetAllEmployees(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/EmployeeService/GetAllEmployees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
type EmployeeServiceServer interface {
	GetAllEmployees(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
}

// UnimplementedEmployeeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (*UnimplementedEmployeeServiceServer) GetAllEmployees(ctx context.Context, req *EmployeeRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEmployees not implemented")
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_GetAllEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetAllEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/GetAllEmployees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetAllEmployees(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllEmployees",
			Handler:    _EmployeeService_GetAllEmployees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/message.proto",
}

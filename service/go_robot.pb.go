// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: go_robot.proto

package service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y uint32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_robot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_go_robot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_go_robot_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type MouseButton struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *MouseButton) Reset() {
	*x = MouseButton{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_robot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MouseButton) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MouseButton) ProtoMessage() {}

func (x *MouseButton) ProtoReflect() protoreflect.Message {
	mi := &file_go_robot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MouseButton.ProtoReflect.Descriptor instead.
func (*MouseButton) Descriptor() ([]byte, []int) {
	return file_go_robot_proto_rawDescGZIP(), []int{1}
}

func (x *MouseButton) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

type MouseScrolledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance uint32 `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *MouseScrolledRequest) Reset() {
	*x = MouseScrolledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_robot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MouseScrolledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MouseScrolledRequest) ProtoMessage() {}

func (x *MouseScrolledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_robot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MouseScrolledRequest.ProtoReflect.Descriptor instead.
func (*MouseScrolledRequest) Descriptor() ([]byte, []int) {
	return file_go_robot_proto_rawDescGZIP(), []int{2}
}

func (x *MouseScrolledRequest) GetDistance() uint32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_robot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_go_robot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_go_robot_proto_rawDescGZIP(), []int{3}
}

type CaptureScreenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  uint32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *CaptureScreenRequest) Reset() {
	*x = CaptureScreenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_robot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureScreenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureScreenRequest) ProtoMessage() {}

func (x *CaptureScreenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_robot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureScreenRequest.ProtoReflect.Descriptor instead.
func (*CaptureScreenRequest) Descriptor() ([]byte, []int) {
	return file_go_robot_proto_rawDescGZIP(), []int{4}
}

func (x *CaptureScreenRequest) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *CaptureScreenRequest) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type CaptureScreenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bitmap []byte `protobuf:"bytes,1,opt,name=bitmap,proto3" json:"bitmap,omitempty"`
}

func (x *CaptureScreenResponse) Reset() {
	*x = CaptureScreenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_robot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureScreenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureScreenResponse) ProtoMessage() {}

func (x *CaptureScreenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_robot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureScreenResponse.ProtoReflect.Descriptor instead.
func (*CaptureScreenResponse) Descriptor() ([]byte, []int) {
	return file_go_robot_proto_rawDescGZIP(), []int{5}
}

func (x *CaptureScreenResponse) GetBitmap() []byte {
	if x != nil {
		return x.Bitmap
	}
	return nil
}

var File_go_robot_proto protoreflect.FileDescriptor

var file_go_robot_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x6f, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79, 0x22, 0x1b,
	0x0a, 0x0b, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x78, 0x22, 0x32, 0x0a, 0x14, 0x4d,
	0x6f, 0x75, 0x73, 0x65, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x44, 0x0a, 0x14, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2f,
	0x0a, 0x15, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x74, 0x6d, 0x61,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x32,
	0xf9, 0x03, 0x0a, 0x07, 0x47, 0x6f, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x4d,
	0x6f, 0x75, 0x73, 0x65, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x4d, 0x6f,
	0x75, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x1a, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x10, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x6f, 0x75, 0x73, 0x65, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c,
	0x4d, 0x6f, 0x75, 0x73, 0x65, 0x50, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4d, 0x6f, 0x75, 0x73, 0x65, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0f, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x55, 0x70,
	0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x73, 0x65,
	0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x11, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_go_robot_proto_rawDescOnce sync.Once
	file_go_robot_proto_rawDescData = file_go_robot_proto_rawDesc
)

func file_go_robot_proto_rawDescGZIP() []byte {
	file_go_robot_proto_rawDescOnce.Do(func() {
		file_go_robot_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_robot_proto_rawDescData)
	})
	return file_go_robot_proto_rawDescData
}

var file_go_robot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_go_robot_proto_goTypes = []interface{}{
	(*Point)(nil),                 // 0: service.Point
	(*MouseButton)(nil),           // 1: service.MouseButton
	(*MouseScrolledRequest)(nil),  // 2: service.MouseScrolledRequest
	(*Empty)(nil),                 // 3: service.Empty
	(*CaptureScreenRequest)(nil),  // 4: service.CaptureScreenRequest
	(*CaptureScreenResponse)(nil), // 5: service.CaptureScreenResponse
}
var file_go_robot_proto_depIdxs = []int32{
	0, // 0: service.GoRobot.MouseMove:input_type -> service.Point
	1, // 1: service.GoRobot.MouseClick:input_type -> service.MouseButton
	1, // 2: service.GoRobot.MouseDoubleClick:input_type -> service.MouseButton
	1, // 3: service.GoRobot.MousePressed:input_type -> service.MouseButton
	1, // 4: service.GoRobot.MouseReleased:input_type -> service.MouseButton
	2, // 5: service.GoRobot.MouseScrolledUp:input_type -> service.MouseScrolledRequest
	2, // 6: service.GoRobot.MouseScrolledDown:input_type -> service.MouseScrolledRequest
	4, // 7: service.GoRobot.CaptureScreen:input_type -> service.CaptureScreenRequest
	3, // 8: service.GoRobot.MouseMove:output_type -> service.Empty
	3, // 9: service.GoRobot.MouseClick:output_type -> service.Empty
	3, // 10: service.GoRobot.MouseDoubleClick:output_type -> service.Empty
	3, // 11: service.GoRobot.MousePressed:output_type -> service.Empty
	3, // 12: service.GoRobot.MouseReleased:output_type -> service.Empty
	3, // 13: service.GoRobot.MouseScrolledUp:output_type -> service.Empty
	3, // 14: service.GoRobot.MouseScrolledDown:output_type -> service.Empty
	5, // 15: service.GoRobot.CaptureScreen:output_type -> service.CaptureScreenResponse
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_robot_proto_init() }
func file_go_robot_proto_init() {
	if File_go_robot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_robot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_go_robot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MouseButton); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_go_robot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MouseScrolledRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_go_robot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_go_robot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureScreenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_go_robot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureScreenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_robot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_robot_proto_goTypes,
		DependencyIndexes: file_go_robot_proto_depIdxs,
		MessageInfos:      file_go_robot_proto_msgTypes,
	}.Build()
	File_go_robot_proto = out.File
	file_go_robot_proto_rawDesc = nil
	file_go_robot_proto_goTypes = nil
	file_go_robot_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GoRobotClient is the client API for GoRobot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoRobotClient interface {
	MouseMove(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Empty, error)
	MouseClick(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error)
	MouseDoubleClick(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error)
	MousePressed(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error)
	MouseReleased(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error)
	MouseScrolledUp(ctx context.Context, in *MouseScrolledRequest, opts ...grpc.CallOption) (*Empty, error)
	MouseScrolledDown(ctx context.Context, in *MouseScrolledRequest, opts ...grpc.CallOption) (*Empty, error)
	CaptureScreen(ctx context.Context, in *CaptureScreenRequest, opts ...grpc.CallOption) (GoRobot_CaptureScreenClient, error)
}

type goRobotClient struct {
	cc grpc.ClientConnInterface
}

func NewGoRobotClient(cc grpc.ClientConnInterface) GoRobotClient {
	return &goRobotClient{cc}
}

func (c *goRobotClient) MouseMove(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.GoRobot/MouseMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRobotClient) MouseClick(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.GoRobot/MouseClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRobotClient) MouseDoubleClick(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.GoRobot/MouseDoubleClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRobotClient) MousePressed(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.GoRobot/MousePressed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRobotClient) MouseReleased(ctx context.Context, in *MouseButton, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.GoRobot/MouseReleased", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRobotClient) MouseScrolledUp(ctx context.Context, in *MouseScrolledRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.GoRobot/MouseScrolledUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRobotClient) MouseScrolledDown(ctx context.Context, in *MouseScrolledRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/service.GoRobot/MouseScrolledDown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goRobotClient) CaptureScreen(ctx context.Context, in *CaptureScreenRequest, opts ...grpc.CallOption) (GoRobot_CaptureScreenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GoRobot_serviceDesc.Streams[0], "/service.GoRobot/CaptureScreen", opts...)
	if err != nil {
		return nil, err
	}
	x := &goRobotCaptureScreenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoRobot_CaptureScreenClient interface {
	Recv() (*CaptureScreenResponse, error)
	grpc.ClientStream
}

type goRobotCaptureScreenClient struct {
	grpc.ClientStream
}

func (x *goRobotCaptureScreenClient) Recv() (*CaptureScreenResponse, error) {
	m := new(CaptureScreenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GoRobotServer is the server API for GoRobot service.
type GoRobotServer interface {
	MouseMove(context.Context, *Point) (*Empty, error)
	MouseClick(context.Context, *MouseButton) (*Empty, error)
	MouseDoubleClick(context.Context, *MouseButton) (*Empty, error)
	MousePressed(context.Context, *MouseButton) (*Empty, error)
	MouseReleased(context.Context, *MouseButton) (*Empty, error)
	MouseScrolledUp(context.Context, *MouseScrolledRequest) (*Empty, error)
	MouseScrolledDown(context.Context, *MouseScrolledRequest) (*Empty, error)
	CaptureScreen(*CaptureScreenRequest, GoRobot_CaptureScreenServer) error
}

// UnimplementedGoRobotServer can be embedded to have forward compatible implementations.
type UnimplementedGoRobotServer struct {
}

func (*UnimplementedGoRobotServer) MouseMove(context.Context, *Point) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MouseMove not implemented")
}
func (*UnimplementedGoRobotServer) MouseClick(context.Context, *MouseButton) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MouseClick not implemented")
}
func (*UnimplementedGoRobotServer) MouseDoubleClick(context.Context, *MouseButton) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MouseDoubleClick not implemented")
}
func (*UnimplementedGoRobotServer) MousePressed(context.Context, *MouseButton) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MousePressed not implemented")
}
func (*UnimplementedGoRobotServer) MouseReleased(context.Context, *MouseButton) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MouseReleased not implemented")
}
func (*UnimplementedGoRobotServer) MouseScrolledUp(context.Context, *MouseScrolledRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MouseScrolledUp not implemented")
}
func (*UnimplementedGoRobotServer) MouseScrolledDown(context.Context, *MouseScrolledRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MouseScrolledDown not implemented")
}
func (*UnimplementedGoRobotServer) CaptureScreen(*CaptureScreenRequest, GoRobot_CaptureScreenServer) error {
	return status.Errorf(codes.Unimplemented, "method CaptureScreen not implemented")
}

func RegisterGoRobotServer(s *grpc.Server, srv GoRobotServer) {
	s.RegisterService(&_GoRobot_serviceDesc, srv)
}

func _GoRobot_MouseMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRobotServer).MouseMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.GoRobot/MouseMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRobotServer).MouseMove(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRobot_MouseClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MouseButton)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRobotServer).MouseClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.GoRobot/MouseClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRobotServer).MouseClick(ctx, req.(*MouseButton))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRobot_MouseDoubleClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MouseButton)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRobotServer).MouseDoubleClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.GoRobot/MouseDoubleClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRobotServer).MouseDoubleClick(ctx, req.(*MouseButton))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRobot_MousePressed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MouseButton)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRobotServer).MousePressed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.GoRobot/MousePressed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRobotServer).MousePressed(ctx, req.(*MouseButton))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRobot_MouseReleased_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MouseButton)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRobotServer).MouseReleased(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.GoRobot/MouseReleased",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRobotServer).MouseReleased(ctx, req.(*MouseButton))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRobot_MouseScrolledUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MouseScrolledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRobotServer).MouseScrolledUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.GoRobot/MouseScrolledUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRobotServer).MouseScrolledUp(ctx, req.(*MouseScrolledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRobot_MouseScrolledDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MouseScrolledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoRobotServer).MouseScrolledDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.GoRobot/MouseScrolledDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoRobotServer).MouseScrolledDown(ctx, req.(*MouseScrolledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoRobot_CaptureScreen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CaptureScreenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoRobotServer).CaptureScreen(m, &goRobotCaptureScreenServer{stream})
}

type GoRobot_CaptureScreenServer interface {
	Send(*CaptureScreenResponse) error
	grpc.ServerStream
}

type goRobotCaptureScreenServer struct {
	grpc.ServerStream
}

func (x *goRobotCaptureScreenServer) Send(m *CaptureScreenResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GoRobot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.GoRobot",
	HandlerType: (*GoRobotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MouseMove",
			Handler:    _GoRobot_MouseMove_Handler,
		},
		{
			MethodName: "MouseClick",
			Handler:    _GoRobot_MouseClick_Handler,
		},
		{
			MethodName: "MouseDoubleClick",
			Handler:    _GoRobot_MouseDoubleClick_Handler,
		},
		{
			MethodName: "MousePressed",
			Handler:    _GoRobot_MousePressed_Handler,
		},
		{
			MethodName: "MouseReleased",
			Handler:    _GoRobot_MouseReleased_Handler,
		},
		{
			MethodName: "MouseScrolledUp",
			Handler:    _GoRobot_MouseScrolledUp_Handler,
		},
		{
			MethodName: "MouseScrolledDown",
			Handler:    _GoRobot_MouseScrolledDown_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CaptureScreen",
			Handler:       _GoRobot_CaptureScreen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "go_robot.proto",
}

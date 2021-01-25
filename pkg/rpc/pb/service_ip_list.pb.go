// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_ip_list.proto

package pb

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

// 创建IP列表
type CreateIPListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	TimeoutJSON []byte `protobuf:"bytes,4,opt,name=timeoutJSON,proto3" json:"timeoutJSON,omitempty"`
}

func (x *CreateIPListRequest) Reset() {
	*x = CreateIPListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIPListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIPListRequest) ProtoMessage() {}

func (x *CreateIPListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIPListRequest.ProtoReflect.Descriptor instead.
func (*CreateIPListRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_list_proto_rawDescGZIP(), []int{0}
}

func (x *CreateIPListRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateIPListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateIPListRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateIPListRequest) GetTimeoutJSON() []byte {
	if x != nil {
		return x.TimeoutJSON
	}
	return nil
}

type CreateIPListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpListId int64 `protobuf:"varint,1,opt,name=ipListId,proto3" json:"ipListId,omitempty"`
}

func (x *CreateIPListResponse) Reset() {
	*x = CreateIPListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIPListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIPListResponse) ProtoMessage() {}

func (x *CreateIPListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIPListResponse.ProtoReflect.Descriptor instead.
func (*CreateIPListResponse) Descriptor() ([]byte, []int) {
	return file_service_ip_list_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIPListResponse) GetIpListId() int64 {
	if x != nil {
		return x.IpListId
	}
	return 0
}

// 修改IP列表
type UpdateIPListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpListId    int64  `protobuf:"varint,1,opt,name=ipListId,proto3" json:"ipListId,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	TimeoutJSON []byte `protobuf:"bytes,4,opt,name=timeoutJSON,proto3" json:"timeoutJSON,omitempty"`
}

func (x *UpdateIPListRequest) Reset() {
	*x = UpdateIPListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIPListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIPListRequest) ProtoMessage() {}

func (x *UpdateIPListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIPListRequest.ProtoReflect.Descriptor instead.
func (*UpdateIPListRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_list_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateIPListRequest) GetIpListId() int64 {
	if x != nil {
		return x.IpListId
	}
	return 0
}

func (x *UpdateIPListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateIPListRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateIPListRequest) GetTimeoutJSON() []byte {
	if x != nil {
		return x.TimeoutJSON
	}
	return nil
}

// 查找IP列表
type FindEnabledIPListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpListId int64 `protobuf:"varint,1,opt,name=ipListId,proto3" json:"ipListId,omitempty"`
}

func (x *FindEnabledIPListRequest) Reset() {
	*x = FindEnabledIPListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledIPListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledIPListRequest) ProtoMessage() {}

func (x *FindEnabledIPListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledIPListRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledIPListRequest) Descriptor() ([]byte, []int) {
	return file_service_ip_list_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledIPListRequest) GetIpListId() int64 {
	if x != nil {
		return x.IpListId
	}
	return 0
}

type FindEnabledIPListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpList *IPList `protobuf:"bytes,1,opt,name=ipList,proto3" json:"ipList,omitempty"`
}

func (x *FindEnabledIPListResponse) Reset() {
	*x = FindEnabledIPListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ip_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledIPListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledIPListResponse) ProtoMessage() {}

func (x *FindEnabledIPListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ip_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledIPListResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledIPListResponse) Descriptor() ([]byte, []int) {
	return file_service_ip_list_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledIPListResponse) GetIpList() *IPList {
	if x != nil {
		return x.IpList
	}
	return nil
}

var File_service_ip_list_proto protoreflect.FileDescriptor

var file_service_ip_list_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x73, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x32, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x36, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x22, 0x3f, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x49,
	0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x06, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x69, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0xdd, 0x01, 0x0a, 0x0d, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x50, 0x0a, 0x11, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x49, 0x50,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_ip_list_proto_rawDescOnce sync.Once
	file_service_ip_list_proto_rawDescData = file_service_ip_list_proto_rawDesc
)

func file_service_ip_list_proto_rawDescGZIP() []byte {
	file_service_ip_list_proto_rawDescOnce.Do(func() {
		file_service_ip_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ip_list_proto_rawDescData)
	})
	return file_service_ip_list_proto_rawDescData
}

var file_service_ip_list_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_ip_list_proto_goTypes = []interface{}{
	(*CreateIPListRequest)(nil),       // 0: pb.CreateIPListRequest
	(*CreateIPListResponse)(nil),      // 1: pb.CreateIPListResponse
	(*UpdateIPListRequest)(nil),       // 2: pb.UpdateIPListRequest
	(*FindEnabledIPListRequest)(nil),  // 3: pb.FindEnabledIPListRequest
	(*FindEnabledIPListResponse)(nil), // 4: pb.FindEnabledIPListResponse
	(*IPList)(nil),                    // 5: pb.IPList
	(*RPCSuccess)(nil),                // 6: pb.RPCSuccess
}
var file_service_ip_list_proto_depIdxs = []int32{
	5, // 0: pb.FindEnabledIPListResponse.ipList:type_name -> pb.IPList
	0, // 1: pb.IPListService.createIPList:input_type -> pb.CreateIPListRequest
	2, // 2: pb.IPListService.updateIPList:input_type -> pb.UpdateIPListRequest
	3, // 3: pb.IPListService.findEnabledIPList:input_type -> pb.FindEnabledIPListRequest
	1, // 4: pb.IPListService.createIPList:output_type -> pb.CreateIPListResponse
	6, // 5: pb.IPListService.updateIPList:output_type -> pb.RPCSuccess
	4, // 6: pb.IPListService.findEnabledIPList:output_type -> pb.FindEnabledIPListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_ip_list_proto_init() }
func file_service_ip_list_proto_init() {
	if File_service_ip_list_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_ip_list_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ip_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIPListRequest); i {
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
		file_service_ip_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIPListResponse); i {
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
		file_service_ip_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIPListRequest); i {
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
		file_service_ip_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledIPListRequest); i {
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
		file_service_ip_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledIPListResponse); i {
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
			RawDescriptor: file_service_ip_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ip_list_proto_goTypes,
		DependencyIndexes: file_service_ip_list_proto_depIdxs,
		MessageInfos:      file_service_ip_list_proto_msgTypes,
	}.Build()
	File_service_ip_list_proto = out.File
	file_service_ip_list_proto_rawDesc = nil
	file_service_ip_list_proto_goTypes = nil
	file_service_ip_list_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IPListServiceClient is the client API for IPListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IPListServiceClient interface {
	// 创建IP列表
	CreateIPList(ctx context.Context, in *CreateIPListRequest, opts ...grpc.CallOption) (*CreateIPListResponse, error)
	// 修改IP列表
	UpdateIPList(ctx context.Context, in *UpdateIPListRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找IP列表
	FindEnabledIPList(ctx context.Context, in *FindEnabledIPListRequest, opts ...grpc.CallOption) (*FindEnabledIPListResponse, error)
}

type iPListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIPListServiceClient(cc grpc.ClientConnInterface) IPListServiceClient {
	return &iPListServiceClient{cc}
}

func (c *iPListServiceClient) CreateIPList(ctx context.Context, in *CreateIPListRequest, opts ...grpc.CallOption) (*CreateIPListResponse, error) {
	out := new(CreateIPListResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/createIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) UpdateIPList(ctx context.Context, in *UpdateIPListRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.IPListService/updateIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPListServiceClient) FindEnabledIPList(ctx context.Context, in *FindEnabledIPListRequest, opts ...grpc.CallOption) (*FindEnabledIPListResponse, error) {
	out := new(FindEnabledIPListResponse)
	err := c.cc.Invoke(ctx, "/pb.IPListService/findEnabledIPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPListServiceServer is the server API for IPListService service.
type IPListServiceServer interface {
	// 创建IP列表
	CreateIPList(context.Context, *CreateIPListRequest) (*CreateIPListResponse, error)
	// 修改IP列表
	UpdateIPList(context.Context, *UpdateIPListRequest) (*RPCSuccess, error)
	// 查找IP列表
	FindEnabledIPList(context.Context, *FindEnabledIPListRequest) (*FindEnabledIPListResponse, error)
}

// UnimplementedIPListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIPListServiceServer struct {
}

func (*UnimplementedIPListServiceServer) CreateIPList(context.Context, *CreateIPListRequest) (*CreateIPListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIPList not implemented")
}
func (*UnimplementedIPListServiceServer) UpdateIPList(context.Context, *UpdateIPListRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIPList not implemented")
}
func (*UnimplementedIPListServiceServer) FindEnabledIPList(context.Context, *FindEnabledIPListRequest) (*FindEnabledIPListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledIPList not implemented")
}

func RegisterIPListServiceServer(s *grpc.Server, srv IPListServiceServer) {
	s.RegisterService(&_IPListService_serviceDesc, srv)
}

func _IPListService_CreateIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).CreateIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/CreateIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).CreateIPList(ctx, req.(*CreateIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_UpdateIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).UpdateIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/UpdateIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).UpdateIPList(ctx, req.(*UpdateIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPListService_FindEnabledIPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledIPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPListServiceServer).FindEnabledIPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.IPListService/FindEnabledIPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPListServiceServer).FindEnabledIPList(ctx, req.(*FindEnabledIPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IPListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IPListService",
	HandlerType: (*IPListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createIPList",
			Handler:    _IPListService_CreateIPList_Handler,
		},
		{
			MethodName: "updateIPList",
			Handler:    _IPListService_UpdateIPList_Handler,
		},
		{
			MethodName: "findEnabledIPList",
			Handler:    _IPListService_FindEnabledIPList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ip_list.proto",
}

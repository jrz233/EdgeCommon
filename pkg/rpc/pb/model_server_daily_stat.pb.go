// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_server_daily_stat.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// 服务每日统计
type ServerDailyStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId             int64 `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	RegionId             int64 `protobuf:"varint,2,opt,name=regionId,proto3" json:"regionId,omitempty"`
	Bytes                int64 `protobuf:"varint,3,opt,name=bytes,proto3" json:"bytes,omitempty"`
	CachedBytes          int64 `protobuf:"varint,5,opt,name=cachedBytes,proto3" json:"cachedBytes,omitempty"`
	CountRequests        int64 `protobuf:"varint,6,opt,name=countRequests,proto3" json:"countRequests,omitempty"`
	CountCachedRequests  int64 `protobuf:"varint,7,opt,name=countCachedRequests,proto3" json:"countCachedRequests,omitempty"`
	CreatedAt            int64 `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	CountAttackRequests  int64 `protobuf:"varint,8,opt,name=countAttackRequests,proto3" json:"countAttackRequests,omitempty"`
	AttackBytes          int64 `protobuf:"varint,9,opt,name=attackBytes,proto3" json:"attackBytes,omitempty"`
	CheckTrafficLimiting bool  `protobuf:"varint,10,opt,name=checkTrafficLimiting,proto3" json:"checkTrafficLimiting,omitempty"`
	PlanId               int64 `protobuf:"varint,11,opt,name=planId,proto3" json:"planId,omitempty"`
}

func (x *ServerDailyStat) Reset() {
	*x = ServerDailyStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_server_daily_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerDailyStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerDailyStat) ProtoMessage() {}

func (x *ServerDailyStat) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_server_daily_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerDailyStat.ProtoReflect.Descriptor instead.
func (*ServerDailyStat) Descriptor() ([]byte, []int) {
	return file_models_model_server_daily_stat_proto_rawDescGZIP(), []int{0}
}

func (x *ServerDailyStat) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *ServerDailyStat) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *ServerDailyStat) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

func (x *ServerDailyStat) GetCachedBytes() int64 {
	if x != nil {
		return x.CachedBytes
	}
	return 0
}

func (x *ServerDailyStat) GetCountRequests() int64 {
	if x != nil {
		return x.CountRequests
	}
	return 0
}

func (x *ServerDailyStat) GetCountCachedRequests() int64 {
	if x != nil {
		return x.CountCachedRequests
	}
	return 0
}

func (x *ServerDailyStat) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ServerDailyStat) GetCountAttackRequests() int64 {
	if x != nil {
		return x.CountAttackRequests
	}
	return 0
}

func (x *ServerDailyStat) GetAttackBytes() int64 {
	if x != nil {
		return x.AttackBytes
	}
	return 0
}

func (x *ServerDailyStat) GetCheckTrafficLimiting() bool {
	if x != nil {
		return x.CheckTrafficLimiting
	}
	return false
}

func (x *ServerDailyStat) GetPlanId() int64 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

var File_models_model_server_daily_stat_proto protoreflect.FileDescriptor

var file_models_model_server_daily_stat_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x97, 0x03, 0x0a, 0x0f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_server_daily_stat_proto_rawDescOnce sync.Once
	file_models_model_server_daily_stat_proto_rawDescData = file_models_model_server_daily_stat_proto_rawDesc
)

func file_models_model_server_daily_stat_proto_rawDescGZIP() []byte {
	file_models_model_server_daily_stat_proto_rawDescOnce.Do(func() {
		file_models_model_server_daily_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_server_daily_stat_proto_rawDescData)
	})
	return file_models_model_server_daily_stat_proto_rawDescData
}

var file_models_model_server_daily_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_server_daily_stat_proto_goTypes = []interface{}{
	(*ServerDailyStat)(nil), // 0: pb.ServerDailyStat
}
var file_models_model_server_daily_stat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_server_daily_stat_proto_init() }
func file_models_model_server_daily_stat_proto_init() {
	if File_models_model_server_daily_stat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_server_daily_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerDailyStat); i {
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
			RawDescriptor: file_models_model_server_daily_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_server_daily_stat_proto_goTypes,
		DependencyIndexes: file_models_model_server_daily_stat_proto_depIdxs,
		MessageInfos:      file_models_model_server_daily_stat_proto_msgTypes,
	}.Build()
	File_models_model_server_daily_stat_proto = out.File
	file_models_model_server_daily_stat_proto_rawDesc = nil
	file_models_model_server_daily_stat_proto_goTypes = nil
	file_models_model_server_daily_stat_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_server_region_city_monthly_stat.proto

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

// 查找前N个城市
type FindTopServerRegionCityMonthlyStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month      string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty"` // YYYYMM
	ServerId   int64  `protobuf:"varint,2,opt,name=serverId,proto3" json:"serverId,omitempty"`
	CountryId  int64  `protobuf:"varint,3,opt,name=countryId,proto3" json:"countryId,omitempty"`
	ProvinceId int64  `protobuf:"varint,4,opt,name=provinceId,proto3" json:"provinceId,omitempty"`
	Offset     int64  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Size       int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) Reset() {
	*x = FindTopServerRegionCityMonthlyStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_city_monthly_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionCityMonthlyStatsRequest) ProtoMessage() {}

func (x *FindTopServerRegionCityMonthlyStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_city_monthly_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionCityMonthlyStatsRequest.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionCityMonthlyStatsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_region_city_monthly_stat_proto_rawDescGZIP(), []int{0}
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) GetProvinceId() int64 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindTopServerRegionCityMonthlyStatsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindTopServerRegionCityMonthlyStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*FindTopServerRegionCityMonthlyStatsResponse_Stat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *FindTopServerRegionCityMonthlyStatsResponse) Reset() {
	*x = FindTopServerRegionCityMonthlyStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_city_monthly_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionCityMonthlyStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionCityMonthlyStatsResponse) ProtoMessage() {}

func (x *FindTopServerRegionCityMonthlyStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_city_monthly_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionCityMonthlyStatsResponse.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionCityMonthlyStatsResponse) Descriptor() ([]byte, []int) {
	return file_service_server_region_city_monthly_stat_proto_rawDescGZIP(), []int{1}
}

func (x *FindTopServerRegionCityMonthlyStatsResponse) GetStats() []*FindTopServerRegionCityMonthlyStatsResponse_Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type FindTopServerRegionCityMonthlyStatsResponse_Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCountry  *RegionCountry  `protobuf:"bytes,1,opt,name=regionCountry,proto3" json:"regionCountry,omitempty"`
	RegionProvince *RegionProvince `protobuf:"bytes,2,opt,name=regionProvince,proto3" json:"regionProvince,omitempty"`
	RegionCity     *RegionCity     `protobuf:"bytes,3,opt,name=regionCity,proto3" json:"regionCity,omitempty"`
	Count          int64           `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FindTopServerRegionCityMonthlyStatsResponse_Stat) Reset() {
	*x = FindTopServerRegionCityMonthlyStatsResponse_Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_region_city_monthly_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerRegionCityMonthlyStatsResponse_Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerRegionCityMonthlyStatsResponse_Stat) ProtoMessage() {}

func (x *FindTopServerRegionCityMonthlyStatsResponse_Stat) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_region_city_monthly_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerRegionCityMonthlyStatsResponse_Stat.ProtoReflect.Descriptor instead.
func (*FindTopServerRegionCityMonthlyStatsResponse_Stat) Descriptor() ([]byte, []int) {
	return file_service_server_region_city_monthly_stat_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindTopServerRegionCityMonthlyStatsResponse_Stat) GetRegionCountry() *RegionCountry {
	if x != nil {
		return x.RegionCountry
	}
	return nil
}

func (x *FindTopServerRegionCityMonthlyStatsResponse_Stat) GetRegionProvince() *RegionProvince {
	if x != nil {
		return x.RegionProvince
	}
	return nil
}

func (x *FindTopServerRegionCityMonthlyStatsResponse_Stat) GetRegionCity() *RegionCity {
	if x != nil {
		return x.RegionCity
	}
	return nil
}

func (x *FindTopServerRegionCityMonthlyStatsResponse_Stat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_service_server_region_city_monthly_stat_proto protoreflect.FileDescriptor

var file_service_server_region_city_monthly_stat_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x2a, 0x46,
	0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xbd, 0x02, 0x0a, 0x2b, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x1a, 0xc1, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52,
	0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12,
	0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xad, 0x01, 0x0a, 0x22, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a,
	0x23, 0x66, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x69, 0x74,
	0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_region_city_monthly_stat_proto_rawDescOnce sync.Once
	file_service_server_region_city_monthly_stat_proto_rawDescData = file_service_server_region_city_monthly_stat_proto_rawDesc
)

func file_service_server_region_city_monthly_stat_proto_rawDescGZIP() []byte {
	file_service_server_region_city_monthly_stat_proto_rawDescOnce.Do(func() {
		file_service_server_region_city_monthly_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_region_city_monthly_stat_proto_rawDescData)
	})
	return file_service_server_region_city_monthly_stat_proto_rawDescData
}

var file_service_server_region_city_monthly_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_server_region_city_monthly_stat_proto_goTypes = []interface{}{
	(*FindTopServerRegionCityMonthlyStatsRequest)(nil),       // 0: pb.FindTopServerRegionCityMonthlyStatsRequest
	(*FindTopServerRegionCityMonthlyStatsResponse)(nil),      // 1: pb.FindTopServerRegionCityMonthlyStatsResponse
	(*FindTopServerRegionCityMonthlyStatsResponse_Stat)(nil), // 2: pb.FindTopServerRegionCityMonthlyStatsResponse.Stat
	(*RegionCountry)(nil),                                    // 3: pb.RegionCountry
	(*RegionProvince)(nil),                                   // 4: pb.RegionProvince
	(*RegionCity)(nil),                                       // 5: pb.RegionCity
}
var file_service_server_region_city_monthly_stat_proto_depIdxs = []int32{
	2, // 0: pb.FindTopServerRegionCityMonthlyStatsResponse.stats:type_name -> pb.FindTopServerRegionCityMonthlyStatsResponse.Stat
	3, // 1: pb.FindTopServerRegionCityMonthlyStatsResponse.Stat.regionCountry:type_name -> pb.RegionCountry
	4, // 2: pb.FindTopServerRegionCityMonthlyStatsResponse.Stat.regionProvince:type_name -> pb.RegionProvince
	5, // 3: pb.FindTopServerRegionCityMonthlyStatsResponse.Stat.regionCity:type_name -> pb.RegionCity
	0, // 4: pb.ServerRegionCityMonthlyStatService.findTopServerRegionCityMonthlyStats:input_type -> pb.FindTopServerRegionCityMonthlyStatsRequest
	1, // 5: pb.ServerRegionCityMonthlyStatService.findTopServerRegionCityMonthlyStats:output_type -> pb.FindTopServerRegionCityMonthlyStatsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_server_region_city_monthly_stat_proto_init() }
func file_service_server_region_city_monthly_stat_proto_init() {
	if File_service_server_region_city_monthly_stat_proto != nil {
		return
	}
	file_models_model_region_country_proto_init()
	file_models_model_region_province_proto_init()
	file_models_model_region_city_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_region_city_monthly_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionCityMonthlyStatsRequest); i {
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
		file_service_server_region_city_monthly_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionCityMonthlyStatsResponse); i {
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
		file_service_server_region_city_monthly_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerRegionCityMonthlyStatsResponse_Stat); i {
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
			RawDescriptor: file_service_server_region_city_monthly_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_region_city_monthly_stat_proto_goTypes,
		DependencyIndexes: file_service_server_region_city_monthly_stat_proto_depIdxs,
		MessageInfos:      file_service_server_region_city_monthly_stat_proto_msgTypes,
	}.Build()
	File_service_server_region_city_monthly_stat_proto = out.File
	file_service_server_region_city_monthly_stat_proto_rawDesc = nil
	file_service_server_region_city_monthly_stat_proto_goTypes = nil
	file_service_server_region_city_monthly_stat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerRegionCityMonthlyStatServiceClient is the client API for ServerRegionCityMonthlyStatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerRegionCityMonthlyStatServiceClient interface {
	// 查找前N个城市
	FindTopServerRegionCityMonthlyStats(ctx context.Context, in *FindTopServerRegionCityMonthlyStatsRequest, opts ...grpc.CallOption) (*FindTopServerRegionCityMonthlyStatsResponse, error)
}

type serverRegionCityMonthlyStatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerRegionCityMonthlyStatServiceClient(cc grpc.ClientConnInterface) ServerRegionCityMonthlyStatServiceClient {
	return &serverRegionCityMonthlyStatServiceClient{cc}
}

func (c *serverRegionCityMonthlyStatServiceClient) FindTopServerRegionCityMonthlyStats(ctx context.Context, in *FindTopServerRegionCityMonthlyStatsRequest, opts ...grpc.CallOption) (*FindTopServerRegionCityMonthlyStatsResponse, error) {
	out := new(FindTopServerRegionCityMonthlyStatsResponse)
	err := c.cc.Invoke(ctx, "/pb.ServerRegionCityMonthlyStatService/findTopServerRegionCityMonthlyStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerRegionCityMonthlyStatServiceServer is the server API for ServerRegionCityMonthlyStatService service.
type ServerRegionCityMonthlyStatServiceServer interface {
	// 查找前N个城市
	FindTopServerRegionCityMonthlyStats(context.Context, *FindTopServerRegionCityMonthlyStatsRequest) (*FindTopServerRegionCityMonthlyStatsResponse, error)
}

// UnimplementedServerRegionCityMonthlyStatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerRegionCityMonthlyStatServiceServer struct {
}

func (*UnimplementedServerRegionCityMonthlyStatServiceServer) FindTopServerRegionCityMonthlyStats(context.Context, *FindTopServerRegionCityMonthlyStatsRequest) (*FindTopServerRegionCityMonthlyStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTopServerRegionCityMonthlyStats not implemented")
}

func RegisterServerRegionCityMonthlyStatServiceServer(s *grpc.Server, srv ServerRegionCityMonthlyStatServiceServer) {
	s.RegisterService(&_ServerRegionCityMonthlyStatService_serviceDesc, srv)
}

func _ServerRegionCityMonthlyStatService_FindTopServerRegionCityMonthlyStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTopServerRegionCityMonthlyStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerRegionCityMonthlyStatServiceServer).FindTopServerRegionCityMonthlyStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ServerRegionCityMonthlyStatService/FindTopServerRegionCityMonthlyStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerRegionCityMonthlyStatServiceServer).FindTopServerRegionCityMonthlyStats(ctx, req.(*FindTopServerRegionCityMonthlyStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerRegionCityMonthlyStatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerRegionCityMonthlyStatService",
	HandlerType: (*ServerRegionCityMonthlyStatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findTopServerRegionCityMonthlyStats",
			Handler:    _ServerRegionCityMonthlyStatService_FindTopServerRegionCityMonthlyStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_region_city_monthly_stat.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_message_task.proto

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

// 创建消息任务
type CreateMessageTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientId int64  `protobuf:"varint,1,opt,name=recipientId,proto3" json:"recipientId,omitempty"`
	InstanceId  int64  `protobuf:"varint,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	User        string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Subject     string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body        string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	IsPrimary   bool   `protobuf:"varint,6,opt,name=isPrimary,proto3" json:"isPrimary,omitempty"`
}

func (x *CreateMessageTaskRequest) Reset() {
	*x = CreateMessageTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageTaskRequest) ProtoMessage() {}

func (x *CreateMessageTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageTaskRequest) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageTaskRequest) GetRecipientId() int64 {
	if x != nil {
		return x.RecipientId
	}
	return 0
}

func (x *CreateMessageTaskRequest) GetInstanceId() int64 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *CreateMessageTaskRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateMessageTaskRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *CreateMessageTaskRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CreateMessageTaskRequest) GetIsPrimary() bool {
	if x != nil {
		return x.IsPrimary
	}
	return false
}

type CreateMessageTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTaskId int64 `protobuf:"varint,1,opt,name=messageTaskId,proto3" json:"messageTaskId,omitempty"`
}

func (x *CreateMessageTaskResponse) Reset() {
	*x = CreateMessageTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageTaskResponse) ProtoMessage() {}

func (x *CreateMessageTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageTaskResponse) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageTaskResponse) GetMessageTaskId() int64 {
	if x != nil {
		return x.MessageTaskId
	}
	return 0
}

// 查找要发送的消息任务
type FindSendingMessageTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindSendingMessageTasksRequest) Reset() {
	*x = FindSendingMessageTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSendingMessageTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSendingMessageTasksRequest) ProtoMessage() {}

func (x *FindSendingMessageTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSendingMessageTasksRequest.ProtoReflect.Descriptor instead.
func (*FindSendingMessageTasksRequest) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{2}
}

func (x *FindSendingMessageTasksRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindSendingMessageTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTasks []*MessageTask `protobuf:"bytes,1,rep,name=messageTasks,proto3" json:"messageTasks,omitempty"`
}

func (x *FindSendingMessageTasksResponse) Reset() {
	*x = FindSendingMessageTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindSendingMessageTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindSendingMessageTasksResponse) ProtoMessage() {}

func (x *FindSendingMessageTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindSendingMessageTasksResponse.ProtoReflect.Descriptor instead.
func (*FindSendingMessageTasksResponse) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{3}
}

func (x *FindSendingMessageTasksResponse) GetMessageTasks() []*MessageTask {
	if x != nil {
		return x.MessageTasks
	}
	return nil
}

// 修改消息任务状态
type UpdateMessageTaskStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTaskId int64              `protobuf:"varint,1,opt,name=messageTaskId,proto3" json:"messageTaskId,omitempty"`
	Status        int32              `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Result        *MessageTaskResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateMessageTaskStatusRequest) Reset() {
	*x = UpdateMessageTaskStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageTaskStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageTaskStatusRequest) ProtoMessage() {}

func (x *UpdateMessageTaskStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageTaskStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageTaskStatusRequest) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMessageTaskStatusRequest) GetMessageTaskId() int64 {
	if x != nil {
		return x.MessageTaskId
	}
	return 0
}

func (x *UpdateMessageTaskStatusRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateMessageTaskStatusRequest) GetResult() *MessageTaskResult {
	if x != nil {
		return x.Result
	}
	return nil
}

// 删除消息任务
type DeleteMessageTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTaskId int64 `protobuf:"varint,1,opt,name=messageTaskId,proto3" json:"messageTaskId,omitempty"`
}

func (x *DeleteMessageTaskRequest) Reset() {
	*x = DeleteMessageTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageTaskRequest) ProtoMessage() {}

func (x *DeleteMessageTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageTaskRequest) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMessageTaskRequest) GetMessageTaskId() int64 {
	if x != nil {
		return x.MessageTaskId
	}
	return 0
}

// 读取消息任务状态
type FindEnabledMessageTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTaskId int64 `protobuf:"varint,1,opt,name=messageTaskId,proto3" json:"messageTaskId,omitempty"`
}

func (x *FindEnabledMessageTaskRequest) Reset() {
	*x = FindEnabledMessageTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageTaskRequest) ProtoMessage() {}

func (x *FindEnabledMessageTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageTaskRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageTaskRequest) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledMessageTaskRequest) GetMessageTaskId() int64 {
	if x != nil {
		return x.MessageTaskId
	}
	return 0
}

type FindEnabledMessageTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTask *MessageTask `protobuf:"bytes,1,opt,name=messageTask,proto3" json:"messageTask,omitempty"`
}

func (x *FindEnabledMessageTaskResponse) Reset() {
	*x = FindEnabledMessageTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledMessageTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledMessageTaskResponse) ProtoMessage() {}

func (x *FindEnabledMessageTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledMessageTaskResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledMessageTaskResponse) Descriptor() ([]byte, []int) {
	return file_service_message_task_proto_rawDescGZIP(), []int{7}
}

func (x *FindEnabledMessageTaskResponse) GetMessageTask() *MessageTask {
	if x != nil {
		return x.MessageTask
	}
	return nil
}

var File_service_message_task_proto protoreflect.FileDescriptor

var file_service_message_task_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x41, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64,
	0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x56,
	0x0a, 0x1f, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22,
	0x53, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x32, 0xbd, 0x03, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x17, 0x66, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x41, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x5f, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_message_task_proto_rawDescOnce sync.Once
	file_service_message_task_proto_rawDescData = file_service_message_task_proto_rawDesc
)

func file_service_message_task_proto_rawDescGZIP() []byte {
	file_service_message_task_proto_rawDescOnce.Do(func() {
		file_service_message_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_message_task_proto_rawDescData)
	})
	return file_service_message_task_proto_rawDescData
}

var file_service_message_task_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_message_task_proto_goTypes = []interface{}{
	(*CreateMessageTaskRequest)(nil),        // 0: pb.CreateMessageTaskRequest
	(*CreateMessageTaskResponse)(nil),       // 1: pb.CreateMessageTaskResponse
	(*FindSendingMessageTasksRequest)(nil),  // 2: pb.FindSendingMessageTasksRequest
	(*FindSendingMessageTasksResponse)(nil), // 3: pb.FindSendingMessageTasksResponse
	(*UpdateMessageTaskStatusRequest)(nil),  // 4: pb.UpdateMessageTaskStatusRequest
	(*DeleteMessageTaskRequest)(nil),        // 5: pb.DeleteMessageTaskRequest
	(*FindEnabledMessageTaskRequest)(nil),   // 6: pb.FindEnabledMessageTaskRequest
	(*FindEnabledMessageTaskResponse)(nil),  // 7: pb.FindEnabledMessageTaskResponse
	(*MessageTask)(nil),                     // 8: pb.MessageTask
	(*MessageTaskResult)(nil),               // 9: pb.MessageTaskResult
	(*RPCSuccess)(nil),                      // 10: pb.RPCSuccess
}
var file_service_message_task_proto_depIdxs = []int32{
	8,  // 0: pb.FindSendingMessageTasksResponse.messageTasks:type_name -> pb.MessageTask
	9,  // 1: pb.UpdateMessageTaskStatusRequest.result:type_name -> pb.MessageTaskResult
	8,  // 2: pb.FindEnabledMessageTaskResponse.messageTask:type_name -> pb.MessageTask
	0,  // 3: pb.MessageTaskService.createMessageTask:input_type -> pb.CreateMessageTaskRequest
	2,  // 4: pb.MessageTaskService.findSendingMessageTasks:input_type -> pb.FindSendingMessageTasksRequest
	4,  // 5: pb.MessageTaskService.updateMessageTaskStatus:input_type -> pb.UpdateMessageTaskStatusRequest
	5,  // 6: pb.MessageTaskService.deleteMessageTask:input_type -> pb.DeleteMessageTaskRequest
	6,  // 7: pb.MessageTaskService.findEnabledMessageTask:input_type -> pb.FindEnabledMessageTaskRequest
	1,  // 8: pb.MessageTaskService.createMessageTask:output_type -> pb.CreateMessageTaskResponse
	3,  // 9: pb.MessageTaskService.findSendingMessageTasks:output_type -> pb.FindSendingMessageTasksResponse
	10, // 10: pb.MessageTaskService.updateMessageTaskStatus:output_type -> pb.RPCSuccess
	10, // 11: pb.MessageTaskService.deleteMessageTask:output_type -> pb.RPCSuccess
	7,  // 12: pb.MessageTaskService.findEnabledMessageTask:output_type -> pb.FindEnabledMessageTaskResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_service_message_task_proto_init() }
func file_service_message_task_proto_init() {
	if File_service_message_task_proto != nil {
		return
	}
	file_models_message_task_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_message_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageTaskRequest); i {
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
		file_service_message_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageTaskResponse); i {
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
		file_service_message_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSendingMessageTasksRequest); i {
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
		file_service_message_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindSendingMessageTasksResponse); i {
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
		file_service_message_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageTaskStatusRequest); i {
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
		file_service_message_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageTaskRequest); i {
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
		file_service_message_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageTaskRequest); i {
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
		file_service_message_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledMessageTaskResponse); i {
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
			RawDescriptor: file_service_message_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_message_task_proto_goTypes,
		DependencyIndexes: file_service_message_task_proto_depIdxs,
		MessageInfos:      file_service_message_task_proto_msgTypes,
	}.Build()
	File_service_message_task_proto = out.File
	file_service_message_task_proto_rawDesc = nil
	file_service_message_task_proto_goTypes = nil
	file_service_message_task_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageTaskServiceClient is the client API for MessageTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageTaskServiceClient interface {
	// 创建消息任务
	CreateMessageTask(ctx context.Context, in *CreateMessageTaskRequest, opts ...grpc.CallOption) (*CreateMessageTaskResponse, error)
	// 查找要发送的消息任务
	FindSendingMessageTasks(ctx context.Context, in *FindSendingMessageTasksRequest, opts ...grpc.CallOption) (*FindSendingMessageTasksResponse, error)
	// 修改消息任务状态
	UpdateMessageTaskStatus(ctx context.Context, in *UpdateMessageTaskStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除消息任务
	DeleteMessageTask(ctx context.Context, in *DeleteMessageTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 读取消息任务状态
	FindEnabledMessageTask(ctx context.Context, in *FindEnabledMessageTaskRequest, opts ...grpc.CallOption) (*FindEnabledMessageTaskResponse, error)
}

type messageTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageTaskServiceClient(cc grpc.ClientConnInterface) MessageTaskServiceClient {
	return &messageTaskServiceClient{cc}
}

func (c *messageTaskServiceClient) CreateMessageTask(ctx context.Context, in *CreateMessageTaskRequest, opts ...grpc.CallOption) (*CreateMessageTaskResponse, error) {
	out := new(CreateMessageTaskResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageTaskService/createMessageTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageTaskServiceClient) FindSendingMessageTasks(ctx context.Context, in *FindSendingMessageTasksRequest, opts ...grpc.CallOption) (*FindSendingMessageTasksResponse, error) {
	out := new(FindSendingMessageTasksResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageTaskService/findSendingMessageTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageTaskServiceClient) UpdateMessageTaskStatus(ctx context.Context, in *UpdateMessageTaskStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MessageTaskService/updateMessageTaskStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageTaskServiceClient) DeleteMessageTask(ctx context.Context, in *DeleteMessageTaskRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.MessageTaskService/deleteMessageTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageTaskServiceClient) FindEnabledMessageTask(ctx context.Context, in *FindEnabledMessageTaskRequest, opts ...grpc.CallOption) (*FindEnabledMessageTaskResponse, error) {
	out := new(FindEnabledMessageTaskResponse)
	err := c.cc.Invoke(ctx, "/pb.MessageTaskService/findEnabledMessageTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageTaskServiceServer is the server API for MessageTaskService service.
type MessageTaskServiceServer interface {
	// 创建消息任务
	CreateMessageTask(context.Context, *CreateMessageTaskRequest) (*CreateMessageTaskResponse, error)
	// 查找要发送的消息任务
	FindSendingMessageTasks(context.Context, *FindSendingMessageTasksRequest) (*FindSendingMessageTasksResponse, error)
	// 修改消息任务状态
	UpdateMessageTaskStatus(context.Context, *UpdateMessageTaskStatusRequest) (*RPCSuccess, error)
	// 删除消息任务
	DeleteMessageTask(context.Context, *DeleteMessageTaskRequest) (*RPCSuccess, error)
	// 读取消息任务状态
	FindEnabledMessageTask(context.Context, *FindEnabledMessageTaskRequest) (*FindEnabledMessageTaskResponse, error)
}

// UnimplementedMessageTaskServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageTaskServiceServer struct {
}

func (*UnimplementedMessageTaskServiceServer) CreateMessageTask(context.Context, *CreateMessageTaskRequest) (*CreateMessageTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessageTask not implemented")
}
func (*UnimplementedMessageTaskServiceServer) FindSendingMessageTasks(context.Context, *FindSendingMessageTasksRequest) (*FindSendingMessageTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSendingMessageTasks not implemented")
}
func (*UnimplementedMessageTaskServiceServer) UpdateMessageTaskStatus(context.Context, *UpdateMessageTaskStatusRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageTaskStatus not implemented")
}
func (*UnimplementedMessageTaskServiceServer) DeleteMessageTask(context.Context, *DeleteMessageTaskRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessageTask not implemented")
}
func (*UnimplementedMessageTaskServiceServer) FindEnabledMessageTask(context.Context, *FindEnabledMessageTaskRequest) (*FindEnabledMessageTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledMessageTask not implemented")
}

func RegisterMessageTaskServiceServer(s *grpc.Server, srv MessageTaskServiceServer) {
	s.RegisterService(&_MessageTaskService_serviceDesc, srv)
}

func _MessageTaskService_CreateMessageTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageTaskServiceServer).CreateMessageTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageTaskService/CreateMessageTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageTaskServiceServer).CreateMessageTask(ctx, req.(*CreateMessageTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageTaskService_FindSendingMessageTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSendingMessageTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageTaskServiceServer).FindSendingMessageTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageTaskService/FindSendingMessageTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageTaskServiceServer).FindSendingMessageTasks(ctx, req.(*FindSendingMessageTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageTaskService_UpdateMessageTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageTaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageTaskServiceServer).UpdateMessageTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageTaskService/UpdateMessageTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageTaskServiceServer).UpdateMessageTaskStatus(ctx, req.(*UpdateMessageTaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageTaskService_DeleteMessageTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageTaskServiceServer).DeleteMessageTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageTaskService/DeleteMessageTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageTaskServiceServer).DeleteMessageTask(ctx, req.(*DeleteMessageTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageTaskService_FindEnabledMessageTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledMessageTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageTaskServiceServer).FindEnabledMessageTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageTaskService/FindEnabledMessageTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageTaskServiceServer).FindEnabledMessageTask(ctx, req.(*FindEnabledMessageTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageTaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageTaskService",
	HandlerType: (*MessageTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createMessageTask",
			Handler:    _MessageTaskService_CreateMessageTask_Handler,
		},
		{
			MethodName: "findSendingMessageTasks",
			Handler:    _MessageTaskService_FindSendingMessageTasks_Handler,
		},
		{
			MethodName: "updateMessageTaskStatus",
			Handler:    _MessageTaskService_UpdateMessageTaskStatus_Handler,
		},
		{
			MethodName: "deleteMessageTask",
			Handler:    _MessageTaskService_DeleteMessageTask_Handler,
		},
		{
			MethodName: "findEnabledMessageTask",
			Handler:    _MessageTaskService_FindEnabledMessageTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_message_task.proto",
}
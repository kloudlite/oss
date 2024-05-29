// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: infra.proto

package infra

import (
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

type GetClusterIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail   string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	AccountName string `protobuf:"bytes,4,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ClusterName string `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *GetClusterIn) Reset() {
	*x = GetClusterIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterIn) ProtoMessage() {}

func (x *GetClusterIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterIn.ProtoReflect.Descriptor instead.
func (*GetClusterIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{0}
}

func (x *GetClusterIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetClusterIn) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetClusterIn) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *GetClusterIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetClusterIn) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type GetClusterOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageQueueTopic string `protobuf:"bytes,1,opt,name=messageQueueTopic,proto3" json:"messageQueueTopic,omitempty"`
	DnsHost           string `protobuf:"bytes,2,opt,name=dnsHost,proto3" json:"dnsHost,omitempty"`
	IACJobName        string `protobuf:"bytes,3,opt,name=IACJobName,proto3" json:"IACJobName,omitempty"`
	IACJobNamespace   string `protobuf:"bytes,4,opt,name=IACJobNamespace,proto3" json:"IACJobNamespace,omitempty"`
}

func (x *GetClusterOut) Reset() {
	*x = GetClusterOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterOut) ProtoMessage() {}

func (x *GetClusterOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterOut.ProtoReflect.Descriptor instead.
func (*GetClusterOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{1}
}

func (x *GetClusterOut) GetMessageQueueTopic() string {
	if x != nil {
		return x.MessageQueueTopic
	}
	return ""
}

func (x *GetClusterOut) GetDnsHost() string {
	if x != nil {
		return x.DnsHost
	}
	return ""
}

func (x *GetClusterOut) GetIACJobName() string {
	if x != nil {
		return x.IACJobName
	}
	return ""
}

func (x *GetClusterOut) GetIACJobNamespace() string {
	if x != nil {
		return x.IACJobNamespace
	}
	return ""
}

type GetNodepoolIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName     string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail    string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	AccountName  string `protobuf:"bytes,4,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ClusterName  string `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	NodepoolName string `protobuf:"bytes,6,opt,name=nodepoolName,proto3" json:"nodepoolName,omitempty"`
}

func (x *GetNodepoolIn) Reset() {
	*x = GetNodepoolIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodepoolIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodepoolIn) ProtoMessage() {}

func (x *GetNodepoolIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodepoolIn.ProtoReflect.Descriptor instead.
func (*GetNodepoolIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodepoolIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetNodepoolIn) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetNodepoolIn) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *GetNodepoolIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetNodepoolIn) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GetNodepoolIn) GetNodepoolName() string {
	if x != nil {
		return x.NodepoolName
	}
	return ""
}

type GetNodepoolOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IACJobName      string `protobuf:"bytes,1,opt,name=IACJobName,proto3" json:"IACJobName,omitempty"`
	IACJobNamespace string `protobuf:"bytes,2,opt,name=IACJobNamespace,proto3" json:"IACJobNamespace,omitempty"`
}

func (x *GetNodepoolOut) Reset() {
	*x = GetNodepoolOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodepoolOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodepoolOut) ProtoMessage() {}

func (x *GetNodepoolOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodepoolOut.ProtoReflect.Descriptor instead.
func (*GetNodepoolOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{3}
}

func (x *GetNodepoolOut) GetIACJobName() string {
	if x != nil {
		return x.IACJobName
	}
	return ""
}

func (x *GetNodepoolOut) GetIACJobNamespace() string {
	if x != nil {
		return x.IACJobNamespace
	}
	return ""
}

type ClusterExistsIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail   string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	AccountName string `protobuf:"bytes,4,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ClusterName string `protobuf:"bytes,5,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *ClusterExistsIn) Reset() {
	*x = ClusterExistsIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterExistsIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterExistsIn) ProtoMessage() {}

func (x *ClusterExistsIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterExistsIn.ProtoReflect.Descriptor instead.
func (*ClusterExistsIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{4}
}

func (x *ClusterExistsIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ClusterExistsIn) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ClusterExistsIn) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *ClusterExistsIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *ClusterExistsIn) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type ClusterExistsOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *ClusterExistsOut) Reset() {
	*x = ClusterExistsOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterExistsOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterExistsOut) ProtoMessage() {}

func (x *ClusterExistsOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterExistsOut.ProtoReflect.Descriptor instead.
func (*ClusterExistsOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{5}
}

func (x *ClusterExistsOut) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type GetClusterKubeconfigOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kubeconfig []byte `protobuf:"bytes,1,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
}

func (x *GetClusterKubeconfigOut) Reset() {
	*x = GetClusterKubeconfigOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterKubeconfigOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterKubeconfigOut) ProtoMessage() {}

func (x *GetClusterKubeconfigOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterKubeconfigOut.ProtoReflect.Descriptor instead.
func (*GetClusterKubeconfigOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{6}
}

func (x *GetClusterKubeconfigOut) GetKubeconfig() []byte {
	if x != nil {
		return x.Kubeconfig
	}
	return nil
}

type GetClusterManagedServiceIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail   string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	AccountName string `protobuf:"bytes,4,opt,name=accountName,proto3" json:"accountName,omitempty"`
	MsvcName    string `protobuf:"bytes,5,opt,name=msvcName,proto3" json:"msvcName,omitempty"`
}

func (x *GetClusterManagedServiceIn) Reset() {
	*x = GetClusterManagedServiceIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterManagedServiceIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterManagedServiceIn) ProtoMessage() {}

func (x *GetClusterManagedServiceIn) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterManagedServiceIn.ProtoReflect.Descriptor instead.
func (*GetClusterManagedServiceIn) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{7}
}

func (x *GetClusterManagedServiceIn) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetClusterManagedServiceIn) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetClusterManagedServiceIn) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *GetClusterManagedServiceIn) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *GetClusterManagedServiceIn) GetMsvcName() string {
	if x != nil {
		return x.MsvcName
	}
	return ""
}

type GetClusterManagedServiceOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetNamespace string `protobuf:"bytes,1,opt,name=targetNamespace,proto3" json:"targetNamespace,omitempty"`
	ClusterName     string `protobuf:"bytes,2,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *GetClusterManagedServiceOut) Reset() {
	*x = GetClusterManagedServiceOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterManagedServiceOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterManagedServiceOut) ProtoMessage() {}

func (x *GetClusterManagedServiceOut) ProtoReflect() protoreflect.Message {
	mi := &file_infra_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterManagedServiceOut.ProtoReflect.Descriptor instead.
func (*GetClusterManagedServiceOut) Descriptor() ([]byte, []int) {
	return file_infra_proto_rawDescGZIP(), []int{8}
}

func (x *GetClusterManagedServiceOut) GetTargetNamespace() string {
	if x != nil {
		return x.TargetNamespace
	}
	return ""
}

func (x *GetClusterManagedServiceOut) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

var File_infra_proto protoreflect.FileDescriptor

var file_infra_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6e, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x70,
	0x6f, 0x6f, 0x6c, 0x4f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x41, 0x43, 0x4a,
	0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x49, 0x41, 0x43, 0x4a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x22, 0xa7, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x75,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0xac, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x76, 0x63, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x73, 0x76, 0x63, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x69, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xb2, 0x02, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x2b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x1a, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f,
	0x75, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f,
	0x6c, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x49,
	0x6e, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x4f,
	0x75, 0x74, 0x12, 0x34, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x1a,
	0x18, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x75, 0x62, 0x65,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x75, 0x74, 0x12, 0x55, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_infra_proto_rawDescOnce sync.Once
	file_infra_proto_rawDescData = file_infra_proto_rawDesc
)

func file_infra_proto_rawDescGZIP() []byte {
	file_infra_proto_rawDescOnce.Do(func() {
		file_infra_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_proto_rawDescData)
	})
	return file_infra_proto_rawDescData
}

var file_infra_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_infra_proto_goTypes = []interface{}{
	(*GetClusterIn)(nil),                // 0: GetClusterIn
	(*GetClusterOut)(nil),               // 1: GetClusterOut
	(*GetNodepoolIn)(nil),               // 2: GetNodepoolIn
	(*GetNodepoolOut)(nil),              // 3: GetNodepoolOut
	(*ClusterExistsIn)(nil),             // 4: ClusterExistsIn
	(*ClusterExistsOut)(nil),            // 5: ClusterExistsOut
	(*GetClusterKubeconfigOut)(nil),     // 6: GetClusterKubeconfigOut
	(*GetClusterManagedServiceIn)(nil),  // 7: GetClusterManagedServiceIn
	(*GetClusterManagedServiceOut)(nil), // 8: GetClusterManagedServiceOut
}
var file_infra_proto_depIdxs = []int32{
	0, // 0: Infra.GetCluster:input_type -> GetClusterIn
	2, // 1: Infra.GetNodepool:input_type -> GetNodepoolIn
	4, // 2: Infra.ClusterExists:input_type -> ClusterExistsIn
	0, // 3: Infra.GetClusterKubeconfig:input_type -> GetClusterIn
	7, // 4: Infra.GetClusterManagedService:input_type -> GetClusterManagedServiceIn
	1, // 5: Infra.GetCluster:output_type -> GetClusterOut
	3, // 6: Infra.GetNodepool:output_type -> GetNodepoolOut
	5, // 7: Infra.ClusterExists:output_type -> ClusterExistsOut
	6, // 8: Infra.GetClusterKubeconfig:output_type -> GetClusterKubeconfigOut
	8, // 9: Infra.GetClusterManagedService:output_type -> GetClusterManagedServiceOut
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_proto_init() }
func file_infra_proto_init() {
	if File_infra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterIn); i {
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
		file_infra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterOut); i {
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
		file_infra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodepoolIn); i {
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
		file_infra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodepoolOut); i {
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
		file_infra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterExistsIn); i {
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
		file_infra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterExistsOut); i {
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
		file_infra_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterKubeconfigOut); i {
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
		file_infra_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterManagedServiceIn); i {
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
		file_infra_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterManagedServiceOut); i {
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
			RawDescriptor: file_infra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_proto_goTypes,
		DependencyIndexes: file_infra_proto_depIdxs,
		MessageInfos:      file_infra_proto_msgTypes,
	}.Build()
	File_infra_proto = out.File
	file_infra_proto_rawDesc = nil
	file_infra_proto_goTypes = nil
	file_infra_proto_depIdxs = nil
}

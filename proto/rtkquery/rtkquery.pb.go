// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: rtkquery/rtkquery.proto

package rtkquerypb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EndpointType int32

const (
	EndpointType_QUERY    EndpointType = 0
	EndpointType_MUTATION EndpointType = 1
)

// Enum value maps for EndpointType.
var (
	EndpointType_name = map[int32]string{
		0: "QUERY",
		1: "MUTATION",
	}
	EndpointType_value = map[string]int32{
		"QUERY":    0,
		"MUTATION": 1,
	}
)

func (x EndpointType) Enum() *EndpointType {
	p := new(EndpointType)
	*p = x
	return p
}

func (x EndpointType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndpointType) Descriptor() protoreflect.EnumDescriptor {
	return file_rtkquery_rtkquery_proto_enumTypes[0].Descriptor()
}

func (EndpointType) Type() protoreflect.EnumType {
	return &file_rtkquery_rtkquery_proto_enumTypes[0]
}

func (x EndpointType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndpointType.Descriptor instead.
func (EndpointType) EnumDescriptor() ([]byte, []int) {
	return file_rtkquery_rtkquery_proto_rawDescGZIP(), []int{0}
}

type ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ServiceOptions) Reset() {
	*x = ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtkquery_rtkquery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOptions) ProtoMessage() {}

func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_rtkquery_rtkquery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
func (*ServiceOptions) Descriptor() ([]byte, []int) {
	return file_rtkquery_rtkquery_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceOptions) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type MethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *EndpointType `protobuf:"varint,1,opt,name=type,proto3,enum=rtkquery.EndpointType,oneof" json:"type,omitempty"`
	// Types that are assignable to ProvidesTags:
	//
	//	*MethodOptions_ProvidesGeneric
	//	*MethodOptions_ProvidesSpecific
	//	*MethodOptions_ProvidesList
	ProvidesTags isMethodOptions_ProvidesTags `protobuf_oneof:"provides_tags"`
	// Types that are assignable to InvalidatesTags:
	//
	//	*MethodOptions_InvalidatesGeneric
	//	*MethodOptions_InvalidatesSpecific
	//	*MethodOptions_InvalidatesList
	InvalidatesTags isMethodOptions_InvalidatesTags `protobuf_oneof:"invalidates_tags"`
}

func (x *MethodOptions) Reset() {
	*x = MethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtkquery_rtkquery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptions) ProtoMessage() {}

func (x *MethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_rtkquery_rtkquery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptions.ProtoReflect.Descriptor instead.
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return file_rtkquery_rtkquery_proto_rawDescGZIP(), []int{1}
}

func (x *MethodOptions) GetType() EndpointType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return EndpointType_QUERY
}

func (m *MethodOptions) GetProvidesTags() isMethodOptions_ProvidesTags {
	if m != nil {
		return m.ProvidesTags
	}
	return nil
}

func (x *MethodOptions) GetProvidesGeneric() string {
	if x, ok := x.GetProvidesTags().(*MethodOptions_ProvidesGeneric); ok {
		return x.ProvidesGeneric
	}
	return ""
}

func (x *MethodOptions) GetProvidesSpecific() *SpecificTag {
	if x, ok := x.GetProvidesTags().(*MethodOptions_ProvidesSpecific); ok {
		return x.ProvidesSpecific
	}
	return nil
}

func (x *MethodOptions) GetProvidesList() *ListTag {
	if x, ok := x.GetProvidesTags().(*MethodOptions_ProvidesList); ok {
		return x.ProvidesList
	}
	return nil
}

func (m *MethodOptions) GetInvalidatesTags() isMethodOptions_InvalidatesTags {
	if m != nil {
		return m.InvalidatesTags
	}
	return nil
}

func (x *MethodOptions) GetInvalidatesGeneric() string {
	if x, ok := x.GetInvalidatesTags().(*MethodOptions_InvalidatesGeneric); ok {
		return x.InvalidatesGeneric
	}
	return ""
}

func (x *MethodOptions) GetInvalidatesSpecific() *SpecificTag {
	if x, ok := x.GetInvalidatesTags().(*MethodOptions_InvalidatesSpecific); ok {
		return x.InvalidatesSpecific
	}
	return nil
}

func (x *MethodOptions) GetInvalidatesList() string {
	if x, ok := x.GetInvalidatesTags().(*MethodOptions_InvalidatesList); ok {
		return x.InvalidatesList
	}
	return ""
}

type isMethodOptions_ProvidesTags interface {
	isMethodOptions_ProvidesTags()
}

type MethodOptions_ProvidesGeneric struct {
	ProvidesGeneric string `protobuf:"bytes,2,opt,name=provides_generic,json=providesGeneric,proto3,oneof"`
}

type MethodOptions_ProvidesSpecific struct {
	ProvidesSpecific *SpecificTag `protobuf:"bytes,3,opt,name=provides_specific,json=providesSpecific,proto3,oneof"`
}

type MethodOptions_ProvidesList struct {
	ProvidesList *ListTag `protobuf:"bytes,4,opt,name=provides_list,json=providesList,proto3,oneof"`
}

func (*MethodOptions_ProvidesGeneric) isMethodOptions_ProvidesTags() {}

func (*MethodOptions_ProvidesSpecific) isMethodOptions_ProvidesTags() {}

func (*MethodOptions_ProvidesList) isMethodOptions_ProvidesTags() {}

type isMethodOptions_InvalidatesTags interface {
	isMethodOptions_InvalidatesTags()
}

type MethodOptions_InvalidatesGeneric struct {
	InvalidatesGeneric string `protobuf:"bytes,5,opt,name=invalidates_generic,json=invalidatesGeneric,proto3,oneof"`
}

type MethodOptions_InvalidatesSpecific struct {
	InvalidatesSpecific *SpecificTag `protobuf:"bytes,6,opt,name=invalidates_specific,json=invalidatesSpecific,proto3,oneof"`
}

type MethodOptions_InvalidatesList struct {
	InvalidatesList string `protobuf:"bytes,7,opt,name=invalidates_list,json=invalidatesList,proto3,oneof"`
}

func (*MethodOptions_InvalidatesGeneric) isMethodOptions_InvalidatesTags() {}

func (*MethodOptions_InvalidatesSpecific) isMethodOptions_InvalidatesTags() {}

func (*MethodOptions_InvalidatesList) isMethodOptions_InvalidatesTags() {}

type ListTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag   string  `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Items *string `protobuf:"bytes,2,opt,name=items,proto3,oneof" json:"items,omitempty"`
}

func (x *ListTag) Reset() {
	*x = ListTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtkquery_rtkquery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTag) ProtoMessage() {}

func (x *ListTag) ProtoReflect() protoreflect.Message {
	mi := &file_rtkquery_rtkquery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTag.ProtoReflect.Descriptor instead.
func (*ListTag) Descriptor() ([]byte, []int) {
	return file_rtkquery_rtkquery_proto_rawDescGZIP(), []int{2}
}

func (x *ListTag) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ListTag) GetItems() string {
	if x != nil && x.Items != nil {
		return *x.Items
	}
	return ""
}

type SpecificTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string  `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Id  *string `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *SpecificTag) Reset() {
	*x = SpecificTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rtkquery_rtkquery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecificTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecificTag) ProtoMessage() {}

func (x *SpecificTag) ProtoReflect() protoreflect.Message {
	mi := &file_rtkquery_rtkquery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecificTag.ProtoReflect.Descriptor instead.
func (*SpecificTag) Descriptor() ([]byte, []int) {
	return file_rtkquery_rtkquery_proto_rawDescGZIP(), []int{3}
}

func (x *SpecificTag) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *SpecificTag) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

var file_rtkquery_rtkquery_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*ServiceOptions)(nil),
		Field:         66699,
		Name:          "rtkquery.api",
		Tag:           "bytes,66699,opt,name=api",
		Filename:      "rtkquery/rtkquery.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptions)(nil),
		Field:         66699,
		Name:          "rtkquery.endpoint",
		Tag:           "bytes,66699,opt,name=endpoint",
		Filename:      "rtkquery/rtkquery.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional rtkquery.ServiceOptions api = 66699;
	E_Api = &file_rtkquery_rtkquery_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional rtkquery.MethodOptions endpoint = 66699;
	E_Endpoint = &file_rtkquery_rtkquery_proto_extTypes[1]
)

var File_rtkquery_rtkquery_proto protoreflect.FileDescriptor

var file_rtkquery_rtkquery_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x74, 0x6b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x72, 0x74, 0x6b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x74, 0x6b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x24, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xc7, 0x03, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x74, 0x6b, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x02,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x12, 0x44, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x72, 0x74, 0x6b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x61, 0x67, 0x48, 0x00, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x38, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x74, 0x6b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x13, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x12, 0x4a, 0x0a, 0x14, 0x69, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x74, 0x6b, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x61, 0x67, 0x48, 0x01,
	0x52, 0x13, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x40, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x19, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x54, 0x61,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x2a,
	0x27, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x55,
	0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x3a, 0x50, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x8b, 0x89, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x74, 0x6b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x03, 0x61, 0x70, 0x69, 0x88, 0x01, 0x01, 0x3a, 0x58, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8b, 0x89, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x72, 0x74, 0x6b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x6c, 0x69, 0x6f, 0x6e, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x72, 0x74, 0x6b, 0x2d,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x74, 0x6b, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x3b, 0x72, 0x74, 0x6b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rtkquery_rtkquery_proto_rawDescOnce sync.Once
	file_rtkquery_rtkquery_proto_rawDescData = file_rtkquery_rtkquery_proto_rawDesc
)

func file_rtkquery_rtkquery_proto_rawDescGZIP() []byte {
	file_rtkquery_rtkquery_proto_rawDescOnce.Do(func() {
		file_rtkquery_rtkquery_proto_rawDescData = protoimpl.X.CompressGZIP(file_rtkquery_rtkquery_proto_rawDescData)
	})
	return file_rtkquery_rtkquery_proto_rawDescData
}

var file_rtkquery_rtkquery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rtkquery_rtkquery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rtkquery_rtkquery_proto_goTypes = []interface{}{
	(EndpointType)(0),                   // 0: rtkquery.EndpointType
	(*ServiceOptions)(nil),              // 1: rtkquery.ServiceOptions
	(*MethodOptions)(nil),               // 2: rtkquery.MethodOptions
	(*ListTag)(nil),                     // 3: rtkquery.ListTag
	(*SpecificTag)(nil),                 // 4: rtkquery.SpecificTag
	(*descriptorpb.ServiceOptions)(nil), // 5: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 6: google.protobuf.MethodOptions
}
var file_rtkquery_rtkquery_proto_depIdxs = []int32{
	0, // 0: rtkquery.MethodOptions.type:type_name -> rtkquery.EndpointType
	4, // 1: rtkquery.MethodOptions.provides_specific:type_name -> rtkquery.SpecificTag
	3, // 2: rtkquery.MethodOptions.provides_list:type_name -> rtkquery.ListTag
	4, // 3: rtkquery.MethodOptions.invalidates_specific:type_name -> rtkquery.SpecificTag
	5, // 4: rtkquery.api:extendee -> google.protobuf.ServiceOptions
	6, // 5: rtkquery.endpoint:extendee -> google.protobuf.MethodOptions
	1, // 6: rtkquery.api:type_name -> rtkquery.ServiceOptions
	2, // 7: rtkquery.endpoint:type_name -> rtkquery.MethodOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	6, // [6:8] is the sub-list for extension type_name
	4, // [4:6] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rtkquery_rtkquery_proto_init() }
func file_rtkquery_rtkquery_proto_init() {
	if File_rtkquery_rtkquery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rtkquery_rtkquery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOptions); i {
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
		file_rtkquery_rtkquery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptions); i {
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
		file_rtkquery_rtkquery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTag); i {
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
		file_rtkquery_rtkquery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecificTag); i {
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
	file_rtkquery_rtkquery_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MethodOptions_ProvidesGeneric)(nil),
		(*MethodOptions_ProvidesSpecific)(nil),
		(*MethodOptions_ProvidesList)(nil),
		(*MethodOptions_InvalidatesGeneric)(nil),
		(*MethodOptions_InvalidatesSpecific)(nil),
		(*MethodOptions_InvalidatesList)(nil),
	}
	file_rtkquery_rtkquery_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_rtkquery_rtkquery_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rtkquery_rtkquery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_rtkquery_rtkquery_proto_goTypes,
		DependencyIndexes: file_rtkquery_rtkquery_proto_depIdxs,
		EnumInfos:         file_rtkquery_rtkquery_proto_enumTypes,
		MessageInfos:      file_rtkquery_rtkquery_proto_msgTypes,
		ExtensionInfos:    file_rtkquery_rtkquery_proto_extTypes,
	}.Build()
	File_rtkquery_rtkquery_proto = out.File
	file_rtkquery_rtkquery_proto_rawDesc = nil
	file_rtkquery_rtkquery_proto_goTypes = nil
	file_rtkquery_rtkquery_proto_depIdxs = nil
}

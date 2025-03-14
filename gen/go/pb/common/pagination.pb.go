// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: common/pagination.proto

package common

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

type PaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor *string `protobuf:"bytes,3,opt,name=cursor,proto3,oneof" json:"cursor,omitempty"`
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_common_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_common_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PaginationRequest) GetCursor() string {
	if x != nil && x.Cursor != nil {
		return *x.Cursor
	}
	return ""
}

type PaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Total      int64   `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	NextPage   *int32  `protobuf:"varint,4,opt,name=next_page,json=nextPage,proto3,oneof" json:"next_page,omitempty"`
	NextCursor *string `protobuf:"bytes,5,opt,name=next_cursor,json=nextCursor,proto3,oneof" json:"next_cursor,omitempty"`
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_common_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_common_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PaginationResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PaginationResponse) GetNextPage() int32 {
	if x != nil && x.NextPage != nil {
		return *x.NextPage
	}
	return 0
}

func (x *PaginationResponse) GetNextCursor() string {
	if x != nil && x.NextCursor != nil {
		return *x.NextCursor
	}
	return ""
}

var File_common_pagination_proto protoreflect.FileDescriptor

var file_common_pagination_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x22, 0x65, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x1b, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xba, 0x01, 0x0a, 0x12, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x20, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x42, 0x8f, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x73, 0x68,
	0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xca, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xe2, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_pagination_proto_rawDescOnce sync.Once
	file_common_pagination_proto_rawDescData = file_common_pagination_proto_rawDesc
)

func file_common_pagination_proto_rawDescGZIP() []byte {
	file_common_pagination_proto_rawDescOnce.Do(func() {
		file_common_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pagination_proto_rawDescData)
	})
	return file_common_pagination_proto_rawDescData
}

var file_common_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_pagination_proto_goTypes = []any{
	(*PaginationRequest)(nil),  // 0: common.PaginationRequest
	(*PaginationResponse)(nil), // 1: common.PaginationResponse
}
var file_common_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_pagination_proto_init() }
func file_common_pagination_proto_init() {
	if File_common_pagination_proto != nil {
		return
	}
	file_common_pagination_proto_msgTypes[0].OneofWrappers = []any{}
	file_common_pagination_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pagination_proto_goTypes,
		DependencyIndexes: file_common_pagination_proto_depIdxs,
		MessageInfos:      file_common_pagination_proto_msgTypes,
	}.Build()
	File_common_pagination_proto = out.File
	file_common_pagination_proto_rawDesc = nil
	file_common_pagination_proto_goTypes = nil
	file_common_pagination_proto_depIdxs = nil
}

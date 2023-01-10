// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: merge.proto

package rpc

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

type MergeBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *WriteRequest `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// head_branch is the source branch we want to merge
	HeadBranch string `protobuf:"bytes,2,opt,name=head_branch,json=headBranch,proto3" json:"head_branch,omitempty"`
	// branch is the branch into which the given commit shall be merged and whose
	// reference is going to be updated.
	Branch string `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	// title is the title to use for the merge commit.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// message is the message to use for the merge commit.
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	// force merge
	Force bool `protobuf:"varint,6,opt,name=force,proto3" json:"force,omitempty"`
	// delete branch after merge
	Delete bool `protobuf:"varint,7,opt,name=delete,proto3" json:"delete,omitempty"`
}

func (x *MergeBranchRequest) Reset() {
	*x = MergeBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeBranchRequest) ProtoMessage() {}

func (x *MergeBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeBranchRequest.ProtoReflect.Descriptor instead.
func (*MergeBranchRequest) Descriptor() ([]byte, []int) {
	return file_merge_proto_rawDescGZIP(), []int{0}
}

func (x *MergeBranchRequest) GetBase() *WriteRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *MergeBranchRequest) GetHeadBranch() string {
	if x != nil {
		return x.HeadBranch
	}
	return ""
}

func (x *MergeBranchRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *MergeBranchRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MergeBranchRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MergeBranchRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

func (x *MergeBranchRequest) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

// This comment is left unintentionally blank.
type MergeBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The merge commit the branch will be updated to. The caller can still abort the merge.
	CommitId string `protobuf:"bytes,1,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
}

func (x *MergeBranchResponse) Reset() {
	*x = MergeBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeBranchResponse) ProtoMessage() {}

func (x *MergeBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeBranchResponse.ProtoReflect.Descriptor instead.
func (*MergeBranchResponse) Descriptor() ([]byte, []int) {
	return file_merge_proto_rawDescGZIP(), []int{1}
}

func (x *MergeBranchResponse) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

var File_merge_proto protoreflect.FileDescriptor

var file_merge_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72,
	0x70, 0x63, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd2, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x32, 0x0a, 0x13, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x32, 0x52, 0x0a, 0x0c, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6e,
	0x65, 0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_merge_proto_rawDescOnce sync.Once
	file_merge_proto_rawDescData = file_merge_proto_rawDesc
)

func file_merge_proto_rawDescGZIP() []byte {
	file_merge_proto_rawDescOnce.Do(func() {
		file_merge_proto_rawDescData = protoimpl.X.CompressGZIP(file_merge_proto_rawDescData)
	})
	return file_merge_proto_rawDescData
}

var file_merge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_merge_proto_goTypes = []interface{}{
	(*MergeBranchRequest)(nil),  // 0: rpc.MergeBranchRequest
	(*MergeBranchResponse)(nil), // 1: rpc.MergeBranchResponse
	(*WriteRequest)(nil),        // 2: rpc.WriteRequest
}
var file_merge_proto_depIdxs = []int32{
	2, // 0: rpc.MergeBranchRequest.base:type_name -> rpc.WriteRequest
	0, // 1: rpc.MergeService.MergeBranch:input_type -> rpc.MergeBranchRequest
	1, // 2: rpc.MergeService.MergeBranch:output_type -> rpc.MergeBranchResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_merge_proto_init() }
func file_merge_proto_init() {
	if File_merge_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_merge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeBranchRequest); i {
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
		file_merge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeBranchResponse); i {
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
			RawDescriptor: file_merge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_merge_proto_goTypes,
		DependencyIndexes: file_merge_proto_depIdxs,
		MessageInfos:      file_merge_proto_msgTypes,
	}.Build()
	File_merge_proto = out.File
	file_merge_proto_rawDesc = nil
	file_merge_proto_goTypes = nil
	file_merge_proto_depIdxs = nil
}

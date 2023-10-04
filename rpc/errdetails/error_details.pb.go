// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: github.com/piotrostr/atlas-app-toolkit/rpc/errdetails/error_details.proto

package errdetails

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

// TargetInfo is a default representation of error details that conforms
// REST API Syntax Specification
type TargetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code is an enumerated error code,
	// which should be an enum value of [google.rpc.Code][google.rpc.Code]
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// The message is a human-readable non-localized message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The target is a resource name
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *TargetInfo) Reset() {
	*x = TargetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetInfo) ProtoMessage() {}

func (x *TargetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetInfo.ProtoReflect.Descriptor instead.
func (*TargetInfo) Descriptor() ([]byte, []int) {
	return file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescGZIP(), []int{0}
}

func (x *TargetInfo) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TargetInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TargetInfo) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

var File_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto protoreflect.FileDescriptor

var file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2d,
	0x61, 0x70, 0x70, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x65, 0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x22, 0x52, 0x0a, 0x0a, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x45, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2d, 0x61,
	0x70, 0x70, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65,
	0x72, 0x72, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescOnce sync.Once
	file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescData = file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDesc
)

func file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescGZIP() []byte {
	file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescOnce.Do(func() {
		file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescData)
	})
	return file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDescData
}

var file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_goTypes = []interface{}{
	(*TargetInfo)(nil), // 0: atlas.rpc.TargetInfo
}
var file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_init() }
func file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_init() {
	if File_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetInfo); i {
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
			RawDescriptor: file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_goTypes,
		DependencyIndexes: file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_depIdxs,
		MessageInfos:      file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_msgTypes,
	}.Build()
	File_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto = out.File
	file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_rawDesc = nil
	file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_goTypes = nil
	file_github_com_infobloxopen_atlas_app_toolkit_rpc_errdetails_error_details_proto_depIdxs = nil
}

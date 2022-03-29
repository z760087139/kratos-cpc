// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.0--rc2
// source: importTemplate/v1/policy.proto

//包名为小写，并且同目录结构一致，例如：my/package/v1/

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ImportTemplateStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Yaml string `protobuf:"bytes,3,opt,name=yaml,proto3" json:"yaml,omitempty"`
}

func (x *ImportTemplateStruct) Reset() {
	*x = ImportTemplateStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_importTemplate_v1_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportTemplateStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportTemplateStruct) ProtoMessage() {}

func (x *ImportTemplateStruct) ProtoReflect() protoreflect.Message {
	mi := &file_importTemplate_v1_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportTemplateStruct.ProtoReflect.Descriptor instead.
func (*ImportTemplateStruct) Descriptor() ([]byte, []int) {
	return file_importTemplate_v1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ImportTemplateStruct) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ImportTemplateStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportTemplateStruct) GetYaml() string {
	if x != nil {
		return x.Yaml
	}
	return ""
}

type GetTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTemplateRequest) Reset() {
	*x = GetTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_importTemplate_v1_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateRequest) ProtoMessage() {}

func (x *GetTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_importTemplate_v1_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateRequest) Descriptor() ([]byte, []int) {
	return file_importTemplate_v1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *GetTemplateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_importTemplate_v1_policy_proto protoreflect.FileDescriptor

var file_importTemplate_v1_policy_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x14, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x79, 0x61, 0x6d, 0x6c, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x97, 0x01, 0x0a, 0x0e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x84,
	0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x29, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x29, 0x5a, 0x27, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_importTemplate_v1_policy_proto_rawDescOnce sync.Once
	file_importTemplate_v1_policy_proto_rawDescData = file_importTemplate_v1_policy_proto_rawDesc
)

func file_importTemplate_v1_policy_proto_rawDescGZIP() []byte {
	file_importTemplate_v1_policy_proto_rawDescOnce.Do(func() {
		file_importTemplate_v1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_importTemplate_v1_policy_proto_rawDescData)
	})
	return file_importTemplate_v1_policy_proto_rawDescData
}

var file_importTemplate_v1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_importTemplate_v1_policy_proto_goTypes = []interface{}{
	(*ImportTemplateStruct)(nil), // 0: api.importTemplate.v1.ImportTemplateStruct
	(*GetTemplateRequest)(nil),   // 1: api.importTemplate.v1.GetTemplateRequest
}
var file_importTemplate_v1_policy_proto_depIdxs = []int32{
	1, // 0: api.importTemplate.v1.ImportTemplate.GetPolicy:input_type -> api.importTemplate.v1.GetTemplateRequest
	0, // 1: api.importTemplate.v1.ImportTemplate.GetPolicy:output_type -> api.importTemplate.v1.ImportTemplateStruct
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_importTemplate_v1_policy_proto_init() }
func file_importTemplate_v1_policy_proto_init() {
	if File_importTemplate_v1_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_importTemplate_v1_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportTemplateStruct); i {
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
		file_importTemplate_v1_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateRequest); i {
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
			RawDescriptor: file_importTemplate_v1_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_importTemplate_v1_policy_proto_goTypes,
		DependencyIndexes: file_importTemplate_v1_policy_proto_depIdxs,
		MessageInfos:      file_importTemplate_v1_policy_proto_msgTypes,
	}.Build()
	File_importTemplate_v1_policy_proto = out.File
	file_importTemplate_v1_policy_proto_rawDesc = nil
	file_importTemplate_v1_policy_proto_goTypes = nil
	file_importTemplate_v1_policy_proto_depIdxs = nil
}
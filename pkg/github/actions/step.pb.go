// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: pkg/github/actions/step.proto

package actions

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

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shell      string            `protobuf:"bytes,1,opt,name=shell,proto3" json:"shell,omitempty"`
	If         string            `protobuf:"bytes,2,opt,name=if,proto3" json:"if,omitempty"`
	Name       string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Id         string            `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Env        map[string]string `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkingDir string            `protobuf:"bytes,6,opt,name=workingDir,proto3" json:"workingDir,omitempty"`
	Uses       string            `protobuf:"bytes,7,opt,name=uses,proto3" json:"uses,omitempty"`
	With       map[string]string `protobuf:"bytes,8,rep,name=with,proto3" json:"with,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Run        string            `protobuf:"bytes,9,opt,name=run,proto3" json:"run,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_github_actions_step_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_github_actions_step_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_pkg_github_actions_step_proto_rawDescGZIP(), []int{0}
}

func (x *Step) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

func (x *Step) GetIf() string {
	if x != nil {
		return x.If
	}
	return ""
}

func (x *Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Step) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Step) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Step) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *Step) GetUses() string {
	if x != nil {
		return x.Uses
	}
	return ""
}

func (x *Step) GetWith() map[string]string {
	if x != nil {
		return x.With
	}
	return nil
}

func (x *Step) GetRun() string {
	if x != nil {
		return x.Run
	}
	return ""
}

var File_pkg_github_actions_step_proto protoreflect.FileDescriptor

var file_pkg_github_actions_step_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xf8, 0x02, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x68, 0x65, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x65,
	0x70, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12,
	0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x04, 0x77, 0x69, 0x74, 0x68, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x2e, 0x57, 0x69,
	0x74, 0x68, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x77, 0x69, 0x74, 0x68, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x75, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x75, 0x6e, 0x1a,
	0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x57, 0x69, 0x74, 0x68, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x72, 0x61, 0x6e, 0x74, 0x6a, 0x63, 0x2f, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_github_actions_step_proto_rawDescOnce sync.Once
	file_pkg_github_actions_step_proto_rawDescData = file_pkg_github_actions_step_proto_rawDesc
)

func file_pkg_github_actions_step_proto_rawDescGZIP() []byte {
	file_pkg_github_actions_step_proto_rawDescOnce.Do(func() {
		file_pkg_github_actions_step_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_github_actions_step_proto_rawDescData)
	})
	return file_pkg_github_actions_step_proto_rawDescData
}

var file_pkg_github_actions_step_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_github_actions_step_proto_goTypes = []interface{}{
	(*Step)(nil), // 0: forge.github.actions.Step
	nil,          // 1: forge.github.actions.Step.EnvEntry
	nil,          // 2: forge.github.actions.Step.WithEntry
}
var file_pkg_github_actions_step_proto_depIdxs = []int32{
	1, // 0: forge.github.actions.Step.env:type_name -> forge.github.actions.Step.EnvEntry
	2, // 1: forge.github.actions.Step.with:type_name -> forge.github.actions.Step.WithEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_github_actions_step_proto_init() }
func file_pkg_github_actions_step_proto_init() {
	if File_pkg_github_actions_step_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_github_actions_step_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
			RawDescriptor: file_pkg_github_actions_step_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_github_actions_step_proto_goTypes,
		DependencyIndexes: file_pkg_github_actions_step_proto_depIdxs,
		MessageInfos:      file_pkg_github_actions_step_proto_msgTypes,
	}.Build()
	File_pkg_github_actions_step_proto = out.File
	file_pkg_github_actions_step_proto_rawDesc = nil
	file_pkg_github_actions_step_proto_goTypes = nil
	file_pkg_github_actions_step_proto_depIdxs = nil
}

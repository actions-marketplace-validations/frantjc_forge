// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: container_config.proto

package forge

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

type ContainerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entrypoint []string `protobuf:"bytes,1,rep,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	Cmd        []string `protobuf:"bytes,2,rep,name=cmd,proto3" json:"cmd,omitempty"`
	WorkingDir string   `protobuf:"bytes,3,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	Env        []string `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty"`
	User       string   `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Privileged bool     `protobuf:"varint,6,opt,name=privileged,proto3" json:"privileged,omitempty"`
	Mounts     []*Mount `protobuf:"bytes,7,rep,name=mounts,proto3" json:"mounts,omitempty"`
}

func (x *ContainerConfig) Reset() {
	*x = ContainerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerConfig) ProtoMessage() {}

func (x *ContainerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_container_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerConfig.ProtoReflect.Descriptor instead.
func (*ContainerConfig) Descriptor() ([]byte, []int) {
	return file_container_config_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerConfig) GetEntrypoint() []string {
	if x != nil {
		return x.Entrypoint
	}
	return nil
}

func (x *ContainerConfig) GetCmd() []string {
	if x != nil {
		return x.Cmd
	}
	return nil
}

func (x *ContainerConfig) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

func (x *ContainerConfig) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *ContainerConfig) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ContainerConfig) GetPrivileged() bool {
	if x != nil {
		return x.Privileged
	}
	return false
}

func (x *ContainerConfig) GetMounts() []*Mount {
	if x != nil {
		return x.Mounts
	}
	return nil
}

var File_container_config_proto protoreflect.FileDescriptor

var file_container_config_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x1a,
	0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x6d, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x44, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x66, 0x6f, 0x72, 0x67,
	0x65, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42,
	0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72,
	0x61, 0x6e, 0x74, 0x6a, 0x63, 0x2f, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_container_config_proto_rawDescOnce sync.Once
	file_container_config_proto_rawDescData = file_container_config_proto_rawDesc
)

func file_container_config_proto_rawDescGZIP() []byte {
	file_container_config_proto_rawDescOnce.Do(func() {
		file_container_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_container_config_proto_rawDescData)
	})
	return file_container_config_proto_rawDescData
}

var file_container_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_container_config_proto_goTypes = []interface{}{
	(*ContainerConfig)(nil), // 0: forge.ContainerConfig
	(*Mount)(nil),           // 1: forge.Mount
}
var file_container_config_proto_depIdxs = []int32{
	1, // 0: forge.ContainerConfig.mounts:type_name -> forge.Mount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_container_config_proto_init() }
func file_container_config_proto_init() {
	if File_container_config_proto != nil {
		return
	}
	file_mount_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_container_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerConfig); i {
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
			RawDescriptor: file_container_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_container_config_proto_goTypes,
		DependencyIndexes: file_container_config_proto_depIdxs,
		MessageInfos:      file_container_config_proto_msgTypes,
	}.Build()
	File_container_config_proto = out.File
	file_container_config_proto_rawDesc = nil
	file_container_config_proto_goTypes = nil
	file_container_config_proto_depIdxs = nil
}

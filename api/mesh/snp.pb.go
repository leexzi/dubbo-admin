// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: snp.proto

package mesh

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

// When dubbo provider start up, it reports its applicationName and its interfaceName,
// and Dubbo consumer will get the service name mapping info by xDS.
type ServiceMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is namespace of proxyless dubbo server
	Namespace       string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ApplicationName string   `protobuf:"bytes,2,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	InterfaceNames  []string `protobuf:"bytes,3,rep,name=interfaceNames,proto3" json:"interfaceNames,omitempty"`
}

func (x *ServiceMappingRequest) Reset() {
	*x = ServiceMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMappingRequest) ProtoMessage() {}

func (x *ServiceMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_snp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMappingRequest.ProtoReflect.Descriptor instead.
func (*ServiceMappingRequest) Descriptor() ([]byte, []int) {
	return file_snp_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceMappingRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ServiceMappingRequest) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *ServiceMappingRequest) GetInterfaceNames() []string {
	if x != nil {
		return x.InterfaceNames
	}
	return nil
}

type ServiceMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceMappingResponse) Reset() {
	*x = ServiceMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMappingResponse) ProtoMessage() {}

func (x *ServiceMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMappingResponse.ProtoReflect.Descriptor instead.
func (*ServiceMappingResponse) Descriptor() ([]byte, []int) {
	return file_snp_proto_rawDescGZIP(), []int{1}
}

type ServiceMappingXdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is namespace of proxyless dubbo server
	Namespace        string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	InterfaceName    string   `protobuf:"bytes,2,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	ApplicationNames []string `protobuf:"bytes,3,rep,name=applicationNames,proto3" json:"applicationNames,omitempty"`
}

func (x *ServiceMappingXdsResponse) Reset() {
	*x = ServiceMappingXdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_snp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMappingXdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMappingXdsResponse) ProtoMessage() {}

func (x *ServiceMappingXdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_snp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMappingXdsResponse.ProtoReflect.Descriptor instead.
func (*ServiceMappingXdsResponse) Descriptor() ([]byte, []int) {
	return file_snp_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceMappingXdsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ServiceMappingXdsResponse) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *ServiceMappingXdsResponse) GetApplicationNames() []string {
	if x != nil {
		return x.ApplicationNames
	}
	return nil
}

var File_snp_proto protoreflect.FileDescriptor

var file_snp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6e, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x18, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x19, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x58, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x7b, 0x0a, 0x19, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x19, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x70, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_snp_proto_rawDescOnce sync.Once
	file_snp_proto_rawDescData = file_snp_proto_rawDesc
)

func file_snp_proto_rawDescGZIP() []byte {
	file_snp_proto_rawDescOnce.Do(func() {
		file_snp_proto_rawDescData = protoimpl.X.CompressGZIP(file_snp_proto_rawDescData)
	})
	return file_snp_proto_rawDescData
}

var file_snp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_snp_proto_goTypes = []interface{}{
	(*ServiceMappingRequest)(nil),     // 0: v1alpha1.ServiceMappingRequest
	(*ServiceMappingResponse)(nil),    // 1: v1alpha1.ServiceMappingResponse
	(*ServiceMappingXdsResponse)(nil), // 2: v1alpha1.ServiceMappingXdsResponse
}
var file_snp_proto_depIdxs = []int32{
	0, // 0: v1alpha1.ServiceNameMappingService.registerServiceAppMapping:input_type -> v1alpha1.ServiceMappingRequest
	1, // 1: v1alpha1.ServiceNameMappingService.registerServiceAppMapping:output_type -> v1alpha1.ServiceMappingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_snp_proto_init() }
func file_snp_proto_init() {
	if File_snp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_snp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMappingRequest); i {
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
		file_snp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMappingResponse); i {
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
		file_snp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMappingXdsResponse); i {
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
			RawDescriptor: file_snp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_snp_proto_goTypes,
		DependencyIndexes: file_snp_proto_depIdxs,
		MessageInfos:      file_snp_proto_msgTypes,
	}.Build()
	File_snp_proto = out.File
	file_snp_proto_rawDesc = nil
	file_snp_proto_goTypes = nil
	file_snp_proto_depIdxs = nil
}

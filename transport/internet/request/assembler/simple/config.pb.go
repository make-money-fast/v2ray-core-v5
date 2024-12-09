package simple

import (
	_ "github.com/make-money-fast/v2ray-core-v5/common/protoext"
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

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxWriteSize             int32   `protobuf:"varint,1,opt,name=max_write_size,json=maxWriteSize,proto3" json:"max_write_size,omitempty"`
	WaitSubsequentWriteMs    int32   `protobuf:"varint,2,opt,name=wait_subsequent_write_ms,json=waitSubsequentWriteMs,proto3" json:"wait_subsequent_write_ms,omitempty"`
	InitialPollingIntervalMs int32   `protobuf:"varint,3,opt,name=initial_polling_interval_ms,json=initialPollingIntervalMs,proto3" json:"initial_polling_interval_ms,omitempty"`
	MaxPollingIntervalMs     int32   `protobuf:"varint,4,opt,name=max_polling_interval_ms,json=maxPollingIntervalMs,proto3" json:"max_polling_interval_ms,omitempty"`
	MinPollingIntervalMs     int32   `protobuf:"varint,5,opt,name=min_polling_interval_ms,json=minPollingIntervalMs,proto3" json:"min_polling_interval_ms,omitempty"`
	BackoffFactor            float32 `protobuf:"fixed32,6,opt,name=backoff_factor,json=backoffFactor,proto3" json:"backoff_factor,omitempty"`
	FailedRetryIntervalMs    int32   `protobuf:"varint,7,opt,name=failed_retry_interval_ms,json=failedRetryIntervalMs,proto3" json:"failed_retry_interval_ms,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_request_assembler_simple_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_request_assembler_simple_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_request_assembler_simple_config_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConfig) GetMaxWriteSize() int32 {
	if x != nil {
		return x.MaxWriteSize
	}
	return 0
}

func (x *ClientConfig) GetWaitSubsequentWriteMs() int32 {
	if x != nil {
		return x.WaitSubsequentWriteMs
	}
	return 0
}

func (x *ClientConfig) GetInitialPollingIntervalMs() int32 {
	if x != nil {
		return x.InitialPollingIntervalMs
	}
	return 0
}

func (x *ClientConfig) GetMaxPollingIntervalMs() int32 {
	if x != nil {
		return x.MaxPollingIntervalMs
	}
	return 0
}

func (x *ClientConfig) GetMinPollingIntervalMs() int32 {
	if x != nil {
		return x.MinPollingIntervalMs
	}
	return 0
}

func (x *ClientConfig) GetBackoffFactor() float32 {
	if x != nil {
		return x.BackoffFactor
	}
	return 0
}

func (x *ClientConfig) GetFailedRetryIntervalMs() int32 {
	if x != nil {
		return x.FailedRetryIntervalMs
	}
	return 0
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxWriteSize int32 `protobuf:"varint,1,opt,name=max_write_size,json=maxWriteSize,proto3" json:"max_write_size,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_request_assembler_simple_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_request_assembler_simple_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_request_assembler_simple_config_proto_rawDescGZIP(), []int{1}
}

func (x *ServerConfig) GetMaxWriteSize() int32 {
	if x != nil {
		return x.MaxWriteSize
	}
	return 0
}

var File_transport_internet_request_assembler_simple_config_proto protoreflect.FileDescriptor

var file_transport_internet_request_assembler_simple_config_proto_rawDesc = []byte{
	0x0a, 0x38, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x6d, 0x62, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x03, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d,
	0x61, 0x78, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x77,
	0x61, 0x69, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x5f, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x77,
	0x61, 0x69, 0x74, 0x53, 0x75, 0x62, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x4d, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x50, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x4d, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x69,
	0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6d, 0x69, 0x6e,
	0x50, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x6f,
	0x66, 0x66, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x18, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d,
	0x73, 0x3a, 0x30, 0x82, 0xb5, 0x18, 0x2c, 0x0a, 0x22, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x6d,
	0x62, 0x6c, 0x65, 0x72, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x22, 0x66, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x3a, 0x30, 0x82, 0xb5, 0x18, 0x2c, 0x0a,
	0x22, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0xc3, 0x01, 0x0a, 0x3a,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62,
	0x6c, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x65,
	0x72, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0xaa, 0x02, 0x36, 0x56, 0x32, 0x52, 0x61, 0x79,
	0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x6d, 0x62, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_request_assembler_simple_config_proto_rawDescOnce sync.Once
	file_transport_internet_request_assembler_simple_config_proto_rawDescData = file_transport_internet_request_assembler_simple_config_proto_rawDesc
)

func file_transport_internet_request_assembler_simple_config_proto_rawDescGZIP() []byte {
	file_transport_internet_request_assembler_simple_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_request_assembler_simple_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_request_assembler_simple_config_proto_rawDescData)
	})
	return file_transport_internet_request_assembler_simple_config_proto_rawDescData
}

var file_transport_internet_request_assembler_simple_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transport_internet_request_assembler_simple_config_proto_goTypes = []interface{}{
	(*ClientConfig)(nil), // 0: v2ray.core.transport.internet.request.assembler.simple.ClientConfig
	(*ServerConfig)(nil), // 1: v2ray.core.transport.internet.request.assembler.simple.ServerConfig
}
var file_transport_internet_request_assembler_simple_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_internet_request_assembler_simple_config_proto_init() }
func file_transport_internet_request_assembler_simple_config_proto_init() {
	if File_transport_internet_request_assembler_simple_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_request_assembler_simple_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
		file_transport_internet_request_assembler_simple_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
			RawDescriptor: file_transport_internet_request_assembler_simple_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_request_assembler_simple_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_request_assembler_simple_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_request_assembler_simple_config_proto_msgTypes,
	}.Build()
	File_transport_internet_request_assembler_simple_config_proto = out.File
	file_transport_internet_request_assembler_simple_config_proto_rawDesc = nil
	file_transport_internet_request_assembler_simple_config_proto_goTypes = nil
	file_transport_internet_request_assembler_simple_config_proto_depIdxs = nil
}

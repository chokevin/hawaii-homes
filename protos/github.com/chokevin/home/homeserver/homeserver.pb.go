// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: homeserver.proto

package homeserver

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

// The request message containing the property ID.
type GetPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPropertyRequest) Reset() {
	*x = GetPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homeserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPropertyRequest) ProtoMessage() {}

func (x *GetPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_homeserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPropertyRequest.ProtoReflect.Descriptor instead.
func (*GetPropertyRequest) Descriptor() ([]byte, []int) {
	return file_homeserver_proto_rawDescGZIP(), []int{0}
}

func (x *GetPropertyRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The response message containing property details.
type GetPropertyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	City    string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	State   string  `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	ZipCode string  `protobuf:"bytes,5,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Price   float32 `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetPropertyResponse) Reset() {
	*x = GetPropertyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homeserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPropertyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPropertyResponse) ProtoMessage() {}

func (x *GetPropertyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_homeserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPropertyResponse.ProtoReflect.Descriptor instead.
func (*GetPropertyResponse) Descriptor() ([]byte, []int) {
	return file_homeserver_proto_rawDescGZIP(), []int{1}
}

func (x *GetPropertyResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetPropertyResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetPropertyResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetPropertyResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetPropertyResponse) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *GetPropertyResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_homeserver_proto protoreflect.FileDescriptor

var file_homeserver_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x24,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x32, 0x5d, 0x0a, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12,
	0x1e, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x6f, 0x6b, 0x65, 0x76, 0x69, 0x6e, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x68, 0x6f, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_homeserver_proto_rawDescOnce sync.Once
	file_homeserver_proto_rawDescData = file_homeserver_proto_rawDesc
)

func file_homeserver_proto_rawDescGZIP() []byte {
	file_homeserver_proto_rawDescOnce.Do(func() {
		file_homeserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_homeserver_proto_rawDescData)
	})
	return file_homeserver_proto_rawDescData
}

var file_homeserver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_homeserver_proto_goTypes = []any{
	(*GetPropertyRequest)(nil),  // 0: homeserver.GetPropertyRequest
	(*GetPropertyResponse)(nil), // 1: homeserver.GetPropertyResponse
}
var file_homeserver_proto_depIdxs = []int32{
	0, // 0: homeserver.HomeService.GetProperty:input_type -> homeserver.GetPropertyRequest
	1, // 1: homeserver.HomeService.GetProperty:output_type -> homeserver.GetPropertyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_homeserver_proto_init() }
func file_homeserver_proto_init() {
	if File_homeserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_homeserver_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPropertyRequest); i {
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
		file_homeserver_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetPropertyResponse); i {
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
			RawDescriptor: file_homeserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_homeserver_proto_goTypes,
		DependencyIndexes: file_homeserver_proto_depIdxs,
		MessageInfos:      file_homeserver_proto_msgTypes,
	}.Build()
	File_homeserver_proto = out.File
	file_homeserver_proto_rawDesc = nil
	file_homeserver_proto_goTypes = nil
	file_homeserver_proto_depIdxs = nil
}

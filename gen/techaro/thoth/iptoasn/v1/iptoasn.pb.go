// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: techaro/thoth/iptoasn/v1/iptoasn.proto

package iptoasnv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LookupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IpAddress     string                 `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	mi := &file_techaro_thoth_iptoasn_v1_iptoasn_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_techaro_thoth_iptoasn_v1_iptoasn_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescGZIP(), []int{0}
}

func (x *LookupRequest) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type LookupResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Announced     bool                   `protobuf:"varint,1,opt,name=announced,proto3" json:"announced,omitempty"`
	AsNumber      uint32                 `protobuf:"varint,2,opt,name=as_number,json=asNumber,proto3" json:"as_number,omitempty"`
	Cidr          []string               `protobuf:"bytes,3,rep,name=cidr,proto3" json:"cidr,omitempty"`
	CountryCode   string                 `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LookupResponse) Reset() {
	*x = LookupResponse{}
	mi := &file_techaro_thoth_iptoasn_v1_iptoasn_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupResponse) ProtoMessage() {}

func (x *LookupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_techaro_thoth_iptoasn_v1_iptoasn_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupResponse.ProtoReflect.Descriptor instead.
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescGZIP(), []int{1}
}

func (x *LookupResponse) GetAnnounced() bool {
	if x != nil {
		return x.Announced
	}
	return false
}

func (x *LookupResponse) GetAsNumber() uint32 {
	if x != nil {
		return x.AsNumber
	}
	return 0
}

func (x *LookupResponse) GetCidr() []string {
	if x != nil {
		return x.Cidr
	}
	return nil
}

func (x *LookupResponse) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *LookupResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_techaro_thoth_iptoasn_v1_iptoasn_proto protoreflect.FileDescriptor

const file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDesc = "" +
	"\n" +
	"&techaro/thoth/iptoasn/v1/iptoasn.proto\x12\x18techaro.thoth.iptoasn.v1\x1a\x1bbuf/validate/validate.proto\"\x82\x01\n" +
	"\rLookupRequest\x12q\n" +
	"\n" +
	"ip_address\x18\x01 \x01(\tBR\xbaHO\xba\x01I\n" +
	"\x18ip_address.is_ip_address\x12 ip_address must be an IP address\x1a\vthis.isIp()\xc8\x01\x01R\tipAddress\"\xa4\x01\n" +
	"\x0eLookupResponse\x12\x1c\n" +
	"\tannounced\x18\x01 \x01(\bR\tannounced\x12\x1b\n" +
	"\tas_number\x18\x02 \x01(\rR\basNumber\x12\x12\n" +
	"\x04cidr\x18\x03 \x03(\tR\x04cidr\x12!\n" +
	"\fcountry_code\x18\x04 \x01(\tR\vcountryCode\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription2m\n" +
	"\x0eIpToASNService\x12[\n" +
	"\x06Lookup\x12'.techaro.thoth.iptoasn.v1.LookupRequest\x1a(.techaro.thoth.iptoasn.v1.LookupResponseB\xf8\x01\n" +
	"\x1ccom.techaro.thoth.iptoasn.v1B\fIptoasnProtoP\x01ZGgithub.com/TecharoHQ/thoth-proto/gen/techaro/thoth/iptoasn/v1;iptoasnv1\xa2\x02\x03TTI\xaa\x02\x18Techaro.Thoth.Iptoasn.V1\xca\x02\x18Techaro\\Thoth\\Iptoasn\\V1\xe2\x02$Techaro\\Thoth\\Iptoasn\\V1\\GPBMetadata\xea\x02\x1bTecharo::Thoth::Iptoasn::V1b\x06proto3"

var (
	file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescOnce sync.Once
	file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescData []byte
)

func file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescGZIP() []byte {
	file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescOnce.Do(func() {
		file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDesc), len(file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDesc)))
	})
	return file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDescData
}

var file_techaro_thoth_iptoasn_v1_iptoasn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_techaro_thoth_iptoasn_v1_iptoasn_proto_goTypes = []any{
	(*LookupRequest)(nil),  // 0: techaro.thoth.iptoasn.v1.LookupRequest
	(*LookupResponse)(nil), // 1: techaro.thoth.iptoasn.v1.LookupResponse
}
var file_techaro_thoth_iptoasn_v1_iptoasn_proto_depIdxs = []int32{
	0, // 0: techaro.thoth.iptoasn.v1.IpToASNService.Lookup:input_type -> techaro.thoth.iptoasn.v1.LookupRequest
	1, // 1: techaro.thoth.iptoasn.v1.IpToASNService.Lookup:output_type -> techaro.thoth.iptoasn.v1.LookupResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_techaro_thoth_iptoasn_v1_iptoasn_proto_init() }
func file_techaro_thoth_iptoasn_v1_iptoasn_proto_init() {
	if File_techaro_thoth_iptoasn_v1_iptoasn_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDesc), len(file_techaro_thoth_iptoasn_v1_iptoasn_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_techaro_thoth_iptoasn_v1_iptoasn_proto_goTypes,
		DependencyIndexes: file_techaro_thoth_iptoasn_v1_iptoasn_proto_depIdxs,
		MessageInfos:      file_techaro_thoth_iptoasn_v1_iptoasn_proto_msgTypes,
	}.Build()
	File_techaro_thoth_iptoasn_v1_iptoasn_proto = out.File
	file_techaro_thoth_iptoasn_v1_iptoasn_proto_goTypes = nil
	file_techaro_thoth_iptoasn_v1_iptoasn_proto_depIdxs = nil
}

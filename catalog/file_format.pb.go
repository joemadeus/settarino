// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: file_format.proto

package catalog

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ElementFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*ElementFile_Rstyle
	//	*ElementFile_Dstyle
	Data isElementFile_Data `protobuf_oneof:"Data"`
}

func (x *ElementFile) Reset() {
	*x = ElementFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_format_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElementFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElementFile) ProtoMessage() {}

func (x *ElementFile) ProtoReflect() protoreflect.Message {
	mi := &file_file_format_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElementFile.ProtoReflect.Descriptor instead.
func (*ElementFile) Descriptor() ([]byte, []int) {
	return file_file_format_proto_rawDescGZIP(), []int{0}
}

func (m *ElementFile) GetData() isElementFile_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ElementFile) GetRstyle() *ReplaceStyleElements {
	if x, ok := x.GetData().(*ElementFile_Rstyle); ok {
		return x.Rstyle
	}
	return nil
}

func (x *ElementFile) GetDstyle() *DeltaStyleElements {
	if x, ok := x.GetData().(*ElementFile_Dstyle); ok {
		return x.Dstyle
	}
	return nil
}

type isElementFile_Data interface {
	isElementFile_Data()
}

type ElementFile_Rstyle struct {
	Rstyle *ReplaceStyleElements `protobuf:"bytes,1,opt,name=rstyle,proto3,oneof"`
}

type ElementFile_Dstyle struct {
	Dstyle *DeltaStyleElements `protobuf:"bytes,2,opt,name=dstyle,proto3,oneof"`
}

func (*ElementFile_Rstyle) isElementFile_Data() {}

func (*ElementFile_Dstyle) isElementFile_Data() {}

type ReplaceStyleHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag       string               `protobuf:"bytes,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	WriteTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=WriteTime,proto3" json:"WriteTime,omitempty"`
}

func (x *ReplaceStyleHeader) Reset() {
	*x = ReplaceStyleHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_format_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceStyleHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceStyleHeader) ProtoMessage() {}

func (x *ReplaceStyleHeader) ProtoReflect() protoreflect.Message {
	mi := &file_file_format_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceStyleHeader.ProtoReflect.Descriptor instead.
func (*ReplaceStyleHeader) Descriptor() ([]byte, []int) {
	return file_file_format_proto_rawDescGZIP(), []int{1}
}

func (x *ReplaceStyleHeader) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ReplaceStyleHeader) GetWriteTime() *timestamp.Timestamp {
	if x != nil {
		return x.WriteTime
	}
	return nil
}

type ReplaceStyleElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *ReplaceStyleElement) Reset() {
	*x = ReplaceStyleElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_format_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceStyleElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceStyleElement) ProtoMessage() {}

func (x *ReplaceStyleElement) ProtoReflect() protoreflect.Message {
	mi := &file_file_format_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceStyleElement.ProtoReflect.Descriptor instead.
func (*ReplaceStyleElement) Descriptor() ([]byte, []int) {
	return file_file_format_proto_rawDescGZIP(), []int{2}
}

func (x *ReplaceStyleElement) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ReplaceStyleElements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *ReplaceStyleHeader    `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Elements []*ReplaceStyleElement `protobuf:"bytes,2,rep,name=Elements,proto3" json:"Elements,omitempty"`
}

func (x *ReplaceStyleElements) Reset() {
	*x = ReplaceStyleElements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_format_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceStyleElements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceStyleElements) ProtoMessage() {}

func (x *ReplaceStyleElements) ProtoReflect() protoreflect.Message {
	mi := &file_file_format_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceStyleElements.ProtoReflect.Descriptor instead.
func (*ReplaceStyleElements) Descriptor() ([]byte, []int) {
	return file_file_format_proto_rawDescGZIP(), []int{3}
}

func (x *ReplaceStyleElements) GetHeader() *ReplaceStyleHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ReplaceStyleElements) GetElements() []*ReplaceStyleElement {
	if x != nil {
		return x.Elements
	}
	return nil
}

// TODO: delta-style element files
type DeltaStyleElements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeltaStyleElements) Reset() {
	*x = DeltaStyleElements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_format_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaStyleElements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaStyleElements) ProtoMessage() {}

func (x *DeltaStyleElements) ProtoReflect() protoreflect.Message {
	mi := &file_file_format_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaStyleElements.ProtoReflect.Descriptor instead.
func (*DeltaStyleElements) Descriptor() ([]byte, []int) {
	return file_file_format_proto_rawDescGZIP(), []int{4}
}

var File_file_format_proto protoreflect.FileDescriptor

var file_file_format_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01,
	0x0a, 0x0b, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x72, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53,
	0x74, 0x79, 0x6c, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x00, 0x52, 0x06,
	0x72, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x79, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x48, 0x00, 0x52, 0x06, 0x64, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x60, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x54,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x54, 0x61, 0x67, 0x12, 0x38, 0x0a,
	0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x22, 0x85, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x79, 0x6c,
	0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x38,
	0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x79, 0x74,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x68, 0x65, 0x6c, 0x69,
	0x78, 0x2d, 0x73, 0x65, 0x74, 0x74, 0x61, 0x72, 0x69, 0x6e, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_format_proto_rawDescOnce sync.Once
	file_file_format_proto_rawDescData = file_file_format_proto_rawDesc
)

func file_file_format_proto_rawDescGZIP() []byte {
	file_file_format_proto_rawDescOnce.Do(func() {
		file_file_format_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_format_proto_rawDescData)
	})
	return file_file_format_proto_rawDescData
}

var file_file_format_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_file_format_proto_goTypes = []interface{}{
	(*ElementFile)(nil),          // 0: catalog.ElementFile
	(*ReplaceStyleHeader)(nil),   // 1: catalog.ReplaceStyleHeader
	(*ReplaceStyleElement)(nil),  // 2: catalog.ReplaceStyleElement
	(*ReplaceStyleElements)(nil), // 3: catalog.ReplaceStyleElements
	(*DeltaStyleElements)(nil),   // 4: catalog.DeltaStyleElements
	(*timestamp.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_file_format_proto_depIdxs = []int32{
	3, // 0: catalog.ElementFile.rstyle:type_name -> catalog.ReplaceStyleElements
	4, // 1: catalog.ElementFile.dstyle:type_name -> catalog.DeltaStyleElements
	5, // 2: catalog.ReplaceStyleHeader.WriteTime:type_name -> google.protobuf.Timestamp
	1, // 3: catalog.ReplaceStyleElements.Header:type_name -> catalog.ReplaceStyleHeader
	2, // 4: catalog.ReplaceStyleElements.Elements:type_name -> catalog.ReplaceStyleElement
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_file_format_proto_init() }
func file_file_format_proto_init() {
	if File_file_format_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_format_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElementFile); i {
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
		file_file_format_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceStyleHeader); i {
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
		file_file_format_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceStyleElement); i {
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
		file_file_format_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceStyleElements); i {
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
		file_file_format_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaStyleElements); i {
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
	file_file_format_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ElementFile_Rstyle)(nil),
		(*ElementFile_Dstyle)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_format_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_format_proto_goTypes,
		DependencyIndexes: file_file_format_proto_depIdxs,
		MessageInfos:      file_file_format_proto_msgTypes,
	}.Build()
	File_file_format_proto = out.File
	file_file_format_proto_rawDesc = nil
	file_file_format_proto_goTypes = nil
	file_file_format_proto_depIdxs = nil
}

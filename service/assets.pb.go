// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: assets.proto

package service

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

type GetAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//
	//	*GetAssetsResponse_File
	//	*GetAssetsResponse_Directory
	Body isGetAssetsResponse_Body `protobuf_oneof:"body"`
}

func (x *GetAssetsResponse) Reset() {
	*x = GetAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsResponse) ProtoMessage() {}

func (x *GetAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsResponse.ProtoReflect.Descriptor instead.
func (*GetAssetsResponse) Descriptor() ([]byte, []int) {
	return file_assets_proto_rawDescGZIP(), []int{0}
}

func (m *GetAssetsResponse) GetBody() isGetAssetsResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *GetAssetsResponse) GetFile() *FileResponse {
	if x, ok := x.GetBody().(*GetAssetsResponse_File); ok {
		return x.File
	}
	return nil
}

func (x *GetAssetsResponse) GetDirectory() *DirectoryResponse {
	if x, ok := x.GetBody().(*GetAssetsResponse_Directory); ok {
		return x.Directory
	}
	return nil
}

type isGetAssetsResponse_Body interface {
	isGetAssetsResponse_Body()
}

type GetAssetsResponse_File struct {
	File *FileResponse `protobuf:"bytes,1,opt,name=file,proto3,oneof"`
}

type GetAssetsResponse_Directory struct {
	Directory *DirectoryResponse `protobuf:"bytes,2,opt,name=directory,proto3,oneof"`
}

func (*GetAssetsResponse_File) isGetAssetsResponse_Body() {}

func (*GetAssetsResponse_Directory) isGetAssetsResponse_Body() {}

type GetAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetAssetsRequest) Reset() {
	*x = GetAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetsRequest) ProtoMessage() {}

func (x *GetAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetsRequest.ProtoReflect.Descriptor instead.
func (*GetAssetsRequest) Descriptor() ([]byte, []int) {
	return file_assets_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssetsRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files       []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Directories []string `protobuf:"bytes,2,rep,name=directories,proto3" json:"directories,omitempty"`
}

func (x *DirectoryResponse) Reset() {
	*x = DirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryResponse) ProtoMessage() {}

func (x *DirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryResponse.ProtoReflect.Descriptor instead.
func (*DirectoryResponse) Descriptor() ([]byte, []int) {
	return file_assets_proto_rawDescGZIP(), []int{2}
}

func (x *DirectoryResponse) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DirectoryResponse) GetDirectories() []string {
	if x != nil {
		return x.Directories
	}
	return nil
}

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentLength int32  `protobuf:"varint,1,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	Body          []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_assets_proto_rawDescGZIP(), []int{3}
}

func (x *FileResponse) GetContentLength() int32 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *FileResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Assets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Assets) Reset() {
	*x = Assets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_assets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assets) ProtoMessage() {}

func (x *Assets) ProtoReflect() protoreflect.Message {
	mi := &file_assets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assets.ProtoReflect.Descriptor instead.
func (*Assets) Descriptor() ([]byte, []int) {
	return file_assets_proto_rawDescGZIP(), []int{4}
}

func (x *Assets) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_assets_proto protoreflect.FileDescriptor

var file_assets_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x26,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x4b, 0x0a, 0x11, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x1c,
	0x0a, 0x06, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x5f, 0x0a, 0x1f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42,
	0x12, 0x43, 0x61, 0x72, 0x74, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_assets_proto_rawDescOnce sync.Once
	file_assets_proto_rawDescData = file_assets_proto_rawDesc
)

func file_assets_proto_rawDescGZIP() []byte {
	file_assets_proto_rawDescOnce.Do(func() {
		file_assets_proto_rawDescData = protoimpl.X.CompressGZIP(file_assets_proto_rawDescData)
	})
	return file_assets_proto_rawDescData
}

var file_assets_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_assets_proto_goTypes = []interface{}{
	(*GetAssetsResponse)(nil), // 0: service.GetAssetsResponse
	(*GetAssetsRequest)(nil),  // 1: service.GetAssetsRequest
	(*DirectoryResponse)(nil), // 2: service.DirectoryResponse
	(*FileResponse)(nil),      // 3: service.FileResponse
	(*Assets)(nil),            // 4: service.Assets
}
var file_assets_proto_depIdxs = []int32{
	3, // 0: service.GetAssetsResponse.file:type_name -> service.FileResponse
	2, // 1: service.GetAssetsResponse.directory:type_name -> service.DirectoryResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_assets_proto_init() }
func file_assets_proto_init() {
	if File_assets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_assets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetsResponse); i {
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
		file_assets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetsRequest); i {
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
		file_assets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectoryResponse); i {
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
		file_assets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
		file_assets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assets); i {
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
	file_assets_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetAssetsResponse_File)(nil),
		(*GetAssetsResponse_Directory)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_assets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_assets_proto_goTypes,
		DependencyIndexes: file_assets_proto_depIdxs,
		MessageInfos:      file_assets_proto_msgTypes,
	}.Build()
	File_assets_proto = out.File
	file_assets_proto_rawDesc = nil
	file_assets_proto_goTypes = nil
	file_assets_proto_depIdxs = nil
}

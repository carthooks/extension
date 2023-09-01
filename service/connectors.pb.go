// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: connectors.proto

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

type GetConnectorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connectors []*Connector `protobuf:"bytes,1,rep,name=connectors,proto3" json:"connectors,omitempty"`
}

func (x *GetConnectorsResponse) Reset() {
	*x = GetConnectorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectorsResponse) ProtoMessage() {}

func (x *GetConnectorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectorsResponse.ProtoReflect.Descriptor instead.
func (*GetConnectorsResponse) Descriptor() ([]byte, []int) {
	return file_connectors_proto_rawDescGZIP(), []int{0}
}

func (x *GetConnectorsResponse) GetConnectors() []*Connector {
	if x != nil {
		return x.Connectors
	}
	return nil
}

type Connector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID             string      `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name             *I18NString `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description      *I18NString `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Icon             *Assets     `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	SetupComponent   *Assets     `protobuf:"bytes,5,opt,name=setupComponent,proto3" json:"setupComponent,omitempty"`
	Service_UUIDList []string    `protobuf:"bytes,7,rep,name=service_UUID_list,json=serviceUUIDList,proto3" json:"service_UUID_list,omitempty"`
}

func (x *Connector) Reset() {
	*x = Connector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connector) ProtoMessage() {}

func (x *Connector) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connector.ProtoReflect.Descriptor instead.
func (*Connector) Descriptor() ([]byte, []int) {
	return file_connectors_proto_rawDescGZIP(), []int{1}
}

func (x *Connector) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *Connector) GetName() *I18NString {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Connector) GetDescription() *I18NString {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Connector) GetIcon() *Assets {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *Connector) GetSetupComponent() *Assets {
	if x != nil {
		return x.SetupComponent
	}
	return nil
}

func (x *Connector) GetService_UUIDList() []string {
	if x != nil {
		return x.Service_UUIDList
	}
	return nil
}

type CreateConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId       int64             `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Connector_UUID string            `protobuf:"bytes,2,opt,name=connector_UUID,json=connectorUUID,proto3" json:"connector_UUID,omitempty"`
	Params         map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateConnectionRequest) Reset() {
	*x = CreateConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConnectionRequest) ProtoMessage() {}

func (x *CreateConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateConnectionRequest) Descriptor() ([]byte, []int) {
	return file_connectors_proto_rawDescGZIP(), []int{2}
}

func (x *CreateConnectionRequest) GetTenantId() int64 {
	if x != nil {
		return x.TenantId
	}
	return 0
}

func (x *CreateConnectionRequest) GetConnector_UUID() string {
	if x != nil {
		return x.Connector_UUID
	}
	return ""
}

func (x *CreateConnectionRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type CreateConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ConnectionId string `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateConnectionResponse) Reset() {
	*x = CreateConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConnectionResponse) ProtoMessage() {}

func (x *CreateConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connectors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConnectionResponse.ProtoReflect.Descriptor instead.
func (*CreateConnectionResponse) Descriptor() ([]byte, []int) {
	return file_connectors_proto_rawDescGZIP(), []int{3}
}

func (x *CreateConnectionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateConnectionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateConnectionResponse) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *CreateConnectionResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_connectors_proto protoreflect.FileDescriptor

var file_connectors_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x31, 0x38, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x31, 0x38, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0e, 0x73,
	0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x55, 0x55, 0x49, 0x44, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0xde, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x55, 0x55, 0x49, 0x44,
	0x12, 0x44, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x87, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x49, 0x0a, 0x1d, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectors_proto_rawDescOnce sync.Once
	file_connectors_proto_rawDescData = file_connectors_proto_rawDesc
)

func file_connectors_proto_rawDescGZIP() []byte {
	file_connectors_proto_rawDescOnce.Do(func() {
		file_connectors_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectors_proto_rawDescData)
	})
	return file_connectors_proto_rawDescData
}

var file_connectors_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_connectors_proto_goTypes = []interface{}{
	(*GetConnectorsResponse)(nil),    // 0: service.GetConnectorsResponse
	(*Connector)(nil),                // 1: service.Connector
	(*CreateConnectionRequest)(nil),  // 2: service.CreateConnectionRequest
	(*CreateConnectionResponse)(nil), // 3: service.CreateConnectionResponse
	nil,                              // 4: service.CreateConnectionRequest.ParamsEntry
	(*I18NString)(nil),               // 5: service.I18nString
	(*Assets)(nil),                   // 6: service.Assets
}
var file_connectors_proto_depIdxs = []int32{
	1, // 0: service.GetConnectorsResponse.connectors:type_name -> service.Connector
	5, // 1: service.Connector.name:type_name -> service.I18nString
	5, // 2: service.Connector.description:type_name -> service.I18nString
	6, // 3: service.Connector.icon:type_name -> service.Assets
	6, // 4: service.Connector.setupComponent:type_name -> service.Assets
	4, // 5: service.CreateConnectionRequest.params:type_name -> service.CreateConnectionRequest.ParamsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_connectors_proto_init() }
func file_connectors_proto_init() {
	if File_connectors_proto != nil {
		return
	}
	file_common_proto_init()
	file_assets_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_connectors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectorsResponse); i {
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
		file_connectors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connector); i {
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
		file_connectors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConnectionRequest); i {
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
		file_connectors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConnectionResponse); i {
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
			RawDescriptor: file_connectors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectors_proto_goTypes,
		DependencyIndexes: file_connectors_proto_depIdxs,
		MessageInfos:      file_connectors_proto_msgTypes,
	}.Build()
	File_connectors_proto = out.File
	file_connectors_proto_rawDesc = nil
	file_connectors_proto_goTypes = nil
	file_connectors_proto_depIdxs = nil
}

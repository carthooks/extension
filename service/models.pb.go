// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: models.proto

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

type FieldType int32

const (
	FieldType_TYPE_UNSPECIFIED FieldType = 0
	FieldType_TEXT             FieldType = 1
	FieldType_TEXTAREA         FieldType = 2
	FieldType_NUMBER           FieldType = 3
	FieldType_DATETIME         FieldType = 4
	FieldType_RADIO            FieldType = 5
	FieldType_CHECKBOX         FieldType = 6
	FieldType_SELECT           FieldType = 7
	FieldType_SELECTS          FieldType = 8
	FieldType_ADDRESS          FieldType = 9
	FieldType_LOCATION         FieldType = 10
	FieldType_IMAGE            FieldType = 11
	FieldType_FILE             FieldType = 12
	FieldType_SUBFORM          FieldType = 13
	FieldType_RELATION         FieldType = 14
	FieldType_RELATIONDATA     FieldType = 15
	FieldType_SIGNATURE        FieldType = 16
	FieldType_SERIALNUMBER     FieldType = 17
	FieldType_MOBILE           FieldType = 18
	FieldType_OCR              FieldType = 19
	FieldType_MEMBER           FieldType = 21
	FieldType_MEMBERS          FieldType = 22
	FieldType_DEPARTMENT       FieldType = 23
	FieldType_DEPARTMENTS      FieldType = 24
)

// Enum value maps for FieldType.
var (
	FieldType_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "TEXT",
		2:  "TEXTAREA",
		3:  "NUMBER",
		4:  "DATETIME",
		5:  "RADIO",
		6:  "CHECKBOX",
		7:  "SELECT",
		8:  "SELECTS",
		9:  "ADDRESS",
		10: "LOCATION",
		11: "IMAGE",
		12: "FILE",
		13: "SUBFORM",
		14: "RELATION",
		15: "RELATIONDATA",
		16: "SIGNATURE",
		17: "SERIALNUMBER",
		18: "MOBILE",
		19: "OCR",
		21: "MEMBER",
		22: "MEMBERS",
		23: "DEPARTMENT",
		24: "DEPARTMENTS",
	}
	FieldType_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TEXT":             1,
		"TEXTAREA":         2,
		"NUMBER":           3,
		"DATETIME":         4,
		"RADIO":            5,
		"CHECKBOX":         6,
		"SELECT":           7,
		"SELECTS":          8,
		"ADDRESS":          9,
		"LOCATION":         10,
		"IMAGE":            11,
		"FILE":             12,
		"SUBFORM":          13,
		"RELATION":         14,
		"RELATIONDATA":     15,
		"SIGNATURE":        16,
		"SERIALNUMBER":     17,
		"MOBILE":           18,
		"OCR":              19,
		"MEMBER":           21,
		"MEMBERS":          22,
		"DEPARTMENT":       23,
		"DEPARTMENTS":      24,
	}
)

func (x FieldType) Enum() *FieldType {
	p := new(FieldType)
	*p = x
	return p
}

func (x FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_models_proto_enumTypes[0].Descriptor()
}

func (FieldType) Type() protoreflect.EnumType {
	return &file_models_proto_enumTypes[0]
}

func (x FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldType.Descriptor instead.
func (FieldType) EnumDescriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

type GetModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Models []*Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
}

func (x *GetModelsResponse) Reset() {
	*x = GetModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelsResponse) ProtoMessage() {}

func (x *GetModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelsResponse.ProtoReflect.Descriptor instead.
func (*GetModelsResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *GetModelsResponse) GetModels() []*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UUID   string   `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Fields []*Field `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *Model) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                   FieldType   `protobuf:"varint,1,opt,name=type,proto3,enum=service.FieldType" json:"type,omitempty"`
	FieldId                int32       `protobuf:"varint,2,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"`
	Name                   *I18NString `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Required               bool        `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	Unique                 bool        `protobuf:"varint,5,opt,name=unique,proto3" json:"unique,omitempty"`
	Editable               bool        `protobuf:"varint,6,opt,name=editable,proto3" json:"editable,omitempty"`
	Visible                bool        `protobuf:"varint,7,opt,name=visible,proto3" json:"visible,omitempty"`
	ShowLabel              bool        `protobuf:"varint,8,opt,name=show_label,json=showLabel,proto3" json:"show_label,omitempty"`
	Description            *I18NString `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Key                    string      `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	NumberDecimal          uint32      `protobuf:"varint,11,opt,name=number_decimal,json=numberDecimal,proto3" json:"number_decimal,omitempty"`
	Options                []string    `protobuf:"bytes,12,rep,name=options,proto3" json:"options,omitempty"`
	SubformModel_UUID      string      `protobuf:"bytes,13,opt,name=subform_model_UUID,json=subformModelUUID,proto3" json:"subform_model_UUID,omitempty"`
	SubformRecordsLimitMin int32       `protobuf:"varint,14,opt,name=subform_records_limit_min,json=subformRecordsLimitMin,proto3" json:"subform_records_limit_min,omitempty"`
	SubformRecordsLimitMax int32       `protobuf:"varint,15,opt,name=subform_records_limit_max,json=subformRecordsLimitMax,proto3" json:"subform_records_limit_max,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{2}
}

func (x *Field) GetType() FieldType {
	if x != nil {
		return x.Type
	}
	return FieldType_TYPE_UNSPECIFIED
}

func (x *Field) GetFieldId() int32 {
	if x != nil {
		return x.FieldId
	}
	return 0
}

func (x *Field) GetName() *I18NString {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Field) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Field) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

func (x *Field) GetEditable() bool {
	if x != nil {
		return x.Editable
	}
	return false
}

func (x *Field) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Field) GetShowLabel() bool {
	if x != nil {
		return x.ShowLabel
	}
	return false
}

func (x *Field) GetDescription() *I18NString {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Field) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Field) GetNumberDecimal() uint32 {
	if x != nil {
		return x.NumberDecimal
	}
	return 0
}

func (x *Field) GetOptions() []string {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Field) GetSubformModel_UUID() string {
	if x != nil {
		return x.SubformModel_UUID
	}
	return ""
}

func (x *Field) GetSubformRecordsLimitMin() int32 {
	if x != nil {
		return x.SubformRecordsLimitMin
	}
	return 0
}

func (x *Field) GetSubformRecordsLimitMax() int32 {
	if x != nil {
		return x.SubformRecordsLimitMax
	}
	return 0
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x22, 0x57, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xaa, 0x04, 0x0a, 0x05,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x31, 0x38, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x68, 0x6f, 0x77, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x68, 0x6f, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x31, 0x38, 0x6e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x73, 0x75, 0x62, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x75, 0x62, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x73, 0x75, 0x62, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x39, 0x0a,
	0x19, 0x73, 0x75, 0x62, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x16, 0x73, 0x75, 0x62, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x2a, 0xd2, 0x02, 0x0a, 0x09, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45, 0x58, 0x54, 0x41, 0x52,
	0x45, 0x41, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x41, 0x44, 0x49, 0x4f, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x45,
	0x43, 0x4b, 0x42, 0x4f, 0x58, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x4c, 0x45, 0x43,
	0x54, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x53, 0x10, 0x08,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x09, 0x12, 0x0c, 0x0a,
	0x08, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x0c,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x42, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0x0d, 0x12, 0x0c, 0x0a,
	0x08, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x44, 0x41, 0x54, 0x41, 0x10, 0x0f, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x10, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x11, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x12, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x43,
	0x52, 0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x15, 0x12,
	0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x10, 0x16, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x17, 0x12, 0x0f, 0x0a, 0x0b,
	0x44, 0x45, 0x50, 0x41, 0x52, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x18, 0x42, 0x5f, 0x0a,
	0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x42, 0x12, 0x43, 0x61, 0x72, 0x74, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_models_proto_goTypes = []interface{}{
	(FieldType)(0),            // 0: service.FieldType
	(*GetModelsResponse)(nil), // 1: service.GetModelsResponse
	(*Model)(nil),             // 2: service.Model
	(*Field)(nil),             // 3: service.Field
	(*I18NString)(nil),        // 4: service.I18nString
}
var file_models_proto_depIdxs = []int32{
	2, // 0: service.GetModelsResponse.models:type_name -> service.Model
	3, // 1: service.Model.fields:type_name -> service.Field
	0, // 2: service.Field.type:type_name -> service.FieldType
	4, // 3: service.Field.name:type_name -> service.I18nString
	4, // 4: service.Field.description:type_name -> service.I18nString
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelsResponse); i {
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
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		EnumInfos:         file_models_proto_enumTypes,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}
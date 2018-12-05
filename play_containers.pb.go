// Code generated by protoc-gen-go. DO NOT EDIT.
// source: play_containers.proto

package libreplay

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContainerMetadata struct {
	BrowseUrl            *string          `protobuf:"bytes,1,opt,name=browseUrl" json:"browseUrl,omitempty"`
	NextPageUrl          *string          `protobuf:"bytes,2,opt,name=nextPageUrl" json:"nextPageUrl,omitempty"`
	Relevance            *float64         `protobuf:"fixed64,3,opt,name=relevance" json:"relevance,omitempty"`
	EstimatedResults     *int64           `protobuf:"varint,4,opt,name=estimatedResults" json:"estimatedResults,omitempty"`
	AnalyticsCookie      *string          `protobuf:"bytes,5,opt,name=analyticsCookie" json:"analyticsCookie,omitempty"`
	Ordered              *bool            `protobuf:"varint,6,opt,name=ordered" json:"ordered,omitempty"`
	ContainerView        []*ContainerView `protobuf:"bytes,7,rep,name=containerView" json:"containerView,omitempty"`
	LeftIcon             *Image           `protobuf:"bytes,8,opt,name=leftIcon" json:"leftIcon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ContainerMetadata) Reset()         { *m = ContainerMetadata{} }
func (m *ContainerMetadata) String() string { return proto.CompactTextString(m) }
func (*ContainerMetadata) ProtoMessage()    {}
func (*ContainerMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c226e85db89812, []int{0}
}

func (m *ContainerMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerMetadata.Unmarshal(m, b)
}
func (m *ContainerMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerMetadata.Marshal(b, m, deterministic)
}
func (m *ContainerMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerMetadata.Merge(m, src)
}
func (m *ContainerMetadata) XXX_Size() int {
	return xxx_messageInfo_ContainerMetadata.Size(m)
}
func (m *ContainerMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerMetadata proto.InternalMessageInfo

func (m *ContainerMetadata) GetBrowseUrl() string {
	if m != nil && m.BrowseUrl != nil {
		return *m.BrowseUrl
	}
	return ""
}

func (m *ContainerMetadata) GetNextPageUrl() string {
	if m != nil && m.NextPageUrl != nil {
		return *m.NextPageUrl
	}
	return ""
}

func (m *ContainerMetadata) GetRelevance() float64 {
	if m != nil && m.Relevance != nil {
		return *m.Relevance
	}
	return 0
}

func (m *ContainerMetadata) GetEstimatedResults() int64 {
	if m != nil && m.EstimatedResults != nil {
		return *m.EstimatedResults
	}
	return 0
}

func (m *ContainerMetadata) GetAnalyticsCookie() string {
	if m != nil && m.AnalyticsCookie != nil {
		return *m.AnalyticsCookie
	}
	return ""
}

func (m *ContainerMetadata) GetOrdered() bool {
	if m != nil && m.Ordered != nil {
		return *m.Ordered
	}
	return false
}

func (m *ContainerMetadata) GetContainerView() []*ContainerView {
	if m != nil {
		return m.ContainerView
	}
	return nil
}

func (m *ContainerMetadata) GetLeftIcon() *Image {
	if m != nil {
		return m.LeftIcon
	}
	return nil
}

type ContainerView struct {
	Selected             *bool    `protobuf:"varint,1,opt,name=selected" json:"selected,omitempty"`
	Title                *string  `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	ListUrl              *string  `protobuf:"bytes,3,opt,name=listUrl" json:"listUrl,omitempty"`
	ServerLogsCookie     []byte   `protobuf:"bytes,4,opt,name=serverLogsCookie" json:"serverLogsCookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerView) Reset()         { *m = ContainerView{} }
func (m *ContainerView) String() string { return proto.CompactTextString(m) }
func (*ContainerView) ProtoMessage()    {}
func (*ContainerView) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c226e85db89812, []int{1}
}

func (m *ContainerView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerView.Unmarshal(m, b)
}
func (m *ContainerView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerView.Marshal(b, m, deterministic)
}
func (m *ContainerView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerView.Merge(m, src)
}
func (m *ContainerView) XXX_Size() int {
	return xxx_messageInfo_ContainerView.Size(m)
}
func (m *ContainerView) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerView.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerView proto.InternalMessageInfo

func (m *ContainerView) GetSelected() bool {
	if m != nil && m.Selected != nil {
		return *m.Selected
	}
	return false
}

func (m *ContainerView) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *ContainerView) GetListUrl() string {
	if m != nil && m.ListUrl != nil {
		return *m.ListUrl
	}
	return ""
}

func (m *ContainerView) GetServerLogsCookie() []byte {
	if m != nil {
		return m.ServerLogsCookie
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerMetadata)(nil), "playapi.proto.finsky.containers.ContainerMetadata")
	proto.RegisterType((*ContainerView)(nil), "playapi.proto.finsky.containers.ContainerView")
}

func init() { proto.RegisterFile("play_containers.proto", fileDescriptor_32c226e85db89812) }

var fileDescriptor_32c226e85db89812 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4e, 0x2a, 0x31,
	0x14, 0xc6, 0x53, 0xe6, 0x72, 0x19, 0xca, 0x25, 0x57, 0x1a, 0x4d, 0x1a, 0x34, 0x61, 0x64, 0x35,
	0x71, 0x31, 0x24, 0x6c, 0xdc, 0x8b, 0x1b, 0x8c, 0x26, 0xa6, 0x51, 0x17, 0x6e, 0x4c, 0x99, 0x39,
	0x0c, 0x0d, 0x9d, 0x96, 0xb4, 0x05, 0xe4, 0x09, 0x7c, 0x5d, 0x1f, 0xc1, 0xcc, 0xc0, 0xf0, 0x47,
	0x4c, 0x5c, 0x7e, 0xbf, 0x7e, 0xe7, 0x34, 0xe7, 0x87, 0xcf, 0x66, 0x92, 0xaf, 0xde, 0x62, 0xad,
	0x1c, 0x17, 0x0a, 0x8c, 0x8d, 0x66, 0x46, 0x3b, 0x4d, 0x3a, 0x39, 0xe6, 0x33, 0xb1, 0x8e, 0xd1,
	0x58, 0x28, 0x3b, 0x5d, 0x45, 0xbb, 0x5a, 0xbb, 0xb5, 0x99, 0xcb, 0x32, 0xad, 0xd6, 0xa5, 0xee,
	0x67, 0x05, 0xb7, 0x06, 0x65, 0xe3, 0x01, 0x1c, 0x4f, 0xb8, 0xe3, 0xe4, 0x02, 0xd7, 0x47, 0x46,
	0x2f, 0x2d, 0x3c, 0x1b, 0x49, 0x51, 0x80, 0xc2, 0x3a, 0xdb, 0x01, 0x12, 0xe0, 0x86, 0x82, 0x77,
	0xf7, 0xc8, 0xd3, 0xe2, 0xbd, 0x52, 0xbc, 0xef, 0xa3, 0x7c, 0xde, 0x80, 0x84, 0x05, 0x57, 0x31,
	0x50, 0x2f, 0x40, 0x21, 0x62, 0x3b, 0x40, 0xae, 0xf0, 0x09, 0x58, 0x27, 0x32, 0xee, 0x20, 0x61,
	0x60, 0xe7, 0xd2, 0x59, 0xfa, 0x27, 0x40, 0xa1, 0xc7, 0x8e, 0x38, 0x09, 0xf1, 0x7f, 0xae, 0xb8,
	0x5c, 0x39, 0x11, 0xdb, 0x81, 0xd6, 0x53, 0x01, 0xb4, 0x5a, 0xfc, 0xf7, 0x1d, 0x13, 0x8a, 0x6b,
	0xda, 0x24, 0x60, 0x20, 0xa1, 0x7f, 0x03, 0x14, 0xfa, 0xac, 0x8c, 0xe4, 0x09, 0x37, 0xb7, 0x12,
	0x5e, 0x04, 0x2c, 0x69, 0x2d, 0xf0, 0xc2, 0x46, 0x3f, 0x8a, 0x7e, 0xf1, 0x15, 0x0d, 0xf6, 0xa7,
	0xd8, 0xe1, 0x12, 0x72, 0x8d, 0x7d, 0x09, 0x63, 0x37, 0x8c, 0xb5, 0xa2, 0x7e, 0x80, 0xc2, 0x46,
	0xff, 0xfc, 0xe7, 0x85, 0xc3, 0x8c, 0xa7, 0xc0, 0xb6, 0xe5, 0xee, 0x07, 0xc2, 0xcd, 0x83, 0xcd,
	0xa4, 0x8d, 0x7d, 0x0b, 0x12, 0x62, 0x07, 0x49, 0x61, 0xdb, 0x67, 0xdb, 0x4c, 0x4e, 0x71, 0xd5,
	0x09, 0x27, 0x61, 0xa3, 0x79, 0x1d, 0xf2, 0x63, 0xa5, 0xb0, 0x2e, 0xd7, 0xef, 0x15, 0xbc, 0x8c,
	0xb9, 0x5c, 0x0b, 0x66, 0x01, 0xe6, 0x5e, 0xa7, 0xa5, 0xb1, 0x5c, 0xee, 0x3f, 0x76, 0xc4, 0x6f,
	0x2e, 0x5f, 0x3b, 0xa9, 0x70, 0x93, 0xf9, 0x28, 0x8a, 0x75, 0xd6, 0xbb, 0xd3, 0x76, 0x32, 0xe7,
	0xb7, 0x1a, 0x6c, 0x4f, 0x8a, 0x91, 0x81, 0xfc, 0x98, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44,
	0x44, 0x16, 0x86, 0x6b, 0x02, 0x00, 0x00,
}

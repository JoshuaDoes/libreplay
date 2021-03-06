// Code generated by protoc-gen-go. DO NOT EDIT.
// source: play_search.proto

package libplay

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

type SearchResponse struct {
	OriginalQuery  *string `protobuf:"bytes,1,opt,name=originalQuery" json:"originalQuery,omitempty"`
	SuggestedQuery *string `protobuf:"bytes,2,opt,name=suggestedQuery" json:"suggestedQuery,omitempty"`
	AggregateQuery *bool   `protobuf:"varint,3,opt,name=aggregateQuery" json:"aggregateQuery,omitempty"`
	// unused: repeated Bucket bucket = 4;
	Doc                  []*DocV2         `protobuf:"bytes,5,rep,name=doc" json:"doc,omitempty"`
	RelatedSearch        []*RelatedSearch `protobuf:"bytes,6,rep,name=relatedSearch" json:"relatedSearch,omitempty"`
	ServerLogsCookie     []byte           `protobuf:"bytes,7,opt,name=serverLogsCookie" json:"serverLogsCookie,omitempty"`
	FullPageReplaced     *bool            `protobuf:"varint,8,opt,name=fullPageReplaced" json:"fullPageReplaced,omitempty"`
	ContainsSnow         *bool            `protobuf:"varint,9,opt,name=containsSnow" json:"containsSnow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4098c258d91b55cd, []int{0}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetOriginalQuery() string {
	if m != nil && m.OriginalQuery != nil {
		return *m.OriginalQuery
	}
	return ""
}

func (m *SearchResponse) GetSuggestedQuery() string {
	if m != nil && m.SuggestedQuery != nil {
		return *m.SuggestedQuery
	}
	return ""
}

func (m *SearchResponse) GetAggregateQuery() bool {
	if m != nil && m.AggregateQuery != nil {
		return *m.AggregateQuery
	}
	return false
}

func (m *SearchResponse) GetDoc() []*DocV2 {
	if m != nil {
		return m.Doc
	}
	return nil
}

func (m *SearchResponse) GetRelatedSearch() []*RelatedSearch {
	if m != nil {
		return m.RelatedSearch
	}
	return nil
}

func (m *SearchResponse) GetServerLogsCookie() []byte {
	if m != nil {
		return m.ServerLogsCookie
	}
	return nil
}

func (m *SearchResponse) GetFullPageReplaced() bool {
	if m != nil && m.FullPageReplaced != nil {
		return *m.FullPageReplaced
	}
	return false
}

func (m *SearchResponse) GetContainsSnow() bool {
	if m != nil && m.ContainsSnow != nil {
		return *m.ContainsSnow
	}
	return false
}

type RelatedSearch struct {
	SearchUrl            *string  `protobuf:"bytes,1,opt,name=searchUrl" json:"searchUrl,omitempty"`
	Header               *string  `protobuf:"bytes,2,opt,name=header" json:"header,omitempty"`
	BackendId            *int32   `protobuf:"varint,3,opt,name=backendId" json:"backendId,omitempty"`
	DocType              *int32   `protobuf:"varint,4,opt,name=docType" json:"docType,omitempty"`
	Current              *bool    `protobuf:"varint,5,opt,name=current" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelatedSearch) Reset()         { *m = RelatedSearch{} }
func (m *RelatedSearch) String() string { return proto.CompactTextString(m) }
func (*RelatedSearch) ProtoMessage()    {}
func (*RelatedSearch) Descriptor() ([]byte, []int) {
	return fileDescriptor_4098c258d91b55cd, []int{1}
}

func (m *RelatedSearch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelatedSearch.Unmarshal(m, b)
}
func (m *RelatedSearch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelatedSearch.Marshal(b, m, deterministic)
}
func (m *RelatedSearch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelatedSearch.Merge(m, src)
}
func (m *RelatedSearch) XXX_Size() int {
	return xxx_messageInfo_RelatedSearch.Size(m)
}
func (m *RelatedSearch) XXX_DiscardUnknown() {
	xxx_messageInfo_RelatedSearch.DiscardUnknown(m)
}

var xxx_messageInfo_RelatedSearch proto.InternalMessageInfo

func (m *RelatedSearch) GetSearchUrl() string {
	if m != nil && m.SearchUrl != nil {
		return *m.SearchUrl
	}
	return ""
}

func (m *RelatedSearch) GetHeader() string {
	if m != nil && m.Header != nil {
		return *m.Header
	}
	return ""
}

func (m *RelatedSearch) GetBackendId() int32 {
	if m != nil && m.BackendId != nil {
		return *m.BackendId
	}
	return 0
}

func (m *RelatedSearch) GetDocType() int32 {
	if m != nil && m.DocType != nil {
		return *m.DocType
	}
	return 0
}

func (m *RelatedSearch) GetCurrent() bool {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return false
}

type SearchSuggestResponse struct {
	Suggestion           []*Suggestion `protobuf:"bytes,1,rep,name=suggestion" json:"suggestion,omitempty"`
	ServerLogsCookie     []byte        `protobuf:"bytes,2,opt,name=serverLogsCookie" json:"serverLogsCookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SearchSuggestResponse) Reset()         { *m = SearchSuggestResponse{} }
func (m *SearchSuggestResponse) String() string { return proto.CompactTextString(m) }
func (*SearchSuggestResponse) ProtoMessage()    {}
func (*SearchSuggestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4098c258d91b55cd, []int{2}
}

func (m *SearchSuggestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchSuggestResponse.Unmarshal(m, b)
}
func (m *SearchSuggestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchSuggestResponse.Marshal(b, m, deterministic)
}
func (m *SearchSuggestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchSuggestResponse.Merge(m, src)
}
func (m *SearchSuggestResponse) XXX_Size() int {
	return xxx_messageInfo_SearchSuggestResponse.Size(m)
}
func (m *SearchSuggestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchSuggestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchSuggestResponse proto.InternalMessageInfo

func (m *SearchSuggestResponse) GetSuggestion() []*Suggestion {
	if m != nil {
		return m.Suggestion
	}
	return nil
}

func (m *SearchSuggestResponse) GetServerLogsCookie() []byte {
	if m != nil {
		return m.ServerLogsCookie
	}
	return nil
}

type Suggestion struct {
	Type                 *int32         `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	SuggestedQuery       *string        `protobuf:"bytes,2,opt,name=suggestedQuery" json:"suggestedQuery,omitempty"`
	NavSuggestion        *NavSuggestion `protobuf:"bytes,3,opt,name=navSuggestion" json:"navSuggestion,omitempty"`
	ServerLogsCookie     []byte         `protobuf:"bytes,4,opt,name=serverLogsCookie" json:"serverLogsCookie,omitempty"`
	Image                *Image         `protobuf:"bytes,5,opt,name=image" json:"image,omitempty"`
	DisplayText          *string        `protobuf:"bytes,6,opt,name=displayText" json:"displayText,omitempty"`
	Link                 *Link          `protobuf:"bytes,7,opt,name=link" json:"link,omitempty"`
	Document             *DocV2         `protobuf:"bytes,8,opt,name=document" json:"document,omitempty"`
	SearchBackend        *int32         `protobuf:"varint,9,opt,name=searchBackend" json:"searchBackend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Suggestion) Reset()         { *m = Suggestion{} }
func (m *Suggestion) String() string { return proto.CompactTextString(m) }
func (*Suggestion) ProtoMessage()    {}
func (*Suggestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_4098c258d91b55cd, []int{3}
}

func (m *Suggestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suggestion.Unmarshal(m, b)
}
func (m *Suggestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suggestion.Marshal(b, m, deterministic)
}
func (m *Suggestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suggestion.Merge(m, src)
}
func (m *Suggestion) XXX_Size() int {
	return xxx_messageInfo_Suggestion.Size(m)
}
func (m *Suggestion) XXX_DiscardUnknown() {
	xxx_messageInfo_Suggestion.DiscardUnknown(m)
}

var xxx_messageInfo_Suggestion proto.InternalMessageInfo

func (m *Suggestion) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Suggestion) GetSuggestedQuery() string {
	if m != nil && m.SuggestedQuery != nil {
		return *m.SuggestedQuery
	}
	return ""
}

func (m *Suggestion) GetNavSuggestion() *NavSuggestion {
	if m != nil {
		return m.NavSuggestion
	}
	return nil
}

func (m *Suggestion) GetServerLogsCookie() []byte {
	if m != nil {
		return m.ServerLogsCookie
	}
	return nil
}

func (m *Suggestion) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Suggestion) GetDisplayText() string {
	if m != nil && m.DisplayText != nil {
		return *m.DisplayText
	}
	return ""
}

func (m *Suggestion) GetLink() *Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Suggestion) GetDocument() *DocV2 {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *Suggestion) GetSearchBackend() int32 {
	if m != nil && m.SearchBackend != nil {
		return *m.SearchBackend
	}
	return 0
}

type NavSuggestion struct {
	DocId                *string  `protobuf:"bytes,1,opt,name=docId" json:"docId,omitempty"`
	ImageBlob            []byte   `protobuf:"bytes,2,opt,name=imageBlob" json:"imageBlob,omitempty"`
	Image                *Image   `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Description          *string  `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NavSuggestion) Reset()         { *m = NavSuggestion{} }
func (m *NavSuggestion) String() string { return proto.CompactTextString(m) }
func (*NavSuggestion) ProtoMessage()    {}
func (*NavSuggestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_4098c258d91b55cd, []int{4}
}

func (m *NavSuggestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NavSuggestion.Unmarshal(m, b)
}
func (m *NavSuggestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NavSuggestion.Marshal(b, m, deterministic)
}
func (m *NavSuggestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NavSuggestion.Merge(m, src)
}
func (m *NavSuggestion) XXX_Size() int {
	return xxx_messageInfo_NavSuggestion.Size(m)
}
func (m *NavSuggestion) XXX_DiscardUnknown() {
	xxx_messageInfo_NavSuggestion.DiscardUnknown(m)
}

var xxx_messageInfo_NavSuggestion proto.InternalMessageInfo

func (m *NavSuggestion) GetDocId() string {
	if m != nil && m.DocId != nil {
		return *m.DocId
	}
	return ""
}

func (m *NavSuggestion) GetImageBlob() []byte {
	if m != nil {
		return m.ImageBlob
	}
	return nil
}

func (m *NavSuggestion) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *NavSuggestion) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchResponse)(nil), "playapi.proto.finsky.search.SearchResponse")
	proto.RegisterType((*RelatedSearch)(nil), "playapi.proto.finsky.search.RelatedSearch")
	proto.RegisterType((*SearchSuggestResponse)(nil), "playapi.proto.finsky.search.SearchSuggestResponse")
	proto.RegisterType((*Suggestion)(nil), "playapi.proto.finsky.search.Suggestion")
	proto.RegisterType((*NavSuggestion)(nil), "playapi.proto.finsky.search.NavSuggestion")
}

func init() { proto.RegisterFile("play_search.proto", fileDescriptor_4098c258d91b55cd) }

var fileDescriptor_4098c258d91b55cd = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0xd6, 0x76, 0x7f, 0x6e, 0xd7, 0xfd, 0x7e, 0x98, 0x3f, 0x8a, 0x36, 0xa4, 0x95, 0x6a,
	0x82, 0x6a, 0x0f, 0x99, 0x28, 0x12, 0xcf, 0x68, 0x4c, 0x42, 0x43, 0x13, 0x1a, 0xee, 0xe0, 0x81,
	0x17, 0xe4, 0x3a, 0x77, 0xa9, 0xd5, 0xd4, 0x8e, 0xec, 0x64, 0xd0, 0xef, 0xb0, 0x77, 0xf8, 0x28,
	0x7c, 0x3c, 0x64, 0x3b, 0x6d, 0x9a, 0xad, 0x8c, 0xf1, 0xd6, 0x7b, 0x72, 0xee, 0xad, 0xcf, 0xf1,
	0xb9, 0x86, 0x07, 0x59, 0xca, 0x66, 0x5f, 0x0d, 0x32, 0xcd, 0xc7, 0x51, 0xa6, 0x55, 0xae, 0xc8,
	0x9e, 0x85, 0x58, 0x26, 0x7c, 0x19, 0x5d, 0x0a, 0x69, 0x26, 0xb3, 0xc8, 0x53, 0x76, 0x3d, 0x9f,
	0xab, 0xe9, 0x54, 0x49, 0x4f, 0xd8, 0x7d, 0xe8, 0xa0, 0x58, 0xf1, 0x62, 0x8a, 0x32, 0x2f, 0xc1,
	0xff, 0x1c, 0x98, 0x0a, 0x39, 0xf1, 0x40, 0xef, 0xba, 0x01, 0x3b, 0x43, 0x37, 0x83, 0xa2, 0xc9,
	0x94, 0x34, 0x48, 0x0e, 0xa0, 0xa3, 0xb4, 0x48, 0x84, 0x64, 0xe9, 0xc7, 0x02, 0xf5, 0x2c, 0x0c,
	0xba, 0x41, 0x7f, 0x8b, 0xd6, 0x41, 0xf2, 0x1c, 0x76, 0x4c, 0x91, 0x24, 0x68, 0x72, 0x8c, 0x3d,
	0x6d, 0xcd, 0xd1, 0x6e, 0xa0, 0x96, 0xc7, 0x92, 0x44, 0x63, 0xc2, 0x72, 0xf4, 0xbc, 0x46, 0x37,
	0xe8, 0x6f, 0xd2, 0x1b, 0x28, 0x79, 0x0d, 0x8d, 0x58, 0xf1, 0xb0, 0xd5, 0x6d, 0xf4, 0xdb, 0x83,
	0x83, 0x68, 0xa5, 0xd8, 0x85, 0x98, 0x13, 0xc5, 0x3f, 0x0f, 0xa8, 0x6d, 0x20, 0xe7, 0xd0, 0xd1,
	0x98, 0xb2, 0x1c, 0x63, 0x2f, 0x23, 0x5c, 0x77, 0x13, 0x0e, 0xa3, 0x3b, 0xec, 0x8a, 0xe8, 0x72,
	0x07, 0xad, 0x0f, 0x20, 0x87, 0xf0, 0xbf, 0x41, 0x7d, 0x85, 0xfa, 0x4c, 0x25, 0xe6, 0xad, 0x52,
	0x13, 0x81, 0xe1, 0x46, 0x37, 0xe8, 0x6f, 0xd3, 0x5b, 0xb8, 0xe5, 0x5e, 0x16, 0x69, 0x7a, 0xce,
	0x12, 0xa4, 0x98, 0xa5, 0x8c, 0x63, 0x1c, 0x6e, 0x3a, 0x7d, 0xb7, 0x70, 0xd2, 0x83, 0x6d, 0xae,
	0x64, 0xce, 0x84, 0x34, 0x43, 0xa9, 0xbe, 0x85, 0x5b, 0x8e, 0x57, 0xc3, 0x7a, 0x3f, 0x02, 0xe8,
	0xd4, 0x0e, 0x47, 0x9e, 0xc2, 0x96, 0x3f, 0xf4, 0x27, 0x9d, 0x96, 0x37, 0x51, 0x01, 0xe4, 0x09,
	0xac, 0x8f, 0x91, 0xc5, 0xa8, 0x4b, 0xf7, 0xcb, 0xca, 0x76, 0x8d, 0x18, 0x9f, 0xa0, 0x8c, 0x4f,
	0x63, 0x67, 0x78, 0x8b, 0x56, 0x00, 0x09, 0x61, 0x23, 0x56, 0xfc, 0x62, 0x96, 0x61, 0xd8, 0x74,
	0xdf, 0xe6, 0xa5, 0xfd, 0xc2, 0x0b, 0xad, 0x51, 0xe6, 0x61, 0xcb, 0x1d, 0x6f, 0x5e, 0xf6, 0xae,
	0x03, 0x78, 0xec, 0x8f, 0x34, 0xf4, 0x17, 0xbc, 0xc8, 0xcb, 0x3b, 0x80, 0xf2, 0xce, 0x85, 0x92,
	0x61, 0xe0, 0xec, 0x7f, 0x71, 0xa7, 0xfd, 0xc3, 0x05, 0x9d, 0x2e, 0xb5, 0xae, 0x34, 0x7e, 0x6d,
	0xb5, 0xf1, 0xbd, 0x5f, 0x0d, 0x80, 0x6a, 0x0c, 0x21, 0xd0, 0xcc, 0xad, 0x9c, 0xc0, 0xc9, 0x71,
	0xbf, 0xef, 0x9d, 0xd0, 0x73, 0xe8, 0x48, 0x76, 0x55, 0x0d, 0x73, 0x7e, 0xfd, 0x2d, 0x41, 0x1f,
	0x96, 0x3b, 0x68, 0x7d, 0xc0, 0x4a, 0x21, 0xcd, 0x3f, 0x24, 0xe8, 0x25, 0xb4, 0xc4, 0x94, 0x25,
	0xe8, 0xfc, 0x6e, 0x0f, 0xf6, 0x56, 0xff, 0xeb, 0xa9, 0xa5, 0x50, 0xcf, 0x24, 0x5d, 0x68, 0xc7,
	0xc2, 0x58, 0xde, 0x05, 0x7e, 0xcf, 0xc3, 0x75, 0xa7, 0x6a, 0x19, 0x22, 0xaf, 0xa0, 0x69, 0x77,
	0xdc, 0xc5, 0xb6, 0x3d, 0xd8, 0x5f, 0x3d, 0xd3, 0xbd, 0x02, 0x67, 0x42, 0x4e, 0xa8, 0x23, 0x93,
	0x37, 0xb0, 0x39, 0x5f, 0x30, 0x97, 0xe1, 0xfb, 0xae, 0xe1, 0xa2, 0xcb, 0xbe, 0x1c, 0xde, 0x9e,
	0x63, 0x1f, 0x35, 0x17, 0xf1, 0x16, 0xad, 0x83, 0xbd, 0x9f, 0x01, 0x74, 0x6a, 0xf6, 0x91, 0x47,
	0xd0, 0x8a, 0x15, 0x3f, 0x8d, 0xcb, 0x7c, 0xfb, 0xc2, 0x66, 0xd8, 0xe9, 0x3d, 0x4e, 0xd5, 0xa8,
	0xcc, 0x41, 0x05, 0x54, 0xbe, 0x35, 0xfe, 0xc9, 0x37, 0x34, 0x5c, 0x8b, 0xcc, 0x5d, 0x73, 0xb3,
	0xf4, 0xad, 0x82, 0x8e, 0x9f, 0x7d, 0xd9, 0x4f, 0x44, 0x3e, 0x2e, 0x46, 0x11, 0x57, 0xd3, 0xa3,
	0xf7, 0xca, 0x8c, 0x0b, 0x76, 0xa2, 0xd0, 0x1c, 0xa5, 0x62, 0xa4, 0xed, 0x1e, 0xcf, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x0c, 0x40, 0x0a, 0x9a, 0x05, 0x00, 0x00,
}

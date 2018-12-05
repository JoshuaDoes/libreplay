// Code generated by protoc-gen-go. DO NOT EDIT.
// source: play_details.proto

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

type BulkDetailsRequest struct {
	Docid                []string `protobuf:"bytes,1,rep,name=docid" json:"docid,omitempty"`
	IncludeChildDocs     *bool    `protobuf:"varint,2,opt,name=includeChildDocs" json:"includeChildDocs,omitempty"`
	IncludeDetails       *bool    `protobuf:"varint,3,opt,name=includeDetails" json:"includeDetails,omitempty"`
	SourcePackageName    *string  `protobuf:"bytes,4,opt,name=sourcePackageName" json:"sourcePackageName,omitempty"`
	InstalledVersionCode []int32  `protobuf:"varint,7,rep,name=installedVersionCode" json:"installedVersionCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkDetailsRequest) Reset()         { *m = BulkDetailsRequest{} }
func (m *BulkDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*BulkDetailsRequest) ProtoMessage()    {}
func (*BulkDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{0}
}

func (m *BulkDetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDetailsRequest.Unmarshal(m, b)
}
func (m *BulkDetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDetailsRequest.Marshal(b, m, deterministic)
}
func (m *BulkDetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDetailsRequest.Merge(m, src)
}
func (m *BulkDetailsRequest) XXX_Size() int {
	return xxx_messageInfo_BulkDetailsRequest.Size(m)
}
func (m *BulkDetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDetailsRequest proto.InternalMessageInfo

func (m *BulkDetailsRequest) GetDocid() []string {
	if m != nil {
		return m.Docid
	}
	return nil
}

func (m *BulkDetailsRequest) GetIncludeChildDocs() bool {
	if m != nil && m.IncludeChildDocs != nil {
		return *m.IncludeChildDocs
	}
	return false
}

func (m *BulkDetailsRequest) GetIncludeDetails() bool {
	if m != nil && m.IncludeDetails != nil {
		return *m.IncludeDetails
	}
	return false
}

func (m *BulkDetailsRequest) GetSourcePackageName() string {
	if m != nil && m.SourcePackageName != nil {
		return *m.SourcePackageName
	}
	return ""
}

func (m *BulkDetailsRequest) GetInstalledVersionCode() []int32 {
	if m != nil {
		return m.InstalledVersionCode
	}
	return nil
}

type BulkDetailsResponse struct {
	Entry                []*BulkDetailsEntry `protobuf:"bytes,1,rep,name=entry" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BulkDetailsResponse) Reset()         { *m = BulkDetailsResponse{} }
func (m *BulkDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*BulkDetailsResponse) ProtoMessage()    {}
func (*BulkDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{1}
}

func (m *BulkDetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDetailsResponse.Unmarshal(m, b)
}
func (m *BulkDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDetailsResponse.Marshal(b, m, deterministic)
}
func (m *BulkDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDetailsResponse.Merge(m, src)
}
func (m *BulkDetailsResponse) XXX_Size() int {
	return xxx_messageInfo_BulkDetailsResponse.Size(m)
}
func (m *BulkDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDetailsResponse proto.InternalMessageInfo

func (m *BulkDetailsResponse) GetEntry() []*BulkDetailsEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type BulkDetailsEntry struct {
	Doc                  *DocV2   `protobuf:"bytes,1,opt,name=doc" json:"doc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BulkDetailsEntry) Reset()         { *m = BulkDetailsEntry{} }
func (m *BulkDetailsEntry) String() string { return proto.CompactTextString(m) }
func (*BulkDetailsEntry) ProtoMessage()    {}
func (*BulkDetailsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{2}
}

func (m *BulkDetailsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BulkDetailsEntry.Unmarshal(m, b)
}
func (m *BulkDetailsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BulkDetailsEntry.Marshal(b, m, deterministic)
}
func (m *BulkDetailsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BulkDetailsEntry.Merge(m, src)
}
func (m *BulkDetailsEntry) XXX_Size() int {
	return xxx_messageInfo_BulkDetailsEntry.Size(m)
}
func (m *BulkDetailsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BulkDetailsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BulkDetailsEntry proto.InternalMessageInfo

func (m *BulkDetailsEntry) GetDoc() *DocV2 {
	if m != nil {
		return m.Doc
	}
	return nil
}

type DetailsResponse struct {
	// deprecated optional DocV1 docV1 = 1;
	UserReview           *Review           `protobuf:"bytes,3,opt,name=userReview" json:"userReview,omitempty"`
	DocV2                *DocV2            `protobuf:"bytes,4,opt,name=docV2" json:"docV2,omitempty"`
	FooterHtml           *string           `protobuf:"bytes,5,opt,name=footerHtml" json:"footerHtml,omitempty"`
	ServerLogsCookie     []byte            `protobuf:"bytes,6,opt,name=serverLogsCookie" json:"serverLogsCookie,omitempty"`
	DiscoveryBadge       []*DiscoveryBadge `protobuf:"bytes,7,rep,name=discoveryBadge" json:"discoveryBadge,omitempty"`
	EnableReviews        *bool             `protobuf:"varint,8,opt,name=enableReviews" json:"enableReviews,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DetailsResponse) Reset()         { *m = DetailsResponse{} }
func (m *DetailsResponse) String() string { return proto.CompactTextString(m) }
func (*DetailsResponse) ProtoMessage()    {}
func (*DetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{3}
}

func (m *DetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailsResponse.Unmarshal(m, b)
}
func (m *DetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailsResponse.Marshal(b, m, deterministic)
}
func (m *DetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailsResponse.Merge(m, src)
}
func (m *DetailsResponse) XXX_Size() int {
	return xxx_messageInfo_DetailsResponse.Size(m)
}
func (m *DetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailsResponse proto.InternalMessageInfo

func (m *DetailsResponse) GetUserReview() *Review {
	if m != nil {
		return m.UserReview
	}
	return nil
}

func (m *DetailsResponse) GetDocV2() *DocV2 {
	if m != nil {
		return m.DocV2
	}
	return nil
}

func (m *DetailsResponse) GetFooterHtml() string {
	if m != nil && m.FooterHtml != nil {
		return *m.FooterHtml
	}
	return ""
}

func (m *DetailsResponse) GetServerLogsCookie() []byte {
	if m != nil {
		return m.ServerLogsCookie
	}
	return nil
}

func (m *DetailsResponse) GetDiscoveryBadge() []*DiscoveryBadge {
	if m != nil {
		return m.DiscoveryBadge
	}
	return nil
}

func (m *DetailsResponse) GetEnableReviews() bool {
	if m != nil && m.EnableReviews != nil {
		return *m.EnableReviews
	}
	return false
}

type DiscoveryBadge struct {
	Title                *string              `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Image                *Image               `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	BackgroundColor      *int32               `protobuf:"varint,3,opt,name=backgroundColor" json:"backgroundColor,omitempty"`
	DiscoveryBadgeLink   *DiscoveryBadgeLink  `protobuf:"bytes,4,opt,name=discoveryBadgeLink" json:"discoveryBadgeLink,omitempty"`
	ServerLogsCookie     []byte               `protobuf:"bytes,5,opt,name=serverLogsCookie" json:"serverLogsCookie,omitempty"`
	IsPlusOne            *bool                `protobuf:"varint,6,opt,name=isPlusOne" json:"isPlusOne,omitempty"`
	AggregateRating      *float32             `protobuf:"fixed32,7,opt,name=aggregateRating" json:"aggregateRating,omitempty"`
	UserStarRating       *int32               `protobuf:"varint,8,opt,name=userStarRating" json:"userStarRating,omitempty"`
	DownloadCount        *string              `protobuf:"bytes,9,opt,name=downloadCount" json:"downloadCount,omitempty"`
	DownloadUnits        *string              `protobuf:"bytes,10,opt,name=downloadUnits" json:"downloadUnits,omitempty"`
	ContentDescription   *string              `protobuf:"bytes,11,opt,name=contentDescription" json:"contentDescription,omitempty"`
	PlayerBadge          *PlayerBadge         `protobuf:"bytes,12,opt,name=playerBadge" json:"playerBadge,omitempty"`
	FamilyAgeRangeBadge  *FamilyAgeRangeBadge `protobuf:"bytes,13,opt,name=familyAgeRangeBadge" json:"familyAgeRangeBadge,omitempty"`
	FamilyCategoryBadge  *FamilyCategoryBadge `protobuf:"bytes,14,opt,name=familyCategoryBadge" json:"familyCategoryBadge,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DiscoveryBadge) Reset()         { *m = DiscoveryBadge{} }
func (m *DiscoveryBadge) String() string { return proto.CompactTextString(m) }
func (*DiscoveryBadge) ProtoMessage()    {}
func (*DiscoveryBadge) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{4}
}

func (m *DiscoveryBadge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryBadge.Unmarshal(m, b)
}
func (m *DiscoveryBadge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryBadge.Marshal(b, m, deterministic)
}
func (m *DiscoveryBadge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryBadge.Merge(m, src)
}
func (m *DiscoveryBadge) XXX_Size() int {
	return xxx_messageInfo_DiscoveryBadge.Size(m)
}
func (m *DiscoveryBadge) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryBadge.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryBadge proto.InternalMessageInfo

func (m *DiscoveryBadge) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *DiscoveryBadge) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *DiscoveryBadge) GetBackgroundColor() int32 {
	if m != nil && m.BackgroundColor != nil {
		return *m.BackgroundColor
	}
	return 0
}

func (m *DiscoveryBadge) GetDiscoveryBadgeLink() *DiscoveryBadgeLink {
	if m != nil {
		return m.DiscoveryBadgeLink
	}
	return nil
}

func (m *DiscoveryBadge) GetServerLogsCookie() []byte {
	if m != nil {
		return m.ServerLogsCookie
	}
	return nil
}

func (m *DiscoveryBadge) GetIsPlusOne() bool {
	if m != nil && m.IsPlusOne != nil {
		return *m.IsPlusOne
	}
	return false
}

func (m *DiscoveryBadge) GetAggregateRating() float32 {
	if m != nil && m.AggregateRating != nil {
		return *m.AggregateRating
	}
	return 0
}

func (m *DiscoveryBadge) GetUserStarRating() int32 {
	if m != nil && m.UserStarRating != nil {
		return *m.UserStarRating
	}
	return 0
}

func (m *DiscoveryBadge) GetDownloadCount() string {
	if m != nil && m.DownloadCount != nil {
		return *m.DownloadCount
	}
	return ""
}

func (m *DiscoveryBadge) GetDownloadUnits() string {
	if m != nil && m.DownloadUnits != nil {
		return *m.DownloadUnits
	}
	return ""
}

func (m *DiscoveryBadge) GetContentDescription() string {
	if m != nil && m.ContentDescription != nil {
		return *m.ContentDescription
	}
	return ""
}

func (m *DiscoveryBadge) GetPlayerBadge() *PlayerBadge {
	if m != nil {
		return m.PlayerBadge
	}
	return nil
}

func (m *DiscoveryBadge) GetFamilyAgeRangeBadge() *FamilyAgeRangeBadge {
	if m != nil {
		return m.FamilyAgeRangeBadge
	}
	return nil
}

func (m *DiscoveryBadge) GetFamilyCategoryBadge() *FamilyCategoryBadge {
	if m != nil {
		return m.FamilyCategoryBadge
	}
	return nil
}

type DiscoveryBadgeLink struct {
	Link                 *Link    `protobuf:"bytes,1,opt,name=link" json:"link,omitempty"`
	UserReviewsUrl       *string  `protobuf:"bytes,2,opt,name=userReviewsUrl" json:"userReviewsUrl,omitempty"`
	CriticReviewsUrl     *string  `protobuf:"bytes,3,opt,name=criticReviewsUrl" json:"criticReviewsUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryBadgeLink) Reset()         { *m = DiscoveryBadgeLink{} }
func (m *DiscoveryBadgeLink) String() string { return proto.CompactTextString(m) }
func (*DiscoveryBadgeLink) ProtoMessage()    {}
func (*DiscoveryBadgeLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{5}
}

func (m *DiscoveryBadgeLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryBadgeLink.Unmarshal(m, b)
}
func (m *DiscoveryBadgeLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryBadgeLink.Marshal(b, m, deterministic)
}
func (m *DiscoveryBadgeLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryBadgeLink.Merge(m, src)
}
func (m *DiscoveryBadgeLink) XXX_Size() int {
	return xxx_messageInfo_DiscoveryBadgeLink.Size(m)
}
func (m *DiscoveryBadgeLink) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryBadgeLink.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryBadgeLink proto.InternalMessageInfo

func (m *DiscoveryBadgeLink) GetLink() *Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *DiscoveryBadgeLink) GetUserReviewsUrl() string {
	if m != nil && m.UserReviewsUrl != nil {
		return *m.UserReviewsUrl
	}
	return ""
}

func (m *DiscoveryBadgeLink) GetCriticReviewsUrl() string {
	if m != nil && m.CriticReviewsUrl != nil {
		return *m.CriticReviewsUrl
	}
	return ""
}

type PlayerBadge struct {
	OverlayIcon          *Image   `protobuf:"bytes,1,opt,name=overlayIcon" json:"overlayIcon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerBadge) Reset()         { *m = PlayerBadge{} }
func (m *PlayerBadge) String() string { return proto.CompactTextString(m) }
func (*PlayerBadge) ProtoMessage()    {}
func (*PlayerBadge) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{6}
}

func (m *PlayerBadge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerBadge.Unmarshal(m, b)
}
func (m *PlayerBadge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerBadge.Marshal(b, m, deterministic)
}
func (m *PlayerBadge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerBadge.Merge(m, src)
}
func (m *PlayerBadge) XXX_Size() int {
	return xxx_messageInfo_PlayerBadge.Size(m)
}
func (m *PlayerBadge) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerBadge.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerBadge proto.InternalMessageInfo

func (m *PlayerBadge) GetOverlayIcon() *Image {
	if m != nil {
		return m.OverlayIcon
	}
	return nil
}

type FamilyAgeRangeBadge struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FamilyAgeRangeBadge) Reset()         { *m = FamilyAgeRangeBadge{} }
func (m *FamilyAgeRangeBadge) String() string { return proto.CompactTextString(m) }
func (*FamilyAgeRangeBadge) ProtoMessage()    {}
func (*FamilyAgeRangeBadge) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{7}
}

func (m *FamilyAgeRangeBadge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FamilyAgeRangeBadge.Unmarshal(m, b)
}
func (m *FamilyAgeRangeBadge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FamilyAgeRangeBadge.Marshal(b, m, deterministic)
}
func (m *FamilyAgeRangeBadge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FamilyAgeRangeBadge.Merge(m, src)
}
func (m *FamilyAgeRangeBadge) XXX_Size() int {
	return xxx_messageInfo_FamilyAgeRangeBadge.Size(m)
}
func (m *FamilyAgeRangeBadge) XXX_DiscardUnknown() {
	xxx_messageInfo_FamilyAgeRangeBadge.DiscardUnknown(m)
}

var xxx_messageInfo_FamilyAgeRangeBadge proto.InternalMessageInfo

type FamilyCategoryBadge struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FamilyCategoryBadge) Reset()         { *m = FamilyCategoryBadge{} }
func (m *FamilyCategoryBadge) String() string { return proto.CompactTextString(m) }
func (*FamilyCategoryBadge) ProtoMessage()    {}
func (*FamilyCategoryBadge) Descriptor() ([]byte, []int) {
	return fileDescriptor_866122eb7c2ee51a, []int{8}
}

func (m *FamilyCategoryBadge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FamilyCategoryBadge.Unmarshal(m, b)
}
func (m *FamilyCategoryBadge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FamilyCategoryBadge.Marshal(b, m, deterministic)
}
func (m *FamilyCategoryBadge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FamilyCategoryBadge.Merge(m, src)
}
func (m *FamilyCategoryBadge) XXX_Size() int {
	return xxx_messageInfo_FamilyCategoryBadge.Size(m)
}
func (m *FamilyCategoryBadge) XXX_DiscardUnknown() {
	xxx_messageInfo_FamilyCategoryBadge.DiscardUnknown(m)
}

var xxx_messageInfo_FamilyCategoryBadge proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BulkDetailsRequest)(nil), "playapi.proto.finsky.details.BulkDetailsRequest")
	proto.RegisterType((*BulkDetailsResponse)(nil), "playapi.proto.finsky.details.BulkDetailsResponse")
	proto.RegisterType((*BulkDetailsEntry)(nil), "playapi.proto.finsky.details.BulkDetailsEntry")
	proto.RegisterType((*DetailsResponse)(nil), "playapi.proto.finsky.details.DetailsResponse")
	proto.RegisterType((*DiscoveryBadge)(nil), "playapi.proto.finsky.details.DiscoveryBadge")
	proto.RegisterType((*DiscoveryBadgeLink)(nil), "playapi.proto.finsky.details.DiscoveryBadgeLink")
	proto.RegisterType((*PlayerBadge)(nil), "playapi.proto.finsky.details.PlayerBadge")
	proto.RegisterType((*FamilyAgeRangeBadge)(nil), "playapi.proto.finsky.details.FamilyAgeRangeBadge")
	proto.RegisterType((*FamilyCategoryBadge)(nil), "playapi.proto.finsky.details.FamilyCategoryBadge")
}

func init() { proto.RegisterFile("play_details.proto", fileDescriptor_866122eb7c2ee51a) }

var fileDescriptor_866122eb7c2ee51a = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xd1, 0x8e, 0x1b, 0x35,
	0x14, 0xd5, 0x6c, 0x9a, 0x76, 0x73, 0xa7, 0xdd, 0x6d, 0xbd, 0x45, 0x1a, 0x95, 0x8a, 0x86, 0xa8,
	0xa0, 0x80, 0xaa, 0x94, 0x06, 0x89, 0x07, 0x24, 0x1e, 0xd8, 0xa4, 0x88, 0x96, 0x15, 0xac, 0x0c,
	0xdd, 0x07, 0x78, 0x00, 0xaf, 0xe7, 0xee, 0xac, 0x15, 0xc7, 0x37, 0xd8, 0x9e, 0xad, 0xf2, 0x1f,
	0xbc, 0xf2, 0x67, 0x3c, 0xf0, 0x29, 0x68, 0xec, 0x59, 0x25, 0x93, 0x4c, 0xc3, 0xf6, 0xd1, 0xc7,
	0xe7, 0x1c, 0x5f, 0xdf, 0x73, 0x3d, 0x03, 0x6c, 0xa1, 0xc5, 0xf2, 0xf7, 0x1c, 0xbd, 0x50, 0xda,
	0x8d, 0x16, 0x96, 0x3c, 0xb1, 0xc7, 0x15, 0x26, 0x16, 0x2a, 0x2e, 0x47, 0x17, 0xca, 0xb8, 0xd9,
	0x72, 0x54, 0x73, 0x1e, 0x3d, 0x08, 0x0a, 0x49, 0xf3, 0x39, 0x99, 0xc8, 0x78, 0x74, 0x14, 0x4d,
	0x48, 0x96, 0x73, 0x34, 0xbe, 0x06, 0x0f, 0x03, 0xa8, 0x95, 0x99, 0x45, 0x60, 0xf0, 0x6f, 0x02,
	0xec, 0xb8, 0xd4, 0xb3, 0x69, 0x34, 0xe2, 0xf8, 0x67, 0x89, 0xce, 0xb3, 0x87, 0xd0, 0xcd, 0x49,
	0xaa, 0x3c, 0x4b, 0xfa, 0x9d, 0x61, 0x8f, 0xc7, 0x05, 0xfb, 0x1c, 0xee, 0x2b, 0x23, 0x75, 0x99,
	0xe3, 0xe4, 0x52, 0xe9, 0x7c, 0x4a, 0xd2, 0x65, 0x7b, 0xfd, 0x64, 0xb8, 0xcf, 0xb7, 0x70, 0xf6,
	0x29, 0x1c, 0xd4, 0x58, 0x6d, 0x9d, 0x75, 0x02, 0x73, 0x03, 0x65, 0xcf, 0xe0, 0x81, 0xa3, 0xd2,
	0x4a, 0x3c, 0x15, 0x72, 0x26, 0x0a, 0xfc, 0x51, 0xcc, 0x31, 0xbb, 0xd5, 0x4f, 0x86, 0x3d, 0xbe,
	0xbd, 0xc1, 0xc6, 0xf0, 0x50, 0x19, 0xe7, 0x85, 0xd6, 0x98, 0x9f, 0xa1, 0x75, 0x8a, 0xcc, 0x84,
	0x72, 0xcc, 0xee, 0xf4, 0x3b, 0xc3, 0x2e, 0x6f, 0xdd, 0x1b, 0xfc, 0x06, 0x47, 0x8d, 0x1b, 0xba,
	0x05, 0x19, 0x87, 0x6c, 0x0a, 0x5d, 0x34, 0xde, 0x2e, 0xc3, 0x15, 0xd3, 0xf1, 0x68, 0xb4, 0xab,
	0xc1, 0xa3, 0x35, 0x87, 0x97, 0x95, 0x8a, 0x47, 0xf1, 0xe0, 0x35, 0xdc, 0xdf, 0xdc, 0x62, 0x5f,
	0x41, 0x27, 0x27, 0x99, 0x25, 0xfd, 0x64, 0x98, 0x8e, 0x9f, 0xbe, 0xc3, 0xf7, 0x3a, 0x97, 0x29,
	0xc9, 0xb3, 0x31, 0xaf, 0x04, 0x83, 0x7f, 0xf6, 0xe0, 0x70, 0xb3, 0xca, 0x97, 0x00, 0xa5, 0x43,
	0xcb, 0xf1, 0x4a, 0xe1, 0xdb, 0xd0, 0xc2, 0x74, 0xfc, 0xc9, 0xff, 0x58, 0x46, 0x32, 0x5f, 0x13,
	0xb2, 0xaf, 0x43, 0x9e, 0x67, 0xe3, 0xd0, 0xd9, 0x9b, 0x16, 0x15, 0x25, 0xec, 0x23, 0x80, 0x0b,
	0x22, 0x8f, 0xf6, 0x7b, 0x3f, 0xd7, 0x59, 0x37, 0x44, 0xb3, 0x86, 0x54, 0x53, 0xe1, 0xd0, 0x5e,
	0xa1, 0x3d, 0xa1, 0xc2, 0x4d, 0x88, 0x66, 0x0a, 0xb3, 0xdb, 0xfd, 0x64, 0x78, 0x97, 0x6f, 0xe1,
	0xec, 0x17, 0x38, 0xc8, 0x95, 0x93, 0x74, 0x85, 0x76, 0x79, 0x2c, 0xf2, 0x22, 0x26, 0x97, 0x8e,
	0x9f, 0xed, 0xee, 0xfe, 0xb4, 0xa1, 0xe1, 0x1b, 0x1e, 0xec, 0x29, 0xdc, 0x43, 0x23, 0xce, 0x35,
	0xc6, 0xdb, 0xba, 0x6c, 0x3f, 0x8c, 0x5a, 0x13, 0x1c, 0xfc, 0x75, 0x1b, 0x0e, 0x9a, 0x46, 0xd5,
	0x98, 0x7b, 0xe5, 0x35, 0x86, 0xac, 0x7a, 0x3c, 0x2e, 0xd8, 0x0b, 0xe8, 0xaa, 0xb9, 0x28, 0x30,
	0xcc, 0x76, 0x3a, 0xfe, 0xb0, 0xbd, 0xb6, 0x57, 0x15, 0x85, 0x47, 0x26, 0x1b, 0xc2, 0xe1, 0xb9,
	0x90, 0xb3, 0xc2, 0x52, 0x69, 0xf2, 0x09, 0x69, 0xb2, 0x21, 0xab, 0x2e, 0xdf, 0x84, 0xd9, 0x1f,
	0xc0, 0x9a, 0xd5, 0x9f, 0x28, 0x33, 0xab, 0x63, 0xf9, 0xe2, 0x7d, 0xba, 0x50, 0xe9, 0x78, 0x8b,
	0x57, 0x6b, 0x1e, 0xdd, 0x77, 0xe4, 0xf1, 0x18, 0x7a, 0xca, 0x9d, 0xea, 0xd2, 0xfd, 0x64, 0x62,
	0x68, 0xfb, 0x7c, 0x05, 0x54, 0xb7, 0x12, 0x45, 0x61, 0xb1, 0x10, 0x1e, 0xb9, 0xf0, 0xca, 0x14,
	0xd9, 0x9d, 0x7e, 0x32, 0xdc, 0xe3, 0x9b, 0x70, 0xf5, 0xda, 0xab, 0x69, 0xfb, 0xd9, 0x0b, 0x5b,
	0x13, 0xf7, 0xc3, 0xf5, 0x37, 0xd0, 0x2a, 0xa9, 0x9c, 0xde, 0x1a, 0x4d, 0x22, 0x9f, 0x50, 0x69,
	0x7c, 0xd6, 0x0b, 0x8d, 0x6f, 0x82, 0xeb, 0xac, 0x37, 0x46, 0x79, 0x97, 0x41, 0x93, 0x15, 0x40,
	0x36, 0x02, 0x26, 0xc9, 0x78, 0x34, 0x7e, 0x8a, 0x4e, 0x5a, 0xb5, 0xf0, 0x8a, 0x4c, 0x96, 0x06,
	0x6a, 0xcb, 0x0e, 0xfb, 0x01, 0xd2, 0xaa, 0xbd, 0x68, 0xe3, 0xe0, 0xdd, 0x0d, 0x2d, 0xff, 0x6c,
	0x77, 0xcb, 0x4f, 0x57, 0x02, 0xbe, 0xae, 0x66, 0x12, 0x8e, 0x2e, 0xc4, 0x5c, 0xe9, 0xe5, 0xb7,
	0x05, 0x72, 0x61, 0x0a, 0x8c, 0xa6, 0xf7, 0x82, 0xe9, 0x8b, 0xdd, 0xa6, 0xdf, 0x6d, 0x0b, 0x79,
	0x9b, 0xdb, 0xea, 0x90, 0x89, 0xf0, 0x58, 0xd0, 0xf5, 0x93, 0x39, 0xb8, 0xf9, 0x21, 0x0d, 0x21,
	0x6f, 0x73, 0x1b, 0xfc, 0x9d, 0x00, 0xdb, 0x9e, 0x2c, 0xf6, 0x25, 0xdc, 0xaa, 0x7e, 0x13, 0xf5,
	0x57, 0xec, 0x49, 0xfb, 0x61, 0xe1, 0x47, 0x12, 0x06, 0x31, 0x90, 0xaf, 0xc7, 0xa0, 0x7e, 0x71,
	0x6f, 0xac, 0x0e, 0x4f, 0xa8, 0xc7, 0x37, 0xd0, 0x6a, 0x44, 0xa5, 0x55, 0x5e, 0xc9, 0x35, 0x66,
	0x27, 0x30, 0xb7, 0xf0, 0xc1, 0x09, 0xa4, 0x6b, 0x29, 0xb0, 0x6f, 0x20, 0xad, 0x2a, 0xd5, 0x62,
	0xf9, 0x4a, 0x92, 0xa9, 0xcb, 0xdb, 0xf9, 0x44, 0xd7, 0xf9, 0x83, 0x0f, 0xe0, 0xa8, 0xa5, 0xfd,
	0x2b, 0xb8, 0xd1, 0x9b, 0xe3, 0x8f, 0x7f, 0x7d, 0x52, 0x28, 0x7f, 0x59, 0x9e, 0x8f, 0x24, 0xcd,
	0x9f, 0xbf, 0x26, 0x77, 0x59, 0x8a, 0x29, 0xa1, 0x7b, 0xae, 0xd5, 0xb9, 0xc5, 0xea, 0xcc, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x06, 0xb7, 0xfa, 0xac, 0x07, 0x00, 0x00,
}

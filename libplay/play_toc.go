// Code generated by protoc-gen-go. DO NOT EDIT.
// source: play_toc.proto

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

type CarrierBillingConfig struct {
	Id                                         *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                                       *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ApiVersion                                 *int32   `protobuf:"varint,3,opt,name=apiVersion" json:"apiVersion,omitempty"`
	ProvisioningUrl                            *string  `protobuf:"bytes,4,opt,name=provisioningUrl" json:"provisioningUrl,omitempty"`
	CredentialsUrl                             *string  `protobuf:"bytes,5,opt,name=credentialsUrl" json:"credentialsUrl,omitempty"`
	TosRequired                                *bool    `protobuf:"varint,6,opt,name=tosRequired" json:"tosRequired,omitempty"`
	PerTransactionCredentialsRequired          *bool    `protobuf:"varint,7,opt,name=perTransactionCredentialsRequired" json:"perTransactionCredentialsRequired,omitempty"`
	SendSubscriberIdWithCarrierBillingRequests *bool    `protobuf:"varint,8,opt,name=sendSubscriberIdWithCarrierBillingRequests" json:"sendSubscriberIdWithCarrierBillingRequests,omitempty"`
	XXX_NoUnkeyedLiteral                       struct{} `json:"-"`
	XXX_unrecognized                           []byte   `json:"-"`
	XXX_sizecache                              int32    `json:"-"`
}

func (m *CarrierBillingConfig) Reset()         { *m = CarrierBillingConfig{} }
func (m *CarrierBillingConfig) String() string { return proto.CompactTextString(m) }
func (*CarrierBillingConfig) ProtoMessage()    {}
func (*CarrierBillingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_968b5b45718eafd3, []int{0}
}

func (m *CarrierBillingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarrierBillingConfig.Unmarshal(m, b)
}
func (m *CarrierBillingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarrierBillingConfig.Marshal(b, m, deterministic)
}
func (m *CarrierBillingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarrierBillingConfig.Merge(m, src)
}
func (m *CarrierBillingConfig) XXX_Size() int {
	return xxx_messageInfo_CarrierBillingConfig.Size(m)
}
func (m *CarrierBillingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CarrierBillingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CarrierBillingConfig proto.InternalMessageInfo

func (m *CarrierBillingConfig) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *CarrierBillingConfig) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CarrierBillingConfig) GetApiVersion() int32 {
	if m != nil && m.ApiVersion != nil {
		return *m.ApiVersion
	}
	return 0
}

func (m *CarrierBillingConfig) GetProvisioningUrl() string {
	if m != nil && m.ProvisioningUrl != nil {
		return *m.ProvisioningUrl
	}
	return ""
}

func (m *CarrierBillingConfig) GetCredentialsUrl() string {
	if m != nil && m.CredentialsUrl != nil {
		return *m.CredentialsUrl
	}
	return ""
}

func (m *CarrierBillingConfig) GetTosRequired() bool {
	if m != nil && m.TosRequired != nil {
		return *m.TosRequired
	}
	return false
}

func (m *CarrierBillingConfig) GetPerTransactionCredentialsRequired() bool {
	if m != nil && m.PerTransactionCredentialsRequired != nil {
		return *m.PerTransactionCredentialsRequired
	}
	return false
}

func (m *CarrierBillingConfig) GetSendSubscriberIdWithCarrierBillingRequests() bool {
	if m != nil && m.SendSubscriberIdWithCarrierBillingRequests != nil {
		return *m.SendSubscriberIdWithCarrierBillingRequests
	}
	return false
}

type BillingConfig struct {
	CarrierBillingConfig *CarrierBillingConfig `protobuf:"bytes,1,opt,name=carrierBillingConfig" json:"carrierBillingConfig,omitempty"`
	MaxIabApiVersion     *int32                `protobuf:"varint,2,opt,name=maxIabApiVersion" json:"maxIabApiVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BillingConfig) Reset()         { *m = BillingConfig{} }
func (m *BillingConfig) String() string { return proto.CompactTextString(m) }
func (*BillingConfig) ProtoMessage()    {}
func (*BillingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_968b5b45718eafd3, []int{1}
}

func (m *BillingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingConfig.Unmarshal(m, b)
}
func (m *BillingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingConfig.Marshal(b, m, deterministic)
}
func (m *BillingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingConfig.Merge(m, src)
}
func (m *BillingConfig) XXX_Size() int {
	return xxx_messageInfo_BillingConfig.Size(m)
}
func (m *BillingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BillingConfig proto.InternalMessageInfo

func (m *BillingConfig) GetCarrierBillingConfig() *CarrierBillingConfig {
	if m != nil {
		return m.CarrierBillingConfig
	}
	return nil
}

func (m *BillingConfig) GetMaxIabApiVersion() int32 {
	if m != nil && m.MaxIabApiVersion != nil {
		return *m.MaxIabApiVersion
	}
	return 0
}

type CorpusMetadata struct {
	Backend              *int32   `protobuf:"varint,1,opt,name=backend" json:"backend,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	LandingUrl           *string  `protobuf:"bytes,3,opt,name=landingUrl" json:"landingUrl,omitempty"`
	LibraryName          *string  `protobuf:"bytes,4,opt,name=libraryName" json:"libraryName,omitempty"`
	RecsWidgetUrl        *string  `protobuf:"bytes,6,opt,name=recsWidgetUrl" json:"recsWidgetUrl,omitempty"`
	ShopName             *string  `protobuf:"bytes,7,opt,name=shopName" json:"shopName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CorpusMetadata) Reset()         { *m = CorpusMetadata{} }
func (m *CorpusMetadata) String() string { return proto.CompactTextString(m) }
func (*CorpusMetadata) ProtoMessage()    {}
func (*CorpusMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_968b5b45718eafd3, []int{2}
}

func (m *CorpusMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CorpusMetadata.Unmarshal(m, b)
}
func (m *CorpusMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CorpusMetadata.Marshal(b, m, deterministic)
}
func (m *CorpusMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CorpusMetadata.Merge(m, src)
}
func (m *CorpusMetadata) XXX_Size() int {
	return xxx_messageInfo_CorpusMetadata.Size(m)
}
func (m *CorpusMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CorpusMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CorpusMetadata proto.InternalMessageInfo

func (m *CorpusMetadata) GetBackend() int32 {
	if m != nil && m.Backend != nil {
		return *m.Backend
	}
	return 0
}

func (m *CorpusMetadata) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CorpusMetadata) GetLandingUrl() string {
	if m != nil && m.LandingUrl != nil {
		return *m.LandingUrl
	}
	return ""
}

func (m *CorpusMetadata) GetLibraryName() string {
	if m != nil && m.LibraryName != nil {
		return *m.LibraryName
	}
	return ""
}

func (m *CorpusMetadata) GetRecsWidgetUrl() string {
	if m != nil && m.RecsWidgetUrl != nil {
		return *m.RecsWidgetUrl
	}
	return ""
}

func (m *CorpusMetadata) GetShopName() string {
	if m != nil && m.ShopName != nil {
		return *m.ShopName
	}
	return ""
}

type Experiments struct {
	ExperimentId         []string `protobuf:"bytes,1,rep,name=experimentId" json:"experimentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Experiments) Reset()         { *m = Experiments{} }
func (m *Experiments) String() string { return proto.CompactTextString(m) }
func (*Experiments) ProtoMessage()    {}
func (*Experiments) Descriptor() ([]byte, []int) {
	return fileDescriptor_968b5b45718eafd3, []int{3}
}

func (m *Experiments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experiments.Unmarshal(m, b)
}
func (m *Experiments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experiments.Marshal(b, m, deterministic)
}
func (m *Experiments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experiments.Merge(m, src)
}
func (m *Experiments) XXX_Size() int {
	return xxx_messageInfo_Experiments.Size(m)
}
func (m *Experiments) XXX_DiscardUnknown() {
	xxx_messageInfo_Experiments.DiscardUnknown(m)
}

var xxx_messageInfo_Experiments proto.InternalMessageInfo

func (m *Experiments) GetExperimentId() []string {
	if m != nil {
		return m.ExperimentId
	}
	return nil
}

type SelfUpdateConfig struct {
	LatestClientVersionCode *int32   `protobuf:"varint,1,opt,name=latestClientVersionCode" json:"latestClientVersionCode,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *SelfUpdateConfig) Reset()         { *m = SelfUpdateConfig{} }
func (m *SelfUpdateConfig) String() string { return proto.CompactTextString(m) }
func (*SelfUpdateConfig) ProtoMessage()    {}
func (*SelfUpdateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_968b5b45718eafd3, []int{4}
}

func (m *SelfUpdateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelfUpdateConfig.Unmarshal(m, b)
}
func (m *SelfUpdateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelfUpdateConfig.Marshal(b, m, deterministic)
}
func (m *SelfUpdateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelfUpdateConfig.Merge(m, src)
}
func (m *SelfUpdateConfig) XXX_Size() int {
	return xxx_messageInfo_SelfUpdateConfig.Size(m)
}
func (m *SelfUpdateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SelfUpdateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SelfUpdateConfig proto.InternalMessageInfo

func (m *SelfUpdateConfig) GetLatestClientVersionCode() int32 {
	if m != nil && m.LatestClientVersionCode != nil {
		return *m.LatestClientVersionCode
	}
	return 0
}

type TocResponse struct {
	Corpus                         []*CorpusMetadata     `protobuf:"bytes,1,rep,name=corpus" json:"corpus,omitempty"`
	TosVersionDeprecated           *int32                `protobuf:"varint,2,opt,name=tosVersionDeprecated" json:"tosVersionDeprecated,omitempty"`
	TosContent                     *string               `protobuf:"bytes,3,opt,name=tosContent" json:"tosContent,omitempty"`
	HomeUrl                        *string               `protobuf:"bytes,4,opt,name=homeUrl" json:"homeUrl,omitempty"`
	Experiments                    *Experiments          `protobuf:"bytes,5,opt,name=experiments" json:"experiments,omitempty"`
	TosCheckboxTextMarketingEmails *string               `protobuf:"bytes,6,opt,name=tosCheckboxTextMarketingEmails" json:"tosCheckboxTextMarketingEmails,omitempty"`
	TosToken                       *string               `protobuf:"bytes,7,opt,name=tosToken" json:"tosToken,omitempty"`
	UserSettings                   *OBSOLETEUserSettings `protobuf:"bytes,8,opt,name=userSettings" json:"userSettings,omitempty"`
	IconOverrideUrl                *string               `protobuf:"bytes,9,opt,name=iconOverrideUrl" json:"iconOverrideUrl,omitempty"`
	SelfUpdateConfig               *SelfUpdateConfig     `protobuf:"bytes,10,opt,name=selfUpdateConfig" json:"selfUpdateConfig,omitempty"`
	RequiresUploadDeviceConfig     *bool                 `protobuf:"varint,11,opt,name=requiresUploadDeviceConfig" json:"requiresUploadDeviceConfig,omitempty"`
	BillingConfig                  *BillingConfig        `protobuf:"bytes,12,opt,name=billingConfig" json:"billingConfig,omitempty"`
	RecsWidgetUrl                  *string               `protobuf:"bytes,13,opt,name=recsWidgetUrl" json:"recsWidgetUrl,omitempty"`
	SocialHomeUrl                  *string               `protobuf:"bytes,15,opt,name=socialHomeUrl" json:"socialHomeUrl,omitempty"`
	AgeVerificationRequired        *bool                 `protobuf:"varint,16,opt,name=ageVerificationRequired" json:"ageVerificationRequired,omitempty"`
	GplusSignupEnabled             *bool                 `protobuf:"varint,17,opt,name=gplusSignupEnabled" json:"gplusSignupEnabled,omitempty"`
	RedeemEnabled                  *bool                 `protobuf:"varint,18,opt,name=redeemEnabled" json:"redeemEnabled,omitempty"`
	HelpUrl                        *string               `protobuf:"bytes,19,opt,name=helpUrl" json:"helpUrl,omitempty"`
	ThemeId                        *int32                `protobuf:"varint,20,opt,name=themeId" json:"themeId,omitempty"`
	EntertainmentHomeUrl           *string               `protobuf:"bytes,21,opt,name=entertainmentHomeUrl" json:"entertainmentHomeUrl,omitempty"`
	Cookie                         *string               `protobuf:"bytes,22,opt,name=cookie" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}              `json:"-"`
	XXX_unrecognized               []byte                `json:"-"`
	XXX_sizecache                  int32                 `json:"-"`
}

func (m *TocResponse) Reset()         { *m = TocResponse{} }
func (m *TocResponse) String() string { return proto.CompactTextString(m) }
func (*TocResponse) ProtoMessage()    {}
func (*TocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_968b5b45718eafd3, []int{5}
}

func (m *TocResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TocResponse.Unmarshal(m, b)
}
func (m *TocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TocResponse.Marshal(b, m, deterministic)
}
func (m *TocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TocResponse.Merge(m, src)
}
func (m *TocResponse) XXX_Size() int {
	return xxx_messageInfo_TocResponse.Size(m)
}
func (m *TocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TocResponse proto.InternalMessageInfo

func (m *TocResponse) GetCorpus() []*CorpusMetadata {
	if m != nil {
		return m.Corpus
	}
	return nil
}

func (m *TocResponse) GetTosVersionDeprecated() int32 {
	if m != nil && m.TosVersionDeprecated != nil {
		return *m.TosVersionDeprecated
	}
	return 0
}

func (m *TocResponse) GetTosContent() string {
	if m != nil && m.TosContent != nil {
		return *m.TosContent
	}
	return ""
}

func (m *TocResponse) GetHomeUrl() string {
	if m != nil && m.HomeUrl != nil {
		return *m.HomeUrl
	}
	return ""
}

func (m *TocResponse) GetExperiments() *Experiments {
	if m != nil {
		return m.Experiments
	}
	return nil
}

func (m *TocResponse) GetTosCheckboxTextMarketingEmails() string {
	if m != nil && m.TosCheckboxTextMarketingEmails != nil {
		return *m.TosCheckboxTextMarketingEmails
	}
	return ""
}

func (m *TocResponse) GetTosToken() string {
	if m != nil && m.TosToken != nil {
		return *m.TosToken
	}
	return ""
}

func (m *TocResponse) GetUserSettings() *OBSOLETEUserSettings {
	if m != nil {
		return m.UserSettings
	}
	return nil
}

func (m *TocResponse) GetIconOverrideUrl() string {
	if m != nil && m.IconOverrideUrl != nil {
		return *m.IconOverrideUrl
	}
	return ""
}

func (m *TocResponse) GetSelfUpdateConfig() *SelfUpdateConfig {
	if m != nil {
		return m.SelfUpdateConfig
	}
	return nil
}

func (m *TocResponse) GetRequiresUploadDeviceConfig() bool {
	if m != nil && m.RequiresUploadDeviceConfig != nil {
		return *m.RequiresUploadDeviceConfig
	}
	return false
}

func (m *TocResponse) GetBillingConfig() *BillingConfig {
	if m != nil {
		return m.BillingConfig
	}
	return nil
}

func (m *TocResponse) GetRecsWidgetUrl() string {
	if m != nil && m.RecsWidgetUrl != nil {
		return *m.RecsWidgetUrl
	}
	return ""
}

func (m *TocResponse) GetSocialHomeUrl() string {
	if m != nil && m.SocialHomeUrl != nil {
		return *m.SocialHomeUrl
	}
	return ""
}

func (m *TocResponse) GetAgeVerificationRequired() bool {
	if m != nil && m.AgeVerificationRequired != nil {
		return *m.AgeVerificationRequired
	}
	return false
}

func (m *TocResponse) GetGplusSignupEnabled() bool {
	if m != nil && m.GplusSignupEnabled != nil {
		return *m.GplusSignupEnabled
	}
	return false
}

func (m *TocResponse) GetRedeemEnabled() bool {
	if m != nil && m.RedeemEnabled != nil {
		return *m.RedeemEnabled
	}
	return false
}

func (m *TocResponse) GetHelpUrl() string {
	if m != nil && m.HelpUrl != nil {
		return *m.HelpUrl
	}
	return ""
}

func (m *TocResponse) GetThemeId() int32 {
	if m != nil && m.ThemeId != nil {
		return *m.ThemeId
	}
	return 0
}

func (m *TocResponse) GetEntertainmentHomeUrl() string {
	if m != nil && m.EntertainmentHomeUrl != nil {
		return *m.EntertainmentHomeUrl
	}
	return ""
}

func (m *TocResponse) GetCookie() string {
	if m != nil && m.Cookie != nil {
		return *m.Cookie
	}
	return ""
}

func init() {
	proto.RegisterType((*CarrierBillingConfig)(nil), "playapi.proto.finsky.toc.CarrierBillingConfig")
	proto.RegisterType((*BillingConfig)(nil), "playapi.proto.finsky.toc.BillingConfig")
	proto.RegisterType((*CorpusMetadata)(nil), "playapi.proto.finsky.toc.CorpusMetadata")
	proto.RegisterType((*Experiments)(nil), "playapi.proto.finsky.toc.Experiments")
	proto.RegisterType((*SelfUpdateConfig)(nil), "playapi.proto.finsky.toc.SelfUpdateConfig")
	proto.RegisterType((*TocResponse)(nil), "playapi.proto.finsky.toc.TocResponse")
}

func init() { proto.RegisterFile("play_toc.proto", fileDescriptor_968b5b45718eafd3) }

var fileDescriptor_968b5b45718eafd3 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x96, 0x9d, 0xcd, 0x5f, 0x4d, 0x9c, 0x0d, 0xbd, 0x61, 0x19, 0xe5, 0xb0, 0x64, 0xad, 0x05,
	0xac, 0x1c, 0xbc, 0xc2, 0x5c, 0x38, 0x21, 0x88, 0x63, 0x20, 0x28, 0x21, 0xd2, 0xd8, 0x49, 0x24,
	0x0e, 0xa0, 0x9e, 0xe9, 0xca, 0xb8, 0xe4, 0x71, 0xf7, 0xd0, 0xdd, 0x8e, 0x9c, 0xa7, 0xe1, 0x15,
	0x78, 0x06, 0x9e, 0x8b, 0x03, 0xea, 0x9e, 0xf1, 0xcf, 0x38, 0xf6, 0xae, 0xb8, 0x4d, 0x7d, 0xf5,
	0xd3, 0x35, 0xf5, 0x55, 0x7d, 0x70, 0x98, 0x67, 0xfc, 0xe9, 0x0f, 0xab, 0x92, 0x76, 0xae, 0x95,
	0x55, 0x2c, 0x74, 0x36, 0xcf, 0xa9, 0x30, 0xdb, 0x0f, 0x24, 0xcd, 0xe8, 0xa9, 0x6d, 0x55, 0x72,
	0xf2, 0xca, 0x47, 0x1a, 0xb4, 0x96, 0x64, 0x6a, 0x0a, 0x7f, 0xf3, 0xdf, 0x3a, 0x1c, 0x77, 0xb9,
	0xd6, 0x84, 0xfa, 0x9c, 0xb2, 0x8c, 0x64, 0xda, 0x55, 0xf2, 0x81, 0x52, 0x76, 0x08, 0x75, 0x12,
	0x61, 0xed, 0xb4, 0xd6, 0xda, 0x8f, 0xea, 0x24, 0x18, 0x83, 0x17, 0x92, 0x8f, 0x31, 0xac, 0x7b,
	0xc4, 0x7f, 0xb3, 0x37, 0x00, 0x3c, 0xa7, 0x3b, 0xd4, 0x86, 0x94, 0x0c, 0xb7, 0x4e, 0x6b, 0xad,
	0xed, 0x68, 0x09, 0x61, 0x2d, 0x78, 0x99, 0x6b, 0xf5, 0x48, 0xce, 0x20, 0x99, 0xde, 0xea, 0x2c,
	0x7c, 0xe1, 0xd3, 0x57, 0x61, 0xf6, 0x25, 0x1c, 0x26, 0x1a, 0x05, 0x4a, 0x4b, 0x3c, 0x33, 0x2e,
	0x70, 0xdb, 0x07, 0xae, 0xa0, 0xec, 0x14, 0x02, 0xab, 0x4c, 0x84, 0x7f, 0x4e, 0x48, 0xa3, 0x08,
	0x77, 0x4e, 0x6b, 0xad, 0xbd, 0x68, 0x19, 0x62, 0x57, 0xf0, 0x36, 0x47, 0x3d, 0xd0, 0x5c, 0x1a,
	0x9e, 0x58, 0x52, 0xb2, 0xbb, 0xa8, 0x30, 0xcf, 0xdb, 0xf5, 0x79, 0x1f, 0x0f, 0x64, 0xbf, 0xc3,
	0x99, 0x41, 0x29, 0xfa, 0x93, 0xd8, 0x24, 0x9a, 0x62, 0xd4, 0x97, 0xe2, 0x9e, 0xec, 0xb0, 0x3a,
	0x31, 0x17, 0x8d, 0xc6, 0x9a, 0x70, 0xcf, 0x97, 0xfd, 0x1f, 0x19, 0xcd, 0xbf, 0x6a, 0xd0, 0xa8,
	0xce, 0x3d, 0x86, 0xe3, 0x64, 0x0d, 0x1f, 0x9e, 0x89, 0xa0, 0xd3, 0x6e, 0x6f, 0xa2, 0xb7, 0xbd,
	0x8e, 0xc5, 0x68, 0x6d, 0x2d, 0x76, 0x06, 0x47, 0x63, 0x3e, 0xbd, 0xe4, 0xf1, 0x0f, 0x0b, 0xf6,
	0xea, 0x9e, 0xbd, 0x67, 0x78, 0xf3, 0x9f, 0x1a, 0x1c, 0x76, 0x95, 0xce, 0x27, 0xe6, 0x1a, 0x2d,
	0x17, 0xdc, 0x72, 0x16, 0xc2, 0x6e, 0xcc, 0x93, 0x11, 0xca, 0x62, 0x3f, 0xb6, 0xa3, 0x99, 0xb9,
	0x69, 0x49, 0x32, 0x2e, 0x45, 0xc9, 0xff, 0x96, 0xf7, 0x2c, 0x21, 0x8e, 0xd2, 0x8c, 0x62, 0xcd,
	0xf5, 0xd3, 0xaf, 0x2e, 0xb5, 0x58, 0x90, 0x65, 0x88, 0xbd, 0x83, 0x86, 0xc6, 0xc4, 0xdc, 0x93,
	0x48, 0xd1, 0xba, 0x22, 0x3b, 0x3e, 0xa6, 0x0a, 0xb2, 0x13, 0xd8, 0x33, 0x43, 0x95, 0xfb, 0x22,
	0xbb, 0x3e, 0x60, 0x6e, 0x37, 0xbf, 0x86, 0xa0, 0x37, 0xcd, 0x51, 0xd3, 0x18, 0xa5, 0x35, 0xac,
	0x09, 0x07, 0x38, 0x37, 0x2f, 0xdd, 0x5f, 0x6c, 0xb5, 0xf6, 0xa3, 0x0a, 0xd6, 0xbc, 0x82, 0xa3,
	0x3e, 0x66, 0x0f, 0xb7, 0xb9, 0xe0, 0x16, 0xcb, 0xb9, 0x7d, 0x0b, 0x9f, 0x65, 0xdc, 0xa2, 0xb1,
	0xdd, 0x8c, 0x50, 0xda, 0x72, 0x44, 0x5d, 0x25, 0xb0, 0x1c, 0xc4, 0x26, 0x77, 0xf3, 0xef, 0x3d,
	0x08, 0x06, 0x2a, 0x89, 0xd0, 0xe4, 0x4a, 0x1a, 0x64, 0xdf, 0xc3, 0x4e, 0xe2, 0x87, 0xea, 0xdf,
	0x0e, 0x3a, 0xad, 0x0f, 0xf0, 0x5a, 0x19, 0x7e, 0x54, 0xe6, 0xb1, 0x0e, 0x1c, 0x5b, 0x65, 0xca,
	0x37, 0x2e, 0x30, 0xd7, 0x98, 0x70, 0x8b, 0xa2, 0xe4, 0x71, 0xad, 0xcf, 0x51, 0x61, 0x95, 0xe9,
	0x2a, 0x69, 0x51, 0xda, 0x19, 0x15, 0x0b, 0xc4, 0x11, 0x3b, 0x54, 0x63, 0x5c, 0xdc, 0xe9, 0xcc,
	0x64, 0x3f, 0x41, 0xb0, 0x98, 0x8e, 0xf1, 0xc7, 0x19, 0x74, 0xbe, 0xd8, 0xdc, 0xf4, 0xd2, 0xb4,
	0xa3, 0xe5, 0x4c, 0xf6, 0x23, 0xbc, 0x71, 0x0f, 0x0e, 0x31, 0x19, 0xc5, 0x6a, 0x3a, 0xc0, 0xa9,
	0xbd, 0xe6, 0x7a, 0x84, 0x4e, 0x92, 0x7a, 0x63, 0x4e, 0x99, 0x29, 0xc9, 0xfd, 0x48, 0x94, 0x63,
	0xdb, 0x2a, 0x33, 0x50, 0x23, 0x94, 0x33, 0xb6, 0x67, 0x36, 0xbb, 0x87, 0x83, 0x89, 0x41, 0xdd,
	0x2f, 0x95, 0xce, 0x9f, 0x65, 0xd0, 0xf9, 0x66, 0x7d, 0xb7, 0x73, 0x3d, 0xbc, 0x39, 0xef, 0xdf,
	0x5c, 0xf5, 0x06, 0xbd, 0xdb, 0xa5, 0xd4, 0xa8, 0x52, 0xc8, 0xe9, 0x19, 0x25, 0x4a, 0xde, 0x3c,
	0xa2, 0xd6, 0x24, 0xfc, 0x9c, 0xf6, 0x0b, 0x3d, 0x5b, 0x81, 0xd9, 0x1d, 0x1c, 0x99, 0x95, 0xed,
	0x09, 0xc1, 0xb7, 0x71, 0xb6, 0x79, 0x68, 0xab, 0xfb, 0x16, 0x3d, 0xab, 0xc1, 0xbe, 0x83, 0x13,
	0x5d, 0x68, 0x93, 0xb9, 0xcd, 0x33, 0xc5, 0xc5, 0x05, 0x3e, 0x52, 0x32, 0x7b, 0x21, 0xf0, 0xfa,
	0xf3, 0x81, 0x08, 0x76, 0x0d, 0x8d, 0xb8, 0x22, 0x2b, 0x07, 0xbe, 0xa9, 0xaf, 0x36, 0x37, 0x55,
	0xd5, 0x93, 0x6a, 0xf6, 0xf3, 0xcb, 0x6c, 0xac, 0xbb, 0xcc, 0x77, 0xd0, 0x30, 0x2a, 0x21, 0x9e,
	0xfd, 0x5c, 0x2e, 0xd7, 0xcb, 0x22, 0xaa, 0x02, 0xba, 0xe3, 0xe2, 0x29, 0xde, 0xa1, 0xa6, 0x07,
	0x4a, 0xb8, 0x13, 0xe4, 0xb9, 0x5c, 0x1f, 0xf9, 0xff, 0xda, 0xe4, 0x66, 0x6d, 0x60, 0x69, 0x9e,
	0x4d, 0x4c, 0x9f, 0x52, 0x39, 0xc9, 0x7b, 0x92, 0xc7, 0x19, 0x8a, 0xf0, 0x13, 0x9f, 0xb4, 0xc6,
	0x53, 0x74, 0x2d, 0x10, 0xc7, 0xb3, 0x50, 0xe6, 0x43, 0xab, 0xa0, 0x3f, 0x06, 0xcc, 0x72, 0xd7,
	0xef, 0xab, 0xf2, 0x18, 0x0a, 0xd3, 0x79, 0xec, 0x10, 0xc7, 0x78, 0x29, 0xc2, 0xe3, 0x42, 0xff,
	0x4a, 0xd3, 0x1d, 0x25, 0x4a, 0x8b, 0xda, 0x72, 0x92, 0x6e, 0xdf, 0x67, 0x3f, 0xfc, 0xa9, 0x2f,
	0xb0, 0xd6, 0xc7, 0x5e, 0x3b, 0x29, 0x50, 0x23, 0xc2, 0xf0, 0xb5, 0x8f, 0x2a, 0xad, 0xf3, 0xb7,
	0xbf, 0x7d, 0x9e, 0x92, 0x1d, 0x4e, 0xe2, 0x76, 0xa2, 0xc6, 0xef, 0x7f, 0x51, 0x66, 0x38, 0xe1,
	0x17, 0x0a, 0xcd, 0x7b, 0x27, 0x8d, 0xe8, 0xf8, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x32,
	0x72, 0x7f, 0xfc, 0x07, 0x00, 0x00,
}

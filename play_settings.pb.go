// Code generated by protoc-gen-go. DO NOT EDIT.
// source: play_settings.proto

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

type UserSettingsConsistencyTokens struct {
	ConsistencyTokenInfo []*UserSettingsConsistencyTokens_ConsistencyTokenInfo `protobuf:"bytes,1,rep,name=consistencyTokenInfo" json:"consistencyTokenInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                              `json:"-"`
	XXX_unrecognized     []byte                                                `json:"-"`
	XXX_sizecache        int32                                                 `json:"-"`
}

func (m *UserSettingsConsistencyTokens) Reset()         { *m = UserSettingsConsistencyTokens{} }
func (m *UserSettingsConsistencyTokens) String() string { return proto.CompactTextString(m) }
func (*UserSettingsConsistencyTokens) ProtoMessage()    {}
func (*UserSettingsConsistencyTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{0}
}

func (m *UserSettingsConsistencyTokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettingsConsistencyTokens.Unmarshal(m, b)
}
func (m *UserSettingsConsistencyTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettingsConsistencyTokens.Marshal(b, m, deterministic)
}
func (m *UserSettingsConsistencyTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingsConsistencyTokens.Merge(m, src)
}
func (m *UserSettingsConsistencyTokens) XXX_Size() int {
	return xxx_messageInfo_UserSettingsConsistencyTokens.Size(m)
}
func (m *UserSettingsConsistencyTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingsConsistencyTokens.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingsConsistencyTokens proto.InternalMessageInfo

func (m *UserSettingsConsistencyTokens) GetConsistencyTokenInfo() []*UserSettingsConsistencyTokens_ConsistencyTokenInfo {
	if m != nil {
		return m.ConsistencyTokenInfo
	}
	return nil
}

type UserSettingsConsistencyTokens_ConsistencyTokenInfo struct {
	RequestHeader        *string  `protobuf:"bytes,1,opt,name=requestHeader" json:"requestHeader,omitempty"`
	ConsistencyToken     *string  `protobuf:"bytes,2,opt,name=consistencyToken" json:"consistencyToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) Reset() {
	*m = UserSettingsConsistencyTokens_ConsistencyTokenInfo{}
}
func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) String() string {
	return proto.CompactTextString(m)
}
func (*UserSettingsConsistencyTokens_ConsistencyTokenInfo) ProtoMessage() {}
func (*UserSettingsConsistencyTokens_ConsistencyTokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{0, 0}
}

func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettingsConsistencyTokens_ConsistencyTokenInfo.Unmarshal(m, b)
}
func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettingsConsistencyTokens_ConsistencyTokenInfo.Marshal(b, m, deterministic)
}
func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingsConsistencyTokens_ConsistencyTokenInfo.Merge(m, src)
}
func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) XXX_Size() int {
	return xxx_messageInfo_UserSettingsConsistencyTokens_ConsistencyTokenInfo.Size(m)
}
func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingsConsistencyTokens_ConsistencyTokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingsConsistencyTokens_ConsistencyTokenInfo proto.InternalMessageInfo

func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) GetRequestHeader() string {
	if m != nil && m.RequestHeader != nil {
		return *m.RequestHeader
	}
	return ""
}

func (m *UserSettingsConsistencyTokens_ConsistencyTokenInfo) GetConsistencyToken() string {
	if m != nil && m.ConsistencyToken != nil {
		return *m.ConsistencyToken
	}
	return ""
}

type MarketingSettings struct {
	MarketingEmailsOptedIn *bool    `protobuf:"varint,1,opt,name=marketingEmailsOptedIn" json:"marketingEmailsOptedIn,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *MarketingSettings) Reset()         { *m = MarketingSettings{} }
func (m *MarketingSettings) String() string { return proto.CompactTextString(m) }
func (*MarketingSettings) ProtoMessage()    {}
func (*MarketingSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{1}
}

func (m *MarketingSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketingSettings.Unmarshal(m, b)
}
func (m *MarketingSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketingSettings.Marshal(b, m, deterministic)
}
func (m *MarketingSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketingSettings.Merge(m, src)
}
func (m *MarketingSettings) XXX_Size() int {
	return xxx_messageInfo_MarketingSettings.Size(m)
}
func (m *MarketingSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketingSettings.DiscardUnknown(m)
}

var xxx_messageInfo_MarketingSettings proto.InternalMessageInfo

func (m *MarketingSettings) GetMarketingEmailsOptedIn() bool {
	if m != nil && m.MarketingEmailsOptedIn != nil {
		return *m.MarketingEmailsOptedIn
	}
	return false
}

type PrivacySetting struct {
	Type                 *int32   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	CurrentStatus        *int32   `protobuf:"varint,2,opt,name=currentStatus" json:"currentStatus,omitempty"`
	EnabledByDefault     *bool    `protobuf:"varint,3,opt,name=enabledByDefault" json:"enabledByDefault,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivacySetting) Reset()         { *m = PrivacySetting{} }
func (m *PrivacySetting) String() string { return proto.CompactTextString(m) }
func (*PrivacySetting) ProtoMessage()    {}
func (*PrivacySetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{2}
}

func (m *PrivacySetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivacySetting.Unmarshal(m, b)
}
func (m *PrivacySetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivacySetting.Marshal(b, m, deterministic)
}
func (m *PrivacySetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivacySetting.Merge(m, src)
}
func (m *PrivacySetting) XXX_Size() int {
	return xxx_messageInfo_PrivacySetting.Size(m)
}
func (m *PrivacySetting) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivacySetting.DiscardUnknown(m)
}

var xxx_messageInfo_PrivacySetting proto.InternalMessageInfo

func (m *PrivacySetting) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *PrivacySetting) GetCurrentStatus() int32 {
	if m != nil && m.CurrentStatus != nil {
		return *m.CurrentStatus
	}
	return 0
}

func (m *PrivacySetting) GetEnabledByDefault() bool {
	if m != nil && m.EnabledByDefault != nil {
		return *m.EnabledByDefault
	}
	return false
}

type PrivacySettings struct {
	PrivacySetting       []*PrivacySetting `protobuf:"bytes,1,rep,name=privacySetting" json:"privacySetting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PrivacySettings) Reset()         { *m = PrivacySettings{} }
func (m *PrivacySettings) String() string { return proto.CompactTextString(m) }
func (*PrivacySettings) ProtoMessage()    {}
func (*PrivacySettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{3}
}

func (m *PrivacySettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivacySettings.Unmarshal(m, b)
}
func (m *PrivacySettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivacySettings.Marshal(b, m, deterministic)
}
func (m *PrivacySettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivacySettings.Merge(m, src)
}
func (m *PrivacySettings) XXX_Size() int {
	return xxx_messageInfo_PrivacySettings.Size(m)
}
func (m *PrivacySettings) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivacySettings.DiscardUnknown(m)
}

var xxx_messageInfo_PrivacySettings proto.InternalMessageInfo

func (m *PrivacySettings) GetPrivacySetting() []*PrivacySetting {
	if m != nil {
		return m.PrivacySetting
	}
	return nil
}

type FamilyInfo struct {
	FamilyMembershipStatus *int32   `protobuf:"varint,1,opt,name=familyMembershipStatus" json:"familyMembershipStatus,omitempty"`
	Family                 *Family  `protobuf:"bytes,2,opt,name=family" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *FamilyInfo) Reset()         { *m = FamilyInfo{} }
func (m *FamilyInfo) String() string { return proto.CompactTextString(m) }
func (*FamilyInfo) ProtoMessage()    {}
func (*FamilyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{4}
}

func (m *FamilyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FamilyInfo.Unmarshal(m, b)
}
func (m *FamilyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FamilyInfo.Marshal(b, m, deterministic)
}
func (m *FamilyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FamilyInfo.Merge(m, src)
}
func (m *FamilyInfo) XXX_Size() int {
	return xxx_messageInfo_FamilyInfo.Size(m)
}
func (m *FamilyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FamilyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FamilyInfo proto.InternalMessageInfo

func (m *FamilyInfo) GetFamilyMembershipStatus() int32 {
	if m != nil && m.FamilyMembershipStatus != nil {
		return *m.FamilyMembershipStatus
	}
	return 0
}

func (m *FamilyInfo) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

type Family struct {
	Member               []*FamilyMember `protobuf:"bytes,1,rep,name=member" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Family) Reset()         { *m = Family{} }
func (m *Family) String() string { return proto.CompactTextString(m) }
func (*Family) ProtoMessage()    {}
func (*Family) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{5}
}

func (m *Family) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Family.Unmarshal(m, b)
}
func (m *Family) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Family.Marshal(b, m, deterministic)
}
func (m *Family) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Family.Merge(m, src)
}
func (m *Family) XXX_Size() int {
	return xxx_messageInfo_Family.Size(m)
}
func (m *Family) XXX_DiscardUnknown() {
	xxx_messageInfo_Family.DiscardUnknown(m)
}

var xxx_messageInfo_Family proto.InternalMessageInfo

func (m *Family) GetMember() []*FamilyMember {
	if m != nil {
		return m.Member
	}
	return nil
}

type FamilyMember struct {
	Role                 *int32   `protobuf:"varint,1,opt,name=role" json:"role,omitempty"`
	PersonDocument       *DocV2   `protobuf:"bytes,2,opt,name=personDocument" json:"personDocument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FamilyMember) Reset()         { *m = FamilyMember{} }
func (m *FamilyMember) String() string { return proto.CompactTextString(m) }
func (*FamilyMember) ProtoMessage()    {}
func (*FamilyMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{6}
}

func (m *FamilyMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FamilyMember.Unmarshal(m, b)
}
func (m *FamilyMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FamilyMember.Marshal(b, m, deterministic)
}
func (m *FamilyMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FamilyMember.Merge(m, src)
}
func (m *FamilyMember) XXX_Size() int {
	return xxx_messageInfo_FamilyMember.Size(m)
}
func (m *FamilyMember) XXX_DiscardUnknown() {
	xxx_messageInfo_FamilyMember.DiscardUnknown(m)
}

var xxx_messageInfo_FamilyMember proto.InternalMessageInfo

func (m *FamilyMember) GetRole() int32 {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return 0
}

func (m *FamilyMember) GetPersonDocument() *DocV2 {
	if m != nil {
		return m.PersonDocument
	}
	return nil
}

type Onboarding struct {
	Type                 *int32   `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Onboarding) Reset()         { *m = Onboarding{} }
func (m *Onboarding) String() string { return proto.CompactTextString(m) }
func (*Onboarding) ProtoMessage()    {}
func (*Onboarding) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{7}
}

func (m *Onboarding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Onboarding.Unmarshal(m, b)
}
func (m *Onboarding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Onboarding.Marshal(b, m, deterministic)
}
func (m *Onboarding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Onboarding.Merge(m, src)
}
func (m *Onboarding) XXX_Size() int {
	return xxx_messageInfo_Onboarding.Size(m)
}
func (m *Onboarding) XXX_DiscardUnknown() {
	xxx_messageInfo_Onboarding.DiscardUnknown(m)
}

var xxx_messageInfo_Onboarding proto.InternalMessageInfo

func (m *Onboarding) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type Onboardings struct {
	Onboarding           []*Onboarding `protobuf:"bytes,1,rep,name=onboarding" json:"onboarding,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Onboardings) Reset()         { *m = Onboardings{} }
func (m *Onboardings) String() string { return proto.CompactTextString(m) }
func (*Onboardings) ProtoMessage()    {}
func (*Onboardings) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{8}
}

func (m *Onboardings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Onboardings.Unmarshal(m, b)
}
func (m *Onboardings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Onboardings.Marshal(b, m, deterministic)
}
func (m *Onboardings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Onboardings.Merge(m, src)
}
func (m *Onboardings) XXX_Size() int {
	return xxx_messageInfo_Onboardings.Size(m)
}
func (m *Onboardings) XXX_DiscardUnknown() {
	xxx_messageInfo_Onboardings.DiscardUnknown(m)
}

var xxx_messageInfo_Onboardings proto.InternalMessageInfo

func (m *Onboardings) GetOnboarding() []*Onboarding {
	if m != nil {
		return m.Onboarding
	}
	return nil
}

type UserSettings struct {
	MarketingSettings    *MarketingSettings `protobuf:"bytes,1,opt,name=marketingSettings" json:"marketingSettings,omitempty"`
	PrivacySettings      *PrivacySettings   `protobuf:"bytes,2,opt,name=privacySettings" json:"privacySettings,omitempty"`
	FamilyInfo           *FamilyInfo        `protobuf:"bytes,3,opt,name=familyInfo" json:"familyInfo,omitempty"`
	DismissedOnboardings *Onboardings       `protobuf:"bytes,4,opt,name=dismissedOnboardings" json:"dismissedOnboardings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UserSettings) Reset()         { *m = UserSettings{} }
func (m *UserSettings) String() string { return proto.CompactTextString(m) }
func (*UserSettings) ProtoMessage()    {}
func (*UserSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{9}
}

func (m *UserSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettings.Unmarshal(m, b)
}
func (m *UserSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettings.Marshal(b, m, deterministic)
}
func (m *UserSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettings.Merge(m, src)
}
func (m *UserSettings) XXX_Size() int {
	return xxx_messageInfo_UserSettings.Size(m)
}
func (m *UserSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettings.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettings proto.InternalMessageInfo

func (m *UserSettings) GetMarketingSettings() *MarketingSettings {
	if m != nil {
		return m.MarketingSettings
	}
	return nil
}

func (m *UserSettings) GetPrivacySettings() *PrivacySettings {
	if m != nil {
		return m.PrivacySettings
	}
	return nil
}

func (m *UserSettings) GetFamilyInfo() *FamilyInfo {
	if m != nil {
		return m.FamilyInfo
	}
	return nil
}

func (m *UserSettings) GetDismissedOnboardings() *Onboardings {
	if m != nil {
		return m.DismissedOnboardings
	}
	return nil
}

type OBSOLETEUserSettings struct {
	TosCheckboxMarketingEmailsOptedIn *bool             `protobuf:"varint,1,opt,name=tosCheckboxMarketingEmailsOptedIn" json:"tosCheckboxMarketingEmailsOptedIn,omitempty"`
	PrivacySetting                    []*PrivacySetting `protobuf:"bytes,2,rep,name=privacySetting" json:"privacySetting,omitempty"`
	XXX_NoUnkeyedLiteral              struct{}          `json:"-"`
	XXX_unrecognized                  []byte            `json:"-"`
	XXX_sizecache                     int32             `json:"-"`
}

func (m *OBSOLETEUserSettings) Reset()         { *m = OBSOLETEUserSettings{} }
func (m *OBSOLETEUserSettings) String() string { return proto.CompactTextString(m) }
func (*OBSOLETEUserSettings) ProtoMessage()    {}
func (*OBSOLETEUserSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{10}
}

func (m *OBSOLETEUserSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OBSOLETEUserSettings.Unmarshal(m, b)
}
func (m *OBSOLETEUserSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OBSOLETEUserSettings.Marshal(b, m, deterministic)
}
func (m *OBSOLETEUserSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OBSOLETEUserSettings.Merge(m, src)
}
func (m *OBSOLETEUserSettings) XXX_Size() int {
	return xxx_messageInfo_OBSOLETEUserSettings.Size(m)
}
func (m *OBSOLETEUserSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_OBSOLETEUserSettings.DiscardUnknown(m)
}

var xxx_messageInfo_OBSOLETEUserSettings proto.InternalMessageInfo

func (m *OBSOLETEUserSettings) GetTosCheckboxMarketingEmailsOptedIn() bool {
	if m != nil && m.TosCheckboxMarketingEmailsOptedIn != nil {
		return *m.TosCheckboxMarketingEmailsOptedIn
	}
	return false
}

func (m *OBSOLETEUserSettings) GetPrivacySetting() []*PrivacySetting {
	if m != nil {
		return m.PrivacySetting
	}
	return nil
}

type UserSettingDirtyData struct {
	Type                 *int32                         `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	ConsistencyTokens    *UserSettingsConsistencyTokens `protobuf:"bytes,2,opt,name=consistencyTokens" json:"consistencyTokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UserSettingDirtyData) Reset()         { *m = UserSettingDirtyData{} }
func (m *UserSettingDirtyData) String() string { return proto.CompactTextString(m) }
func (*UserSettingDirtyData) ProtoMessage()    {}
func (*UserSettingDirtyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{11}
}

func (m *UserSettingDirtyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSettingDirtyData.Unmarshal(m, b)
}
func (m *UserSettingDirtyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSettingDirtyData.Marshal(b, m, deterministic)
}
func (m *UserSettingDirtyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingDirtyData.Merge(m, src)
}
func (m *UserSettingDirtyData) XXX_Size() int {
	return xxx_messageInfo_UserSettingDirtyData.Size(m)
}
func (m *UserSettingDirtyData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingDirtyData.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingDirtyData proto.InternalMessageInfo

func (m *UserSettingDirtyData) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *UserSettingDirtyData) GetConsistencyTokens() *UserSettingsConsistencyTokens {
	if m != nil {
		return m.ConsistencyTokens
	}
	return nil
}

type GetUserSettingsResponse struct {
	UserSettings         *UserSettings                  `protobuf:"bytes,1,opt,name=userSettings" json:"userSettings,omitempty"`
	ConsistencyTokens    *UserSettingsConsistencyTokens `protobuf:"bytes,2,opt,name=consistencyTokens" json:"consistencyTokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetUserSettingsResponse) Reset()         { *m = GetUserSettingsResponse{} }
func (m *GetUserSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserSettingsResponse) ProtoMessage()    {}
func (*GetUserSettingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_698b561ffae0b126, []int{12}
}

func (m *GetUserSettingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserSettingsResponse.Unmarshal(m, b)
}
func (m *GetUserSettingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserSettingsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserSettingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserSettingsResponse.Merge(m, src)
}
func (m *GetUserSettingsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserSettingsResponse.Size(m)
}
func (m *GetUserSettingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserSettingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserSettingsResponse proto.InternalMessageInfo

func (m *GetUserSettingsResponse) GetUserSettings() *UserSettings {
	if m != nil {
		return m.UserSettings
	}
	return nil
}

func (m *GetUserSettingsResponse) GetConsistencyTokens() *UserSettingsConsistencyTokens {
	if m != nil {
		return m.ConsistencyTokens
	}
	return nil
}

func init() {
	proto.RegisterType((*UserSettingsConsistencyTokens)(nil), "playapi.proto.finsky.settings.UserSettingsConsistencyTokens")
	proto.RegisterType((*UserSettingsConsistencyTokens_ConsistencyTokenInfo)(nil), "playapi.proto.finsky.settings.UserSettingsConsistencyTokens.ConsistencyTokenInfo")
	proto.RegisterType((*MarketingSettings)(nil), "playapi.proto.finsky.settings.MarketingSettings")
	proto.RegisterType((*PrivacySetting)(nil), "playapi.proto.finsky.settings.PrivacySetting")
	proto.RegisterType((*PrivacySettings)(nil), "playapi.proto.finsky.settings.PrivacySettings")
	proto.RegisterType((*FamilyInfo)(nil), "playapi.proto.finsky.settings.FamilyInfo")
	proto.RegisterType((*Family)(nil), "playapi.proto.finsky.settings.Family")
	proto.RegisterType((*FamilyMember)(nil), "playapi.proto.finsky.settings.FamilyMember")
	proto.RegisterType((*Onboarding)(nil), "playapi.proto.finsky.settings.Onboarding")
	proto.RegisterType((*Onboardings)(nil), "playapi.proto.finsky.settings.Onboardings")
	proto.RegisterType((*UserSettings)(nil), "playapi.proto.finsky.settings.UserSettings")
	proto.RegisterType((*OBSOLETEUserSettings)(nil), "playapi.proto.finsky.settings.OBSOLETEUserSettings")
	proto.RegisterType((*UserSettingDirtyData)(nil), "playapi.proto.finsky.settings.UserSettingDirtyData")
	proto.RegisterType((*GetUserSettingsResponse)(nil), "playapi.proto.finsky.settings.GetUserSettingsResponse")
}

func init() { proto.RegisterFile("play_settings.proto", fileDescriptor_698b561ffae0b126) }

var fileDescriptor_698b561ffae0b126 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x96, 0xd3, 0x8b, 0xfe, 0xff, 0xa4, 0xb4, 0x64, 0x88, 0x4a, 0x54, 0xa9, 0xa2, 0xb5, 0x8a,
	0x54, 0x5a, 0xe1, 0xa2, 0x2c, 0xba, 0x82, 0x4d, 0x93, 0x02, 0x85, 0x46, 0x81, 0x69, 0x8b, 0x2a,
	0x16, 0x45, 0x13, 0xe7, 0xa4, 0x36, 0xb1, 0x3d, 0x66, 0x66, 0x5c, 0xe1, 0x35, 0xe2, 0x01, 0xd8,
	0xf0, 0x32, 0xbc, 0x07, 0x2f, 0xc2, 0x0b, 0x20, 0x5f, 0x92, 0xda, 0x8e, 0xdb, 0xb4, 0x02, 0xb1,
	0xf3, 0x7c, 0x73, 0xce, 0x77, 0xbe, 0x73, 0x1b, 0xc3, 0x3d, 0xdf, 0x61, 0xe1, 0x07, 0x89, 0x4a,
	0xd9, 0xde, 0xb9, 0x34, 0x7c, 0xc1, 0x15, 0x27, 0xab, 0x11, 0xc8, 0x7c, 0x3b, 0x39, 0x1a, 0x03,
	0xdb, 0x93, 0xc3, 0xd0, 0x18, 0x19, 0xad, 0x24, 0x3e, 0x7d, 0x6e, 0x06, 0x2e, 0x7a, 0x2a, 0x31,
	0x5a, 0xa9, 0xa5, 0x44, 0x4c, 0x98, 0x56, 0x02, 0xe9, 0xdf, 0x2a, 0xb0, 0x7a, 0x22, 0x51, 0x1c,
	0xa5, 0x8e, 0x2d, 0xee, 0x49, 0x5b, 0x2a, 0xf4, 0xcc, 0xf0, 0x98, 0x0f, 0xd1, 0x93, 0xe4, 0xab,
	0x06, 0x75, 0xb3, 0x80, 0x1e, 0x78, 0x03, 0xde, 0xd0, 0xd6, 0x66, 0x36, 0xab, 0xcd, 0xb7, 0xc6,
	0xb5, 0x42, 0x8c, 0x6b, 0xc9, 0x8d, 0x56, 0x09, 0x31, 0x2d, 0x0d, 0xb7, 0x62, 0x41, 0xbd, 0xcc,
	0x9a, 0x6c, 0xc0, 0x1d, 0x81, 0x9f, 0x02, 0x94, 0xea, 0x25, 0xb2, 0x3e, 0x8a, 0x86, 0xb6, 0xa6,
	0x6d, 0xfe, 0x4f, 0xf3, 0x20, 0xd9, 0x82, 0xbb, 0x45, 0xd6, 0x46, 0x25, 0x36, 0x9c, 0xc0, 0xf5,
	0xd7, 0x50, 0xeb, 0x30, 0x31, 0xc4, 0x48, 0xf3, 0x48, 0x3a, 0xd9, 0x85, 0x65, 0x77, 0x04, 0xee,
	0xbb, 0xcc, 0x76, 0x64, 0xd7, 0x57, 0xd8, 0x3f, 0xf0, 0xe2, 0x78, 0xff, 0xd1, 0x2b, 0x6e, 0xf5,
	0x0b, 0x58, 0x7c, 0x23, 0xec, 0x0b, 0x66, 0x86, 0x29, 0x15, 0x21, 0x30, 0xab, 0x42, 0x1f, 0x63,
	0xbf, 0x39, 0x1a, 0x7f, 0x47, 0x49, 0x98, 0x81, 0x10, 0xe8, 0xa9, 0x23, 0xc5, 0x54, 0x20, 0x63,
	0x6d, 0x73, 0x34, 0x0f, 0x46, 0x49, 0xa0, 0xc7, 0x7a, 0x0e, 0xf6, 0xf7, 0xc2, 0x36, 0x0e, 0x58,
	0xe0, 0xa8, 0xc6, 0x4c, 0x1c, 0x7d, 0x02, 0xd7, 0x2d, 0x58, 0xca, 0xc7, 0x95, 0xe4, 0x04, 0x16,
	0xfd, 0x1c, 0x94, 0xb6, 0xf0, 0xf1, 0x94, 0x16, 0xe6, 0x79, 0x68, 0x81, 0x44, 0xff, 0xa2, 0x01,
	0x3c, 0x67, 0xae, 0xed, 0x84, 0x71, 0x3f, 0x76, 0x61, 0x79, 0x10, 0x9f, 0x3a, 0xe8, 0xf6, 0x50,
	0x48, 0xcb, 0xf6, 0xd3, 0x9c, 0x92, 0x84, 0xaf, 0xb8, 0x25, 0xcf, 0x60, 0x3e, 0xb9, 0x89, 0x73,
	0xaf, 0x36, 0x1f, 0x4e, 0x51, 0x95, 0x84, 0xa4, 0xa9, 0x93, 0xde, 0x81, 0xf9, 0x04, 0x21, 0x2d,
	0x98, 0x77, 0x63, 0xf2, 0x34, 0xbd, 0xed, 0x1b, 0x11, 0x25, 0x7a, 0x68, 0xea, 0xaa, 0xfb, 0xb0,
	0x90, 0xc5, 0xa3, 0xa6, 0x09, 0xee, 0x8c, 0x9b, 0x16, 0x7d, 0x93, 0x43, 0x58, 0xf4, 0x51, 0x48,
	0xee, 0xb5, 0xd3, 0x35, 0x4b, 0x95, 0x6f, 0x94, 0x07, 0x1c, 0x2f, 0x63, 0x9b, 0x9b, 0xef, 0x9a,
	0xb4, 0xe0, 0xab, 0xaf, 0x01, 0x74, 0xbd, 0x1e, 0x67, 0xa2, 0x7f, 0xc5, 0x90, 0xe8, 0xa7, 0x50,
	0xbd, 0xb4, 0x90, 0xe4, 0x00, 0x80, 0x8f, 0x8f, 0x69, 0xae, 0x8f, 0xa6, 0xe4, 0x7a, 0xe9, 0x4f,
	0x33, 0xce, 0xfa, 0xaf, 0x0a, 0x2c, 0x64, 0x17, 0x95, 0x9c, 0x41, 0xcd, 0x2d, 0xae, 0x40, 0xac,
	0xa5, 0xda, 0x7c, 0x32, 0x25, 0xc4, 0xc4, 0xea, 0xd0, 0x49, 0x2a, 0x72, 0x0a, 0x4b, 0xf9, 0x29,
	0x92, 0x69, 0xed, 0x8c, 0x5b, 0xcd, 0xa2, 0xa4, 0x45, 0x9a, 0xa8, 0x2a, 0x83, 0xf1, 0x30, 0xc6,
	0xdb, 0x31, 0xbd, 0x2a, 0x97, 0xd3, 0x4b, 0x33, 0xce, 0xe4, 0x0c, 0xea, 0x7d, 0x5b, 0xba, 0xb6,
	0x94, 0xd8, 0xcf, 0x14, 0xbe, 0x31, 0x1b, 0x93, 0x6e, 0xdd, 0xb8, 0xd4, 0x92, 0x96, 0xf2, 0xe8,
	0x3f, 0x34, 0xa8, 0x77, 0xf7, 0x8e, 0xba, 0x87, 0xfb, 0xc7, 0xfb, 0xb9, 0xea, 0x1f, 0xc2, 0xba,
	0xe2, 0xb2, 0x65, 0xa1, 0x39, 0xec, 0xf1, 0xcf, 0x9d, 0xeb, 0x9e, 0x9d, 0xe9, 0x86, 0x25, 0x6b,
	0x5f, 0xf9, 0x1b, 0x6b, 0xff, 0x5d, 0x83, 0x7a, 0x46, 0x75, 0xdb, 0x16, 0x2a, 0x6c, 0x33, 0xc5,
	0x4a, 0xdf, 0xb7, 0x8f, 0x50, 0x2b, 0x3e, 0xb3, 0xa3, 0x8e, 0x3f, 0xfd, 0x93, 0x1f, 0x08, 0x9d,
	0xa4, 0xd5, 0x7f, 0x6a, 0x70, 0xff, 0x05, 0xaa, 0xac, 0x1f, 0x45, 0xe9, 0x73, 0x4f, 0x22, 0xe9,
	0xc2, 0x42, 0x90, 0xc1, 0xd3, 0x91, 0xde, 0xbe, 0x85, 0x04, 0x9a, 0x23, 0xf8, 0x97, 0x89, 0xed,
	0xad, 0xbf, 0x7f, 0x70, 0x6e, 0x2b, 0x2b, 0xe8, 0x19, 0x26, 0x77, 0x77, 0x5e, 0x71, 0x69, 0x05,
	0xac, 0xcd, 0x51, 0xee, 0x38, 0x76, 0x4f, 0x60, 0x14, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0x01, 0x29, 0x8f, 0x2b, 0x08, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: play_filter_rules.proto

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

type Availability struct {
	Restriction                      *int32                                           `protobuf:"varint,5,opt,name=restriction" json:"restriction,omitempty"`
	OfferType                        *int32                                           `protobuf:"varint,6,opt,name=offerType" json:"offerType,omitempty"`
	Rule                             *Rule                                            `protobuf:"bytes,7,opt,name=rule" json:"rule,omitempty"`
	Perdeviceavailabilityrestriction []*Availability_PerDeviceAvailabilityRestriction `protobuf:"group,9,rep,name=PerDeviceAvailabilityRestriction,json=perdeviceavailabilityrestriction" json:"perdeviceavailabilityrestriction,omitempty"`
	AvailableIfOwned                 *bool                                            `protobuf:"varint,13,opt,name=availableIfOwned" json:"availableIfOwned,omitempty"`
	Install                          []*Install                                       `protobuf:"bytes,14,rep,name=install" json:"install,omitempty"`
	FilterInfo                       *FilterEvaluationInfo                            `protobuf:"bytes,16,opt,name=filterInfo" json:"filterInfo,omitempty"`
	OwnershipInfo                    *OwnershipInfo                                   `protobuf:"bytes,17,opt,name=ownershipInfo" json:"ownershipInfo,omitempty"`
	AvailabilityProblem              []*AvailabilityProblem                           `protobuf:"bytes,18,rep,name=availabilityProblem" json:"availabilityProblem,omitempty"`
	Hidden                           *bool                                            `protobuf:"varint,21,opt,name=hidden" json:"hidden,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                                         `json:"-"`
	XXX_unrecognized                 []byte                                           `json:"-"`
	XXX_sizecache                    int32                                            `json:"-"`
}

func (m *Availability) Reset()         { *m = Availability{} }
func (m *Availability) String() string { return proto.CompactTextString(m) }
func (*Availability) ProtoMessage()    {}
func (*Availability) Descriptor() ([]byte, []int) {
	return fileDescriptor_866bbbb49451b6e5, []int{0}
}

func (m *Availability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Availability.Unmarshal(m, b)
}
func (m *Availability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Availability.Marshal(b, m, deterministic)
}
func (m *Availability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Availability.Merge(m, src)
}
func (m *Availability) XXX_Size() int {
	return xxx_messageInfo_Availability.Size(m)
}
func (m *Availability) XXX_DiscardUnknown() {
	xxx_messageInfo_Availability.DiscardUnknown(m)
}

var xxx_messageInfo_Availability proto.InternalMessageInfo

func (m *Availability) GetRestriction() int32 {
	if m != nil && m.Restriction != nil {
		return *m.Restriction
	}
	return 0
}

func (m *Availability) GetOfferType() int32 {
	if m != nil && m.OfferType != nil {
		return *m.OfferType
	}
	return 0
}

func (m *Availability) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Availability) GetPerdeviceavailabilityrestriction() []*Availability_PerDeviceAvailabilityRestriction {
	if m != nil {
		return m.Perdeviceavailabilityrestriction
	}
	return nil
}

func (m *Availability) GetAvailableIfOwned() bool {
	if m != nil && m.AvailableIfOwned != nil {
		return *m.AvailableIfOwned
	}
	return false
}

func (m *Availability) GetInstall() []*Install {
	if m != nil {
		return m.Install
	}
	return nil
}

func (m *Availability) GetFilterInfo() *FilterEvaluationInfo {
	if m != nil {
		return m.FilterInfo
	}
	return nil
}

func (m *Availability) GetOwnershipInfo() *OwnershipInfo {
	if m != nil {
		return m.OwnershipInfo
	}
	return nil
}

func (m *Availability) GetAvailabilityProblem() []*AvailabilityProblem {
	if m != nil {
		return m.AvailabilityProblem
	}
	return nil
}

func (m *Availability) GetHidden() bool {
	if m != nil && m.Hidden != nil {
		return *m.Hidden
	}
	return false
}

type Availability_PerDeviceAvailabilityRestriction struct {
	AndroidId            *uint64               `protobuf:"fixed64,10,opt,name=androidId" json:"androidId,omitempty"`
	DeviceRestriction    *int32                `protobuf:"varint,11,opt,name=deviceRestriction" json:"deviceRestriction,omitempty"`
	ChannelId            *int64                `protobuf:"varint,12,opt,name=channelId" json:"channelId,omitempty"`
	FilterInfo           *FilterEvaluationInfo `protobuf:"bytes,15,opt,name=filterInfo" json:"filterInfo,omitempty"`
	AvailableIfOwned     *bool                 `protobuf:"varint,22,opt,name=availableIfOwned" json:"availableIfOwned,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Availability_PerDeviceAvailabilityRestriction) Reset() {
	*m = Availability_PerDeviceAvailabilityRestriction{}
}
func (m *Availability_PerDeviceAvailabilityRestriction) String() string {
	return proto.CompactTextString(m)
}
func (*Availability_PerDeviceAvailabilityRestriction) ProtoMessage() {}
func (*Availability_PerDeviceAvailabilityRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_866bbbb49451b6e5, []int{0, 0}
}

func (m *Availability_PerDeviceAvailabilityRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Availability_PerDeviceAvailabilityRestriction.Unmarshal(m, b)
}
func (m *Availability_PerDeviceAvailabilityRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Availability_PerDeviceAvailabilityRestriction.Marshal(b, m, deterministic)
}
func (m *Availability_PerDeviceAvailabilityRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Availability_PerDeviceAvailabilityRestriction.Merge(m, src)
}
func (m *Availability_PerDeviceAvailabilityRestriction) XXX_Size() int {
	return xxx_messageInfo_Availability_PerDeviceAvailabilityRestriction.Size(m)
}
func (m *Availability_PerDeviceAvailabilityRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_Availability_PerDeviceAvailabilityRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_Availability_PerDeviceAvailabilityRestriction proto.InternalMessageInfo

func (m *Availability_PerDeviceAvailabilityRestriction) GetAndroidId() uint64 {
	if m != nil && m.AndroidId != nil {
		return *m.AndroidId
	}
	return 0
}

func (m *Availability_PerDeviceAvailabilityRestriction) GetDeviceRestriction() int32 {
	if m != nil && m.DeviceRestriction != nil {
		return *m.DeviceRestriction
	}
	return 0
}

func (m *Availability_PerDeviceAvailabilityRestriction) GetChannelId() int64 {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return 0
}

func (m *Availability_PerDeviceAvailabilityRestriction) GetFilterInfo() *FilterEvaluationInfo {
	if m != nil {
		return m.FilterInfo
	}
	return nil
}

func (m *Availability_PerDeviceAvailabilityRestriction) GetAvailableIfOwned() bool {
	if m != nil && m.AvailableIfOwned != nil {
		return *m.AvailableIfOwned
	}
	return false
}

type AvailabilityProblem struct {
	ProblemType          *int32   `protobuf:"varint,1,opt,name=problemType" json:"problemType,omitempty"`
	MissingValue         []string `protobuf:"bytes,2,rep,name=missingValue" json:"missingValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvailabilityProblem) Reset()         { *m = AvailabilityProblem{} }
func (m *AvailabilityProblem) String() string { return proto.CompactTextString(m) }
func (*AvailabilityProblem) ProtoMessage()    {}
func (*AvailabilityProblem) Descriptor() ([]byte, []int) {
	return fileDescriptor_866bbbb49451b6e5, []int{1}
}

func (m *AvailabilityProblem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailabilityProblem.Unmarshal(m, b)
}
func (m *AvailabilityProblem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailabilityProblem.Marshal(b, m, deterministic)
}
func (m *AvailabilityProblem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailabilityProblem.Merge(m, src)
}
func (m *AvailabilityProblem) XXX_Size() int {
	return xxx_messageInfo_AvailabilityProblem.Size(m)
}
func (m *AvailabilityProblem) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailabilityProblem.DiscardUnknown(m)
}

var xxx_messageInfo_AvailabilityProblem proto.InternalMessageInfo

func (m *AvailabilityProblem) GetProblemType() int32 {
	if m != nil && m.ProblemType != nil {
		return *m.ProblemType
	}
	return 0
}

func (m *AvailabilityProblem) GetMissingValue() []string {
	if m != nil {
		return m.MissingValue
	}
	return nil
}

type FilterEvaluationInfo struct {
	RuleEvaluation       []*RuleEvaluation `protobuf:"bytes,1,rep,name=ruleEvaluation" json:"ruleEvaluation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FilterEvaluationInfo) Reset()         { *m = FilterEvaluationInfo{} }
func (m *FilterEvaluationInfo) String() string { return proto.CompactTextString(m) }
func (*FilterEvaluationInfo) ProtoMessage()    {}
func (*FilterEvaluationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_866bbbb49451b6e5, []int{2}
}

func (m *FilterEvaluationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterEvaluationInfo.Unmarshal(m, b)
}
func (m *FilterEvaluationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterEvaluationInfo.Marshal(b, m, deterministic)
}
func (m *FilterEvaluationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterEvaluationInfo.Merge(m, src)
}
func (m *FilterEvaluationInfo) XXX_Size() int {
	return xxx_messageInfo_FilterEvaluationInfo.Size(m)
}
func (m *FilterEvaluationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterEvaluationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FilterEvaluationInfo proto.InternalMessageInfo

func (m *FilterEvaluationInfo) GetRuleEvaluation() []*RuleEvaluation {
	if m != nil {
		return m.RuleEvaluation
	}
	return nil
}

type Rule struct {
	Negate                  *bool     `protobuf:"varint,1,opt,name=negate" json:"negate,omitempty"`
	Operator                *int32    `protobuf:"varint,2,opt,name=operator" json:"operator,omitempty"`
	Key                     *int32    `protobuf:"varint,3,opt,name=key" json:"key,omitempty"`
	StringArg               []string  `protobuf:"bytes,4,rep,name=stringArg" json:"stringArg,omitempty"`
	LongArg                 []int64   `protobuf:"varint,5,rep,name=longArg" json:"longArg,omitempty"`
	DoubleArg               []float64 `protobuf:"fixed64,6,rep,name=doubleArg" json:"doubleArg,omitempty"`
	Subrule                 []*Rule   `protobuf:"bytes,7,rep,name=subrule" json:"subrule,omitempty"`
	ResponseCode            *int32    `protobuf:"varint,8,opt,name=responseCode" json:"responseCode,omitempty"`
	Comment                 *string   `protobuf:"bytes,9,opt,name=comment" json:"comment,omitempty"`
	StringArgHash           []uint64  `protobuf:"fixed64,10,rep,name=stringArgHash" json:"stringArgHash,omitempty"`
	AvailabilityProblemType *int32    `protobuf:"varint,12,opt,name=availabilityProblemType" json:"availabilityProblemType,omitempty"`
	IncludeMissingValues    *bool     `protobuf:"varint,13,opt,name=includeMissingValues" json:"includeMissingValues,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}  `json:"-"`
	XXX_unrecognized        []byte    `json:"-"`
	XXX_sizecache           int32     `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_866bbbb49451b6e5, []int{3}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetNegate() bool {
	if m != nil && m.Negate != nil {
		return *m.Negate
	}
	return false
}

func (m *Rule) GetOperator() int32 {
	if m != nil && m.Operator != nil {
		return *m.Operator
	}
	return 0
}

func (m *Rule) GetKey() int32 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

func (m *Rule) GetStringArg() []string {
	if m != nil {
		return m.StringArg
	}
	return nil
}

func (m *Rule) GetLongArg() []int64 {
	if m != nil {
		return m.LongArg
	}
	return nil
}

func (m *Rule) GetDoubleArg() []float64 {
	if m != nil {
		return m.DoubleArg
	}
	return nil
}

func (m *Rule) GetSubrule() []*Rule {
	if m != nil {
		return m.Subrule
	}
	return nil
}

func (m *Rule) GetResponseCode() int32 {
	if m != nil && m.ResponseCode != nil {
		return *m.ResponseCode
	}
	return 0
}

func (m *Rule) GetComment() string {
	if m != nil && m.Comment != nil {
		return *m.Comment
	}
	return ""
}

func (m *Rule) GetStringArgHash() []uint64 {
	if m != nil {
		return m.StringArgHash
	}
	return nil
}

func (m *Rule) GetAvailabilityProblemType() int32 {
	if m != nil && m.AvailabilityProblemType != nil {
		return *m.AvailabilityProblemType
	}
	return 0
}

func (m *Rule) GetIncludeMissingValues() bool {
	if m != nil && m.IncludeMissingValues != nil {
		return *m.IncludeMissingValues
	}
	return false
}

type RuleEvaluation struct {
	Rule                 *Rule     `protobuf:"bytes,1,opt,name=rule" json:"rule,omitempty"`
	ActualStringValue    []string  `protobuf:"bytes,2,rep,name=actualStringValue" json:"actualStringValue,omitempty"`
	ActualLongValue      []int64   `protobuf:"varint,3,rep,name=actualLongValue" json:"actualLongValue,omitempty"`
	ActualBoolValue      []bool    `protobuf:"varint,4,rep,name=actualBoolValue" json:"actualBoolValue,omitempty"`
	ActualDoubleValue    []float64 `protobuf:"fixed64,5,rep,name=actualDoubleValue" json:"actualDoubleValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RuleEvaluation) Reset()         { *m = RuleEvaluation{} }
func (m *RuleEvaluation) String() string { return proto.CompactTextString(m) }
func (*RuleEvaluation) ProtoMessage()    {}
func (*RuleEvaluation) Descriptor() ([]byte, []int) {
	return fileDescriptor_866bbbb49451b6e5, []int{4}
}

func (m *RuleEvaluation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleEvaluation.Unmarshal(m, b)
}
func (m *RuleEvaluation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleEvaluation.Marshal(b, m, deterministic)
}
func (m *RuleEvaluation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleEvaluation.Merge(m, src)
}
func (m *RuleEvaluation) XXX_Size() int {
	return xxx_messageInfo_RuleEvaluation.Size(m)
}
func (m *RuleEvaluation) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleEvaluation.DiscardUnknown(m)
}

var xxx_messageInfo_RuleEvaluation proto.InternalMessageInfo

func (m *RuleEvaluation) GetRule() *Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *RuleEvaluation) GetActualStringValue() []string {
	if m != nil {
		return m.ActualStringValue
	}
	return nil
}

func (m *RuleEvaluation) GetActualLongValue() []int64 {
	if m != nil {
		return m.ActualLongValue
	}
	return nil
}

func (m *RuleEvaluation) GetActualBoolValue() []bool {
	if m != nil {
		return m.ActualBoolValue
	}
	return nil
}

func (m *RuleEvaluation) GetActualDoubleValue() []float64 {
	if m != nil {
		return m.ActualDoubleValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Availability)(nil), "playapi.proto.finsky.filter_rules.Availability")
	proto.RegisterType((*Availability_PerDeviceAvailabilityRestriction)(nil), "playapi.proto.finsky.filter_rules.Availability.PerDeviceAvailabilityRestriction")
	proto.RegisterType((*AvailabilityProblem)(nil), "playapi.proto.finsky.filter_rules.AvailabilityProblem")
	proto.RegisterType((*FilterEvaluationInfo)(nil), "playapi.proto.finsky.filter_rules.FilterEvaluationInfo")
	proto.RegisterType((*Rule)(nil), "playapi.proto.finsky.filter_rules.Rule")
	proto.RegisterType((*RuleEvaluation)(nil), "playapi.proto.finsky.filter_rules.RuleEvaluation")
}

func init() { proto.RegisterFile("play_filter_rules.proto", fileDescriptor_866bbbb49451b6e5) }

var fileDescriptor_866bbbb49451b6e5 = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x86, 0xa2, 0xf8, 0xef, 0xd8, 0xf9, 0x31, 0x93, 0x25, 0x84, 0x31, 0x60, 0x8a, 0x37, 0x6c,
	0xc2, 0x10, 0x38, 0x98, 0x2f, 0xb6, 0x01, 0xbb, 0x59, 0xb2, 0xac, 0xa8, 0x8b, 0x16, 0x09, 0xd8,
	0xa2, 0x45, 0xdb, 0x8b, 0x80, 0xb6, 0x68, 0x9b, 0x08, 0x4d, 0xaa, 0x94, 0x94, 0xc0, 0x8f, 0xd1,
	0xfb, 0xbe, 0x47, 0xdf, 0xa3, 0xef, 0x53, 0xa0, 0x20, 0x65, 0x5b, 0xb2, 0xad, 0x22, 0x49, 0xd1,
	0x3b, 0x9f, 0xef, 0xfc, 0xf8, 0x3b, 0xfc, 0x3e, 0x52, 0x70, 0x18, 0x0a, 0x3a, 0xbd, 0x1a, 0x72,
	0x11, 0x33, 0x7d, 0xa5, 0x13, 0xc1, 0xa2, 0x4e, 0xa8, 0x55, 0xac, 0xd0, 0x91, 0x49, 0xd0, 0x90,
	0xa7, 0x61, 0x67, 0xc8, 0x65, 0x74, 0x3d, 0xed, 0xe4, 0x0b, 0x5b, 0x4d, 0xdb, 0x3b, 0x50, 0x93,
	0x89, 0x92, 0x69, 0x59, 0x6b, 0xcf, 0x42, 0x81, 0xba, 0x95, 0x42, 0xd1, 0x20, 0x05, 0xdb, 0x9f,
	0x2a, 0xd0, 0x38, 0xbd, 0xa1, 0x5c, 0xd0, 0x3e, 0x17, 0x3c, 0x9e, 0x22, 0x0f, 0xea, 0x9a, 0x45,
	0xb1, 0xe6, 0x83, 0x98, 0x2b, 0x89, 0x4b, 0x9e, 0xe3, 0x97, 0x48, 0x1e, 0x42, 0x3f, 0x42, 0x4d,
	0x0d, 0x87, 0x4c, 0xbf, 0x98, 0x86, 0x0c, 0x97, 0x6d, 0x3e, 0x03, 0xd0, 0x3f, 0xb0, 0x69, 0x18,
	0xe0, 0x8a, 0xe7, 0xf8, 0xf5, 0xee, 0x6f, 0x9d, 0x3b, 0xa9, 0x76, 0x48, 0x22, 0x18, 0xb1, 0x4d,
	0xe8, 0x83, 0x03, 0x5e, 0xc8, 0x74, 0xc0, 0x6e, 0xf8, 0x80, 0xd1, 0x1c, 0xad, 0x3c, 0xa5, 0x9a,
	0xe7, 0xfa, 0xd0, 0xbd, 0xbc, 0xc7, 0xe4, 0xfc, 0x62, 0x9d, 0x4b, 0xa6, 0xcf, 0xed, 0xdc, 0x3c,
	0x4a, 0xb2, 0xb9, 0xe4, 0xce, 0x7f, 0x46, 0xbf, 0xc3, 0xee, 0x2c, 0x25, 0x58, 0x6f, 0x78, 0x71,
	0x2b, 0x59, 0x80, 0xb7, 0x3c, 0xc7, 0xaf, 0x92, 0x35, 0x1c, 0xfd, 0x0b, 0x15, 0x2e, 0xa3, 0x98,
	0x0a, 0x81, 0xb7, 0x3d, 0xd7, 0xaf, 0x77, 0x7f, 0x2d, 0x26, 0xbc, 0xd0, 0xa3, 0x97, 0x56, 0x93,
	0x79, 0x1b, 0x7a, 0x05, 0x90, 0x6e, 0xd3, 0x93, 0x43, 0x85, 0x77, 0xed, 0x79, 0xfe, 0x75, 0x8f,
	0xad, 0x1f, 0xd9, 0xe0, 0xff, 0x1b, 0x2a, 0x12, 0x6a, 0x68, 0x9b, 0x76, 0x92, 0x1b, 0x85, 0x7a,
	0xb0, 0xa5, 0x6e, 0x25, 0xd3, 0xd1, 0x98, 0x87, 0x76, 0x76, 0xd3, 0xce, 0xfe, 0xb9, 0x78, 0xf6,
	0x45, 0xbe, 0x94, 0x2c, 0x77, 0xa2, 0x31, 0xec, 0xe5, 0x0f, 0xeb, 0x52, 0xab, 0xbe, 0x60, 0x13,
	0x8c, 0xec, 0xc6, 0x7f, 0x3e, 0x50, 0xa2, 0x59, 0x37, 0x29, 0x1a, 0x89, 0x0e, 0xa0, 0x3c, 0xe6,
	0x41, 0xc0, 0x24, 0xfe, 0xc1, 0x9e, 0xf8, 0x2c, 0x6a, 0xbd, 0xdf, 0x00, 0xef, 0x2e, 0x69, 0x8d,
	0x65, 0xa9, 0x0c, 0xb4, 0xe2, 0x41, 0x2f, 0xc0, 0xe0, 0x39, 0x7e, 0x99, 0x64, 0x00, 0x3a, 0x86,
	0x66, 0xaa, 0x7b, 0xae, 0x05, 0xd7, 0xad, 0xb1, 0xd7, 0x13, 0x66, 0xd6, 0x60, 0x4c, 0xa5, 0x64,
	0xa2, 0x17, 0xe0, 0x86, 0xe7, 0xf8, 0x2e, 0xc9, 0x80, 0x15, 0xd1, 0x76, 0xbe, 0x9f, 0x68, 0x45,
	0xde, 0x3b, 0x28, 0xf6, 0x5e, 0xfb, 0x2d, 0xec, 0x15, 0x9c, 0xab, 0xb9, 0xda, 0x61, 0xfa, 0xd3,
	0x5e, 0x5d, 0x27, 0xbd, 0xda, 0x39, 0x08, 0xb5, 0xa1, 0x31, 0xe1, 0x51, 0xc4, 0xe5, 0xe8, 0x25,
	0x15, 0x09, 0xc3, 0x1b, 0x9e, 0xeb, 0xd7, 0xc8, 0x12, 0xd6, 0x7e, 0x07, 0xfb, 0x45, 0x64, 0xd1,
	0x6b, 0xd8, 0x36, 0xab, 0x64, 0x28, 0x76, 0xac, 0x0b, 0xfe, 0xb8, 0xe7, 0x13, 0x90, 0x35, 0x92,
	0x95, 0x41, 0xed, 0x8f, 0x2e, 0x6c, 0x9a, 0x12, 0x63, 0x02, 0xc9, 0x46, 0x34, 0x4e, 0xc9, 0x57,
	0xc9, 0x2c, 0x42, 0x2d, 0xa8, 0xaa, 0x90, 0x69, 0x1a, 0x2b, 0x8d, 0x37, 0xec, 0x5a, 0x8b, 0x18,
	0xed, 0x82, 0x7b, 0xcd, 0xa6, 0xd8, 0xb5, 0xb0, 0xf9, 0x69, 0x14, 0x34, 0x72, 0xca, 0xd1, 0xa9,
	0x1e, 0xe1, 0x4d, 0xbb, 0x62, 0x06, 0x20, 0x0c, 0x15, 0xa1, 0xd2, 0x5c, 0xc9, 0x73, 0x7d, 0x97,
	0xcc, 0x43, 0xd3, 0x17, 0xa8, 0xa4, 0x2f, 0x98, 0xc9, 0x95, 0x3d, 0xd7, 0x77, 0x48, 0x06, 0xa0,
	0x53, 0xa8, 0x44, 0x49, 0x7f, 0xf6, 0xf6, 0xb9, 0x0f, 0x79, 0xfb, 0xe6, 0x7d, 0xe6, 0xf8, 0x35,
	0x8b, 0x42, 0x25, 0x23, 0xf6, 0x9f, 0x0a, 0x18, 0xae, 0x5a, 0xce, 0x4b, 0x98, 0xa1, 0x67, 0x5e,
	0x75, 0x26, 0x63, 0x5c, 0xf3, 0x1c, 0xbf, 0x46, 0xe6, 0x21, 0xfa, 0x05, 0xb6, 0x16, 0x5b, 0x3c,
	0xa6, 0xd1, 0x18, 0x83, 0xe7, 0xfa, 0x65, 0xb2, 0x0c, 0xa2, 0xbf, 0xe1, 0xb0, 0xe0, 0x7a, 0x59,
	0x43, 0x34, 0xec, 0xdf, 0x7d, 0x2d, 0x8d, 0xba, 0xb0, 0xcf, 0xe5, 0x40, 0x24, 0x01, 0x7b, 0x96,
	0xf3, 0x43, 0x34, 0x7b, 0x01, 0x0b, 0x73, 0xed, 0xcf, 0x0e, 0x6c, 0x2f, 0x8b, 0xbb, 0xf8, 0x40,
	0x38, 0xdf, 0xf2, 0x81, 0x38, 0x86, 0x26, 0x1d, 0xc4, 0x09, 0x15, 0xcf, 0xed, 0x52, 0x79, 0x97,
	0xae, 0x27, 0x90, 0x0f, 0x3b, 0x29, 0xf8, 0x54, 0xcd, 0x6b, 0x5d, 0x2b, 0xe9, 0x2a, 0x9c, 0x55,
	0x9e, 0x29, 0x25, 0xd2, 0x4a, 0x63, 0x8c, 0x2a, 0x59, 0x85, 0x33, 0x06, 0xe7, 0x56, 0xf9, 0xb4,
	0xb6, 0x64, 0xcd, 0xb0, 0x9e, 0x38, 0x3b, 0x7a, 0xf3, 0xd3, 0x88, 0xc7, 0xe3, 0xa4, 0xdf, 0x19,
	0xa8, 0xc9, 0xc9, 0x13, 0x15, 0x8d, 0x13, 0x7a, 0xae, 0x58, 0x74, 0x22, 0x78, 0x5f, 0x33, 0xb3,
	0xfa, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0x6a, 0xea, 0x3b, 0xe6, 0x07, 0x00, 0x00,
}

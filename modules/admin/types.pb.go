// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/admin/types.proto

package admin

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	gitlab_bianjie_ai_cschain_cschain_sdk_go_types "gitlab.bianjie.ai/cschain/cschain-sdk-go/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Role is a type alias that represents a proposal status as a byte
type Role int32

const (
	// ROLE_ROOT_ADMIN defines the root admin role index.
	RoleRootAdmin Role = 0
	// ROLE_PERM_ADMIN defines the permission admin role index.
	RolePermAdmin Role = 1
	// ROLE_BLACKLIST_ADMIN defines the blacklist admin role index.
	RoleBlacklistAdmin Role = 2
	// ROLE_NODE_ADMIN defines the validator node admin role index.
	RoleNodeAdmin Role = 3
	// ROLE_PARAM_ADMIN defines the param admin role index.
	RoleParamAdmin Role = 4
	// ROLE_POWER_USER defines the power user role index.
	RolePowerUser Role = 5
)

var Role_name = map[int32]string{
	0: "ROLE_ROOT_ADMIN",
	1: "ROLE_PERM_ADMIN",
	2: "ROLE_BLACKLIST_ADMIN",
	3: "ROLE_NODE_ADMIN",
	4: "ROLE_PARAM_ADMIN",
	5: "ROLE_POWER_USER",
}

var Role_value = map[string]int32{
	"ROLE_ROOT_ADMIN":      0,
	"ROLE_PERM_ADMIN":      1,
	"ROLE_BLACKLIST_ADMIN": 2,
	"ROLE_NODE_ADMIN":      3,
	"ROLE_PARAM_ADMIN":     4,
	"ROLE_POWER_USER":      5,
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cfc93ac7f115920f, []int{0}
}

// MsgAddRoles defines an SDK message for adding roles to a address.
type MsgAddRoles struct {
	Address  gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"address,omitempty"`
	Roles    []Role                                                    `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=cschain.modules.admin.Role" json:"roles,omitempty"`
	Operator gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,3,opt,name=operator,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"operator,omitempty"`
}

func (m *MsgAddRoles) Reset()         { *m = MsgAddRoles{} }
func (m *MsgAddRoles) String() string { return proto.CompactTextString(m) }
func (*MsgAddRoles) ProtoMessage()    {}
func (*MsgAddRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc93ac7f115920f, []int{0}
}
func (m *MsgAddRoles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddRoles.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddRoles.Merge(m, src)
}
func (m *MsgAddRoles) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddRoles.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddRoles proto.InternalMessageInfo

// MsgRemoveRoles defines an SDK message for removing roles from an existing address.
type MsgRemoveRoles struct {
	Address  gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"address,omitempty"`
	Roles    []Role                                                    `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=cschain.modules.admin.Role" json:"roles,omitempty"`
	Operator gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,3,opt,name=operator,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"operator,omitempty"`
}

func (m *MsgRemoveRoles) Reset()         { *m = MsgRemoveRoles{} }
func (m *MsgRemoveRoles) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveRoles) ProtoMessage()    {}
func (*MsgRemoveRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc93ac7f115920f, []int{1}
}
func (m *MsgRemoveRoles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveRoles.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveRoles.Merge(m, src)
}
func (m *MsgRemoveRoles) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveRoles.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveRoles proto.InternalMessageInfo

// MsgBlockAccount defines an SDK message for blocking an account.
type MsgBlockAccount struct {
	Address  gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"address,omitempty"`
	Operator gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,2,opt,name=operator,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"operator,omitempty"`
}

func (m *MsgBlockAccount) Reset()         { *m = MsgBlockAccount{} }
func (m *MsgBlockAccount) String() string { return proto.CompactTextString(m) }
func (*MsgBlockAccount) ProtoMessage()    {}
func (*MsgBlockAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc93ac7f115920f, []int{2}
}
func (m *MsgBlockAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBlockAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBlockAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBlockAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBlockAccount.Merge(m, src)
}
func (m *MsgBlockAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgBlockAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBlockAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBlockAccount proto.InternalMessageInfo

// MsgUnblockAccount defines an SDK message for unblocking an account.
type MsgUnblockAccount struct {
	Address  gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"address,omitempty"`
	Operator gitlab_bianjie_ai_cschain_cschain_sdk_go_types.AccAddress `protobuf:"bytes,2,opt,name=operator,proto3,casttype=gitlab.bianjie.ai/cschain/cschain-sdk-go/types.AccAddress" json:"operator,omitempty"`
}

func (m *MsgUnblockAccount) Reset()         { *m = MsgUnblockAccount{} }
func (m *MsgUnblockAccount) String() string { return proto.CompactTextString(m) }
func (*MsgUnblockAccount) ProtoMessage()    {}
func (*MsgUnblockAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc93ac7f115920f, []int{3}
}
func (m *MsgUnblockAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnblockAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnblockAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnblockAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnblockAccount.Merge(m, src)
}
func (m *MsgUnblockAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnblockAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnblockAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnblockAccount proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cschain.modules.admin.Role", Role_name, Role_value)
	proto.RegisterType((*MsgAddRoles)(nil), "cschain.modules.admin.MsgAddRoles")
	proto.RegisterType((*MsgRemoveRoles)(nil), "cschain.modules.admin.MsgRemoveRoles")
	proto.RegisterType((*MsgBlockAccount)(nil), "cschain.modules.admin.MsgBlockAccount")
	proto.RegisterType((*MsgUnblockAccount)(nil), "cschain.modules.admin.MsgUnblockAccount")
}

func init() { proto.RegisterFile("modules/admin/types.proto", fileDescriptor_cfc93ac7f115920f) }

var fileDescriptor_cfc93ac7f115920f = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0x2e, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83,
	0x2a, 0xd1, 0x03, 0x2b, 0x91, 0x52, 0x2b, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a,
	0xa9, 0xd4, 0x07, 0xab, 0xd4, 0x4f, 0xcf, 0x4f, 0xcf, 0x47, 0xb0, 0x20, 0xda, 0x95, 0x7e, 0x31,
	0x72, 0x71, 0xfb, 0x16, 0xa7, 0x3b, 0xa6, 0xa4, 0x04, 0xe5, 0xe7, 0xa4, 0x16, 0x0b, 0x85, 0x73,
	0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x38, 0xd9,
	0xfe, 0xba, 0x27, 0x6f, 0x99, 0x9e, 0x59, 0x92, 0x93, 0x98, 0xa4, 0x97, 0x94, 0x99, 0x98, 0x97,
	0x95, 0x99, 0xaa, 0x97, 0x98, 0xa9, 0x0f, 0xb5, 0x15, 0x46, 0xeb, 0x16, 0xa7, 0x64, 0xeb, 0xa6,
	0xe7, 0x43, 0x5d, 0xe6, 0x98, 0x9c, 0xec, 0x08, 0x31, 0x24, 0x08, 0x66, 0x9a, 0x90, 0x21, 0x17,
	0x6b, 0x11, 0xc8, 0x06, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x3e, 0x23, 0x69, 0x3d, 0xac, 0xee, 0xd6,
	0x03, 0xb9, 0x22, 0x08, 0xa2, 0x52, 0x28, 0x92, 0x8b, 0x23, 0xbf, 0x20, 0xb5, 0x28, 0xb1, 0x24,
	0xbf, 0x48, 0x82, 0x99, 0x1a, 0x8e, 0x81, 0x1b, 0x67, 0xc5, 0xf2, 0x62, 0x81, 0x3c, 0xa3, 0xd2,
	0x5f, 0x46, 0x2e, 0x3e, 0xdf, 0xe2, 0xf4, 0xa0, 0xd4, 0xdc, 0xfc, 0xb2, 0xd4, 0x11, 0xe8, 0xff,
	0xd3, 0x8c, 0x5c, 0xfc, 0xbe, 0xc5, 0xe9, 0x4e, 0x39, 0xf9, 0xc9, 0xd9, 0x8e, 0xc9, 0xc9, 0xf9,
	0xa5, 0x79, 0x25, 0xb4, 0x0b, 0x00, 0x64, 0xdf, 0x30, 0xd1, 0xc2, 0x37, 0x67, 0x19, 0xb9, 0x04,
	0x7d, 0x8b, 0xd3, 0x43, 0xf3, 0x92, 0x86, 0x85, 0x7f, 0xb4, 0x26, 0x30, 0x71, 0xb1, 0x80, 0x92,
	0x83, 0x90, 0x1a, 0x17, 0x7f, 0x90, 0xbf, 0x8f, 0x6b, 0x7c, 0x90, 0xbf, 0x7f, 0x48, 0xbc, 0xa3,
	0x8b, 0xaf, 0xa7, 0x9f, 0x00, 0x83, 0x94, 0x60, 0xd7, 0x5c, 0x05, 0x5e, 0x70, 0x6a, 0xc9, 0xcf,
	0x2f, 0x71, 0x04, 0xa5, 0x1d, 0xb8, 0xba, 0x00, 0xd7, 0x20, 0x5f, 0xa8, 0x3a, 0x46, 0x84, 0xba,
	0x80, 0xd4, 0xa2, 0x5c, 0x88, 0x3a, 0x03, 0x2e, 0x11, 0xb0, 0x3a, 0x27, 0x1f, 0x47, 0x67, 0x6f,
	0x1f, 0xcf, 0x60, 0x98, 0xa1, 0x4c, 0x52, 0x62, 0x5d, 0x73, 0x15, 0x84, 0x40, 0x8a, 0x9d, 0x72,
	0x12, 0x93, 0xb3, 0x73, 0x32, 0x8b, 0xd1, 0x4c, 0xf6, 0xf3, 0x77, 0x71, 0x85, 0x2a, 0x66, 0x46,
	0x98, 0xec, 0x97, 0x9f, 0x92, 0x0a, 0x51, 0xa7, 0xc1, 0x25, 0x00, 0x71, 0x81, 0x63, 0x90, 0x23,
	0xcc, 0x09, 0x2c, 0x52, 0x42, 0x5d, 0x73, 0x15, 0xf8, 0xc0, 0x4e, 0x48, 0x2c, 0x4a, 0xcc, 0x45,
	0x73, 0xab, 0x7f, 0xb8, 0x6b, 0x50, 0x7c, 0x68, 0xb0, 0x6b, 0x90, 0x00, 0x2b, 0x92, 0x5b, 0xf3,
	0xcb, 0x53, 0x8b, 0x42, 0x8b, 0x53, 0x8b, 0xa4, 0x78, 0x3a, 0x16, 0xcb, 0x31, 0xac, 0x58, 0x22,
	0xc7, 0xb0, 0x61, 0x89, 0x1c, 0x83, 0x53, 0xc8, 0x89, 0x87, 0x72, 0x0c, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x46, 0x74, 0xd8, 0xa3, 0x14, 0xa7, 0x49, 0x6c, 0xe0,
	0xa2, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x62, 0x3c, 0xf1, 0xe7, 0x66, 0x05, 0x00, 0x00,
}

func (this *MsgAddRoles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgAddRoles)
	if !ok {
		that2, ok := that.(MsgAddRoles)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if this.Roles[i] != that1.Roles[i] {
			return false
		}
	}
	if !bytes.Equal(this.Operator, that1.Operator) {
		return false
	}
	return true
}
func (this *MsgRemoveRoles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgRemoveRoles)
	if !ok {
		that2, ok := that.(MsgRemoveRoles)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if this.Roles[i] != that1.Roles[i] {
			return false
		}
	}
	if !bytes.Equal(this.Operator, that1.Operator) {
		return false
	}
	return true
}
func (this *MsgBlockAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgBlockAccount)
	if !ok {
		that2, ok := that.(MsgBlockAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if !bytes.Equal(this.Operator, that1.Operator) {
		return false
	}
	return true
}
func (this *MsgUnblockAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUnblockAccount)
	if !ok {
		that2, ok := that.(MsgUnblockAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if !bytes.Equal(this.Operator, that1.Operator) {
		return false
	}
	return true
}
func (m *MsgAddRoles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddRoles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddRoles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Roles) > 0 {
		dAtA2 := make([]byte, len(m.Roles)*10)
		var j1 int
		for _, num := range m.Roles {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTypes(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRemoveRoles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveRoles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveRoles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Roles) > 0 {
		dAtA4 := make([]byte, len(m.Roles)*10)
		var j3 int
		for _, num := range m.Roles {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintTypes(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBlockAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBlockAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBlockAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnblockAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnblockAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnblockAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddRoles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MsgRemoveRoles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Roles) > 0 {
		l = 0
		for _, e := range m.Roles {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MsgBlockAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *MsgUnblockAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddRoles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddRoles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddRoles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Role
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Role(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Roles) == 0 {
					m.Roles = make([]Role, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Role
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Role(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = append(m.Operator[:0], dAtA[iNdEx:postIndex]...)
			if m.Operator == nil {
				m.Operator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRemoveRoles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveRoles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveRoles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Role
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Role(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Roles = append(m.Roles, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Roles) == 0 {
					m.Roles = make([]Role, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Role
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Role(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Roles = append(m.Roles, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = append(m.Operator[:0], dAtA[iNdEx:postIndex]...)
			if m.Operator == nil {
				m.Operator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgBlockAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgBlockAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBlockAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = append(m.Operator[:0], dAtA[iNdEx:postIndex]...)
			if m.Operator == nil {
				m.Operator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnblockAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnblockAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnblockAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = append(m.Operator[:0], dAtA[iNdEx:postIndex]...)
			if m.Operator == nil {
				m.Operator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: auth.proto

package authProto

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

type UserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Role        string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Permissions string `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserRole) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRole) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserRole) GetPermissions() string {
	if x != nil {
		return x.Permissions
	}
	return ""
}

type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AccessToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *RefreshToken) Reset() {
	*x = RefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshToken) ProtoMessage() {}

func (x *RefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshToken.ProtoReflect.Descriptor instead.
func (*RefreshToken) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type GeneratedTokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *GeneratedTokens) Reset() {
	*x = GeneratedTokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratedTokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratedTokens) ProtoMessage() {}

func (x *GeneratedTokens) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratedTokens.ProtoReflect.Descriptor instead.
func (*GeneratedTokens) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *GeneratedTokens) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GeneratedTokens) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Role   string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type ResultBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultBinding) Reset() {
	*x = ResultBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultBinding) ProtoMessage() {}

func (x *ResultBinding) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultBinding.ProtoReflect.Descriptor instead.
func (*ResultBinding) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ResultBinding) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

type Roles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles string `protobuf:"bytes,1,opt,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Roles) Reset() {
	*x = Roles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roles) ProtoMessage() {}

func (x *Roles) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roles.ProtoReflect.Descriptor instead.
func (*Roles) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{7}
}

func (x *Roles) GetRoles() string {
	if x != nil {
		return x.Roles
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x58, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2f, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a,
	0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x27,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x1d, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x32, 0xac, 0x02, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12,
	0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x18, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x17, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0a,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x10, 0x5a, 0x0e, 0x47, 0x52, 0x50, 0x43, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_proto_goTypes = []interface{}{
	(*UserRole)(nil),        // 0: auth.UserRole
	(*AccessToken)(nil),     // 1: auth.AccessToken
	(*RefreshToken)(nil),    // 2: auth.RefreshToken
	(*GeneratedTokens)(nil), // 3: auth.GeneratedTokens
	(*User)(nil),            // 4: auth.User
	(*ResultBinding)(nil),   // 5: auth.ResultBinding
	(*Request)(nil),         // 6: auth.Request
	(*Roles)(nil),           // 7: auth.Roles
}
var file_auth_proto_depIdxs = []int32{
	1, // 0: auth.Auth.GetUserWithRights:input_type -> auth.AccessToken
	4, // 1: auth.Auth.BindUserAndRole:input_type -> auth.User
	2, // 2: auth.Auth.TokenGenerationByRefresh:input_type -> auth.RefreshToken
	4, // 3: auth.Auth.TokenGenerationByUserId:input_type -> auth.User
	6, // 4: auth.Auth.GetAllRoles:input_type -> auth.Request
	0, // 5: auth.Auth.GetUserWithRights:output_type -> auth.UserRole
	5, // 6: auth.Auth.BindUserAndRole:output_type -> auth.ResultBinding
	3, // 7: auth.Auth.TokenGenerationByRefresh:output_type -> auth.GeneratedTokens
	3, // 8: auth.Auth.TokenGenerationByUserId:output_type -> auth.GeneratedTokens
	7, // 9: auth.Auth.GetAllRoles:output_type -> auth.Roles
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRole); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessToken); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshToken); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratedTokens); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultBinding); i {
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
		file_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roles); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

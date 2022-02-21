// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: prot_bee_body.proto

// import "prot_bee_base.proto";

package protbee

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

//cmd=1, login, cmd=2, logout
type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{0}
}

func (x *UserReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Jwt         string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Tel         string   `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty"`
	ExpiredTime int64    `protobuf:"varint,5,opt,name=expired_time,json=expiredTime,proto3" json:"expired_time,omitempty"`
}

func (x *UserRes) Reset() {
	*x = UserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRes) ProtoMessage() {}

func (x *UserRes) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRes.ProtoReflect.Descriptor instead.
func (*UserRes) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{1}
}

func (x *UserRes) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UserRes) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *UserRes) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *UserRes) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *UserRes) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

//thumbnail
type Thumbnail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type      int32  `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	LifeCycle int32  `protobuf:"varint,4,opt,name=life_cycle,json=lifeCycle,proto3" json:"life_cycle,omitempty"`
}

func (x *Thumbnail) Reset() {
	*x = Thumbnail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thumbnail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thumbnail) ProtoMessage() {}

func (x *Thumbnail) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thumbnail.ProtoReflect.Descriptor instead.
func (*Thumbnail) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{2}
}

func (x *Thumbnail) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Thumbnail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Thumbnail) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Thumbnail) GetLifeCycle() int32 {
	if x != nil {
		return x.LifeCycle
	}
	return 0
}

type MapThumbnails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thumbnails []*Thumbnail `protobuf:"bytes,1,rep,name=thumbnails,proto3" json:"thumbnails,omitempty"`
}

func (x *MapThumbnails) Reset() {
	*x = MapThumbnails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapThumbnails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapThumbnails) ProtoMessage() {}

func (x *MapThumbnails) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapThumbnails.ProtoReflect.Descriptor instead.
func (*MapThumbnails) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{3}
}

func (x *MapThumbnails) GetThumbnails() []*Thumbnail {
	if x != nil {
		return x.Thumbnails
	}
	return nil
}

type MapOneDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thumbnail *Thumbnail `protobuf:"bytes,1,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *MapOneDetail) Reset() {
	*x = MapOneDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapOneDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapOneDetail) ProtoMessage() {}

func (x *MapOneDetail) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapOneDetail.ProtoReflect.Descriptor instead.
func (*MapOneDetail) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{4}
}

func (x *MapOneDetail) GetThumbnail() *Thumbnail {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

// cmd=3, /test
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test string `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{5}
}

func (x *Test) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

// cmd=4 /get_users
type UsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UsersReq) Reset() {
	*x = UsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersReq) ProtoMessage() {}

func (x *UsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersReq.ProtoReflect.Descriptor instead.
func (*UsersReq) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{6}
}

type UsersRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserRes `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersRes) Reset() {
	*x = UsersRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRes) ProtoMessage() {}

func (x *UsersRes) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRes.ProtoReflect.Descriptor instead.
func (*UsersRes) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{7}
}

func (x *UsersRes) GetUsers() []*UserRes {
	if x != nil {
		return x.Users
	}
	return nil
}

// cmd=5 /get_roles
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{8}
}

func (x *Role) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RolesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RolesReq) Reset() {
	*x = RolesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolesReq) ProtoMessage() {}

func (x *RolesReq) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolesReq.ProtoReflect.Descriptor instead.
func (*RolesReq) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{9}
}

type RolesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *RolesRes) Reset() {
	*x = RolesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolesRes) ProtoMessage() {}

func (x *RolesRes) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolesRes.ProtoReflect.Descriptor instead.
func (*RolesRes) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{10}
}

func (x *RolesRes) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

// cmd=6 /manage_user 添加用户
// cmd=7 /manage_user 删除用户
// cmd=8 /manage_user 更新用户
type ManageUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User     string  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password string  `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Mobile   string  `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Roles    []*Role `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *ManageUserReq) Reset() {
	*x = ManageUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageUserReq) ProtoMessage() {}

func (x *ManageUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageUserReq.ProtoReflect.Descriptor instead.
func (*ManageUserReq) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{11}
}

func (x *ManageUserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ManageUserReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ManageUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ManageUserReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ManageUserReq) GetRoles() []*Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type ManageUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManageUserRes) Reset() {
	*x = ManageUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prot_bee_body_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageUserRes) ProtoMessage() {}

func (x *ManageUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_prot_bee_body_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageUserRes.ProtoReflect.Descriptor instead.
func (*ManageUserRes) Descriptor() ([]byte, []int) {
	return file_prot_bee_body_proto_rawDescGZIP(), []int{12}
}

var File_prot_bee_body_proto protoreflect.FileDescriptor

var file_prot_bee_body_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x62, 0x65, 0x65, 0x22, 0x39,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x07, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x66, 0x0a, 0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x69, 0x66, 0x65, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6c, 0x69, 0x66, 0x65, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x4d, 0x61,
	0x70, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x62, 0x65, 0x65, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x52, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0x40, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x4f, 0x6e, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x30, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x62, 0x65, 0x65, 0x2e, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x22, 0x1a, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0x0a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x22, 0x32, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x62, 0x65, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x3e, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0a, 0x0a,
	0x08, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x22, 0x2f, 0x0a, 0x08, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x62, 0x65, 0x65, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x62, 0x65, 0x65, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x26, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x74, 0x62, 0x65, 0x65, 0x2e, 0x70, 0x62, 0x66, 0x48, 0x03, 0x5a, 0x15, 0x7a, 0x6b, 0x2d,
	0x61, 0x74, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x70, 0x62, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x62,
	0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prot_bee_body_proto_rawDescOnce sync.Once
	file_prot_bee_body_proto_rawDescData = file_prot_bee_body_proto_rawDesc
)

func file_prot_bee_body_proto_rawDescGZIP() []byte {
	file_prot_bee_body_proto_rawDescOnce.Do(func() {
		file_prot_bee_body_proto_rawDescData = protoimpl.X.CompressGZIP(file_prot_bee_body_proto_rawDescData)
	})
	return file_prot_bee_body_proto_rawDescData
}

var file_prot_bee_body_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_prot_bee_body_proto_goTypes = []interface{}{
	(*UserReq)(nil),       // 0: protbee.UserReq
	(*UserRes)(nil),       // 1: protbee.UserRes
	(*Thumbnail)(nil),     // 2: protbee.Thumbnail
	(*MapThumbnails)(nil), // 3: protbee.MapThumbnails
	(*MapOneDetail)(nil),  // 4: protbee.MapOneDetail
	(*Test)(nil),          // 5: protbee.Test
	(*UsersReq)(nil),      // 6: protbee.UsersReq
	(*UsersRes)(nil),      // 7: protbee.UsersRes
	(*Role)(nil),          // 8: protbee.Role
	(*RolesReq)(nil),      // 9: protbee.RolesReq
	(*RolesRes)(nil),      // 10: protbee.RolesRes
	(*ManageUserReq)(nil), // 11: protbee.ManageUserReq
	(*ManageUserRes)(nil), // 12: protbee.ManageUserRes
}
var file_prot_bee_body_proto_depIdxs = []int32{
	2, // 0: protbee.MapThumbnails.thumbnails:type_name -> protbee.Thumbnail
	2, // 1: protbee.MapOneDetail.thumbnail:type_name -> protbee.Thumbnail
	1, // 2: protbee.UsersRes.users:type_name -> protbee.UserRes
	8, // 3: protbee.RolesRes.roles:type_name -> protbee.Role
	8, // 4: protbee.ManageUserReq.roles:type_name -> protbee.Role
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_prot_bee_body_proto_init() }
func file_prot_bee_body_proto_init() {
	if File_prot_bee_body_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prot_bee_body_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
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
		file_prot_bee_body_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRes); i {
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
		file_prot_bee_body_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thumbnail); i {
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
		file_prot_bee_body_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapThumbnails); i {
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
		file_prot_bee_body_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapOneDetail); i {
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
		file_prot_bee_body_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_prot_bee_body_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersReq); i {
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
		file_prot_bee_body_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersRes); i {
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
		file_prot_bee_body_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_prot_bee_body_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolesReq); i {
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
		file_prot_bee_body_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolesRes); i {
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
		file_prot_bee_body_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageUserReq); i {
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
		file_prot_bee_body_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageUserRes); i {
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
			RawDescriptor: file_prot_bee_body_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prot_bee_body_proto_goTypes,
		DependencyIndexes: file_prot_bee_body_proto_depIdxs,
		MessageInfos:      file_prot_bee_body_proto_msgTypes,
	}.Build()
	File_prot_bee_body_proto = out.File
	file_prot_bee_body_proto_rawDesc = nil
	file_prot_bee_body_proto_goTypes = nil
	file_prot_bee_body_proto_depIdxs = nil
}

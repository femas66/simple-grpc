// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: simplecrud/user.proto

package simplecrud

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

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_simplecrud_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_simplecrud_user_proto_msgTypes[0]
	if x != nil {
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
	return file_simplecrud_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*User                `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserList) Reset() {
	*x = UserList{}
	mi := &file_simplecrud_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_simplecrud_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_simplecrud_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserList) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

type UserID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserID) Reset() {
	*x = UserID{}
	mi := &file_simplecrud_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_simplecrud_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_simplecrud_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_simplecrud_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_simplecrud_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_simplecrud_user_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_simplecrud_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_simplecrud_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_simplecrud_user_proto_rawDescGZIP(), []int{4}
}

var File_simplecrud_user_proto protoreflect.FileDescriptor

var file_simplecrud_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x63,
	0x72, 0x75, 0x64, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x30, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xe7, 0x01, 0x0a, 0x0b, 0x43,
	0x52, 0x55, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x14, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x63, 0x72, 0x75,
	0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x14, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x63, 0x72, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simplecrud_user_proto_rawDescOnce sync.Once
	file_simplecrud_user_proto_rawDescData = file_simplecrud_user_proto_rawDesc
)

func file_simplecrud_user_proto_rawDescGZIP() []byte {
	file_simplecrud_user_proto_rawDescOnce.Do(func() {
		file_simplecrud_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_simplecrud_user_proto_rawDescData)
	})
	return file_simplecrud_user_proto_rawDescData
}

var file_simplecrud_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_simplecrud_user_proto_goTypes = []any{
	(*User)(nil),     // 0: simplecrud.User
	(*UserList)(nil), // 1: simplecrud.UserList
	(*UserID)(nil),   // 2: simplecrud.UserID
	(*Response)(nil), // 3: simplecrud.Response
	(*Empty)(nil),    // 4: simplecrud.Empty
}
var file_simplecrud_user_proto_depIdxs = []int32{
	0, // 0: simplecrud.UserList.list:type_name -> simplecrud.User
	0, // 1: simplecrud.CRUDService.CreateUser:input_type -> simplecrud.User
	0, // 2: simplecrud.CRUDService.UpdateUser:input_type -> simplecrud.User
	2, // 3: simplecrud.CRUDService.DeleteUser:input_type -> simplecrud.UserID
	4, // 4: simplecrud.CRUDService.ReadUsers:input_type -> simplecrud.Empty
	3, // 5: simplecrud.CRUDService.CreateUser:output_type -> simplecrud.Response
	3, // 6: simplecrud.CRUDService.UpdateUser:output_type -> simplecrud.Response
	3, // 7: simplecrud.CRUDService.DeleteUser:output_type -> simplecrud.Response
	1, // 8: simplecrud.CRUDService.ReadUsers:output_type -> simplecrud.UserList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_simplecrud_user_proto_init() }
func file_simplecrud_user_proto_init() {
	if File_simplecrud_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_simplecrud_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simplecrud_user_proto_goTypes,
		DependencyIndexes: file_simplecrud_user_proto_depIdxs,
		MessageInfos:      file_simplecrud_user_proto_msgTypes,
	}.Build()
	File_simplecrud_user_proto = out.File
	file_simplecrud_user_proto_rawDesc = nil
	file_simplecrud_user_proto_goTypes = nil
	file_simplecrud_user_proto_depIdxs = nil
}
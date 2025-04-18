// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0
// source: proto/main.proto

package become_better

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MainCategoriesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MainCategoriesRequest) Reset() {
	*x = MainCategoriesRequest{}
	mi := &file_proto_main_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MainCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainCategoriesRequest) ProtoMessage() {}

func (x *MainCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainCategoriesRequest.ProtoReflect.Descriptor instead.
func (*MainCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{0}
}

type EmptyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	mi := &file_proto_main_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{1}
}

type MainCategoriesResponse struct {
	state          protoimpl.MessageState   `protogen:"open.v1"`
	MainCategories []*MainCategoriesMessage `protobuf:"bytes,1,rep,name=MainCategories,proto3" json:"MainCategories,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MainCategoriesResponse) Reset() {
	*x = MainCategoriesResponse{}
	mi := &file_proto_main_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MainCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainCategoriesResponse) ProtoMessage() {}

func (x *MainCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainCategoriesResponse.ProtoReflect.Descriptor instead.
func (*MainCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{2}
}

func (x *MainCategoriesResponse) GetMainCategories() []*MainCategoriesMessage {
	if x != nil {
		return x.MainCategories
	}
	return nil
}

type MainCategoriesMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	MainCategory  string                 `protobuf:"bytes,4,opt,name=main_category,json=mainCategory,proto3" json:"main_category,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MainCategoriesMessage) Reset() {
	*x = MainCategoriesMessage{}
	mi := &file_proto_main_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MainCategoriesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainCategoriesMessage) ProtoMessage() {}

func (x *MainCategoriesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainCategoriesMessage.ProtoReflect.Descriptor instead.
func (*MainCategoriesMessage) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{3}
}

func (x *MainCategoriesMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MainCategoriesMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MainCategoriesMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MainCategoriesMessage) GetMainCategory() string {
	if x != nil {
		return x.MainCategory
	}
	return ""
}

type AddCategoryMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MainCategory  int32                  `protobuf:"varint,3,opt,name=main_category,json=mainCategory,proto3" json:"main_category,omitempty"`
	CategoryType  int32                  `protobuf:"varint,4,opt,name=category_type,json=categoryType,proto3" json:"category_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCategoryMessage) Reset() {
	*x = AddCategoryMessage{}
	mi := &file_proto_main_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCategoryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCategoryMessage) ProtoMessage() {}

func (x *AddCategoryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCategoryMessage.ProtoReflect.Descriptor instead.
func (*AddCategoryMessage) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{4}
}

func (x *AddCategoryMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCategoryMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddCategoryMessage) GetMainCategory() int32 {
	if x != nil {
		return x.MainCategory
	}
	return 0
}

func (x *AddCategoryMessage) GetCategoryType() int32 {
	if x != nil {
		return x.CategoryType
	}
	return 0
}

type DeleteProgressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProgressId    string                 `protobuf:"bytes,1,opt,name=progress_id,json=progressId,proto3" json:"progress_id,omitempty"` // uuid прогресса
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProgressRequest) Reset() {
	*x = DeleteProgressRequest{}
	mi := &file_proto_main_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProgressRequest) ProtoMessage() {}

func (x *DeleteProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProgressRequest.ProtoReflect.Descriptor instead.
func (*DeleteProgressRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProgressRequest) GetProgressId() string {
	if x != nil {
		return x.ProgressId
	}
	return ""
}

func (x *DeleteProgressRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FillProgressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CategoryId    string                 `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"` // uuid категории
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Result        int32                  `protobuf:"varint,4,opt,name=result,proto3" json:"result,omitempty"` // Минуты, страницы и тп(количественное значение)
	Date          string                 `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`      // Когда было сделано DD.MM.YYYY
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FillProgressRequest) Reset() {
	*x = FillProgressRequest{}
	mi := &file_proto_main_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FillProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FillProgressRequest) ProtoMessage() {}

func (x *FillProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FillProgressRequest.ProtoReflect.Descriptor instead.
func (*FillProgressRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{6}
}

func (x *FillProgressRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *FillProgressRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FillProgressRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FillProgressRequest) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *FillProgressRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetProgressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CategoryId    string                 `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"` // uuid категории
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DateFrom      string                 `protobuf:"bytes,3,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"` // Когда было сделано DD.MM.YYYY
	DateTo        string                 `protobuf:"bytes,4,opt,name=dateTo,proto3" json:"dateTo,omitempty"`     // Когда было сделано DD.MM.YYYY
	Page          int32                  `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProgressRequest) Reset() {
	*x = GetProgressRequest{}
	mi := &file_proto_main_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgressRequest) ProtoMessage() {}

func (x *GetProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgressRequest.ProtoReflect.Descriptor instead.
func (*GetProgressRequest) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{7}
}

func (x *GetProgressRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *GetProgressRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetProgressRequest) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *GetProgressRequest) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *GetProgressRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetProgressRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetProgressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Progress      []*Progress            `protobuf:"bytes,1,rep,name=progress,proto3" json:"progress,omitempty"`
	CountPages    int32                  `protobuf:"varint,2,opt,name=count_pages,json=countPages,proto3" json:"count_pages,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProgressResponse) Reset() {
	*x = GetProgressResponse{}
	mi := &file_proto_main_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProgressResponse) ProtoMessage() {}

func (x *GetProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProgressResponse.ProtoReflect.Descriptor instead.
func (*GetProgressResponse) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{8}
}

func (x *GetProgressResponse) GetProgress() []*Progress {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *GetProgressResponse) GetCountPages() int32 {
	if x != nil {
		return x.CountPages
	}
	return 0
}

func (x *GetProgressResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetProgressResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Progress struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ID                string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`                                   // uuid
	CategoryId        string                 `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"` // uuid категории
	UserId            string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Date              string                 `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"` // Когда было сделано DD.MM.YYYY
	ProgressType      string                 `protobuf:"bytes,5,opt,name=progress_type,json=progressType,proto3" json:"progress_type,omitempty"`
	ResultInt         int32                  `protobuf:"varint,6,opt,name=result_int,json=resultInt,proto3" json:"result_int,omitempty"`
	ResultDescription string                 `protobuf:"bytes,7,opt,name=result_description,json=resultDescription,proto3" json:"result_description,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Progress) Reset() {
	*x = Progress{}
	mi := &file_proto_main_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress) ProtoMessage() {}

func (x *Progress) ProtoReflect() protoreflect.Message {
	mi := &file_proto_main_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress.ProtoReflect.Descriptor instead.
func (*Progress) Descriptor() ([]byte, []int) {
	return file_proto_main_proto_rawDescGZIP(), []int{9}
}

func (x *Progress) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Progress) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Progress) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Progress) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Progress) GetProgressType() string {
	if x != nil {
		return x.ProgressType
	}
	return ""
}

func (x *Progress) GetResultInt() int32 {
	if x != nil {
		return x.ResultInt
	}
	return 0
}

func (x *Progress) GetResultDescription() string {
	if x != nil {
		return x.ResultDescription
	}
	return ""
}

var File_proto_main_proto protoreflect.FileDescriptor

var file_proto_main_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x4d, 0x61, 0x69,
	0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x16, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x51, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f,
	0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x49, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0x82, 0x04, 0x0a, 0x0c, 0x42, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x42, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x6e, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x69, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x1e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x69,
	0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x57, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x5e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x5e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x62, 0x65, 0x74, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_main_proto_rawDescOnce sync.Once
	file_proto_main_proto_rawDescData = file_proto_main_proto_rawDesc
)

func file_proto_main_proto_rawDescGZIP() []byte {
	file_proto_main_proto_rawDescOnce.Do(func() {
		file_proto_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_main_proto_rawDescData)
	})
	return file_proto_main_proto_rawDescData
}

var file_proto_main_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_main_proto_goTypes = []any{
	(*MainCategoriesRequest)(nil),  // 0: example.MainCategoriesRequest
	(*EmptyResponse)(nil),          // 1: example.EmptyResponse
	(*MainCategoriesResponse)(nil), // 2: example.MainCategoriesResponse
	(*MainCategoriesMessage)(nil),  // 3: example.MainCategoriesMessage
	(*AddCategoryMessage)(nil),     // 4: example.AddCategoryMessage
	(*DeleteProgressRequest)(nil),  // 5: example.DeleteProgressRequest
	(*FillProgressRequest)(nil),    // 6: example.FillProgressRequest
	(*GetProgressRequest)(nil),     // 7: example.GetProgressRequest
	(*GetProgressResponse)(nil),    // 8: example.GetProgressResponse
	(*Progress)(nil),               // 9: example.Progress
}
var file_proto_main_proto_depIdxs = []int32{
	3, // 0: example.MainCategoriesResponse.MainCategories:type_name -> example.MainCategoriesMessage
	9, // 1: example.GetProgressResponse.progress:type_name -> example.Progress
	0, // 2: example.BecomeBetter.MainCategories:input_type -> example.MainCategoriesRequest
	4, // 3: example.BecomeBetter.AddCategories:input_type -> example.AddCategoryMessage
	6, // 4: example.BecomeBetter.FillProgress:input_type -> example.FillProgressRequest
	5, // 5: example.BecomeBetter.DeleteProgress:input_type -> example.DeleteProgressRequest
	7, // 6: example.BecomeBetter.GetProgress:input_type -> example.GetProgressRequest
	2, // 7: example.BecomeBetter.MainCategories:output_type -> example.MainCategoriesResponse
	3, // 8: example.BecomeBetter.AddCategories:output_type -> example.MainCategoriesMessage
	1, // 9: example.BecomeBetter.FillProgress:output_type -> example.EmptyResponse
	1, // 10: example.BecomeBetter.DeleteProgress:output_type -> example.EmptyResponse
	8, // 11: example.BecomeBetter.GetProgress:output_type -> example.GetProgressResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_main_proto_init() }
func file_proto_main_proto_init() {
	if File_proto_main_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_main_proto_goTypes,
		DependencyIndexes: file_proto_main_proto_depIdxs,
		MessageInfos:      file_proto_main_proto_msgTypes,
	}.Build()
	File_proto_main_proto = out.File
	file_proto_main_proto_rawDesc = nil
	file_proto_main_proto_goTypes = nil
	file_proto_main_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: douyin_publish.proto

package publish

import (
	context "context"
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

type DouyinPublishActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	Data  []byte  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	Title *string `protobuf:"bytes,3,req,name=title" json:"title,omitempty"`
}

func (x *DouyinPublishActionRequest) Reset() {
	*x = DouyinPublishActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishActionRequest) ProtoMessage() {}

func (x *DouyinPublishActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinPublishActionRequest) Descriptor() ([]byte, []int) {
	return file_douyin_publish_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinPublishActionRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *DouyinPublishActionRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DouyinPublishActionRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type DouyinPublishActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`
}

func (x *DouyinPublishActionResponse) Reset() {
	*x = DouyinPublishActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishActionResponse) ProtoMessage() {}

func (x *DouyinPublishActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinPublishActionResponse) Descriptor() ([]byte, []int) {
	return file_douyin_publish_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinPublishActionResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinPublishActionResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

type DouyinPublishListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *int64  `protobuf:"varint,1,req,name=user_id,json=userId" json:"user_id,omitempty"`
	Token  *string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
}

func (x *DouyinPublishListRequest) Reset() {
	*x = DouyinPublishListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_publish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishListRequest) ProtoMessage() {}

func (x *DouyinPublishListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_publish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishListRequest.ProtoReflect.Descriptor instead.
func (*DouyinPublishListRequest) Descriptor() ([]byte, []int) {
	return file_douyin_publish_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinPublishListRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *DouyinPublishListRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type DouyinPublishListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode *int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"`
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"`
}

func (x *DouyinPublishListResponse) Reset() {
	*x = DouyinPublishListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_publish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishListResponse) ProtoMessage() {}

func (x *DouyinPublishListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_publish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishListResponse.ProtoReflect.Descriptor instead.
func (*DouyinPublishListResponse) Descriptor() ([]byte, []int) {
	return file_douyin_publish_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinPublishListResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *DouyinPublishListResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinPublishListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`                                            //视频唯一标识
	Author        *User   `protobuf:"bytes,2,req,name=author" json:"author,omitempty"`                                     //视频作者信息
	PlayUrl       *string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,omitempty"`                    //视频播放地址
	CoverUrl      *string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,omitempty"`                 //视频封面地址
	FavoriteCount *int64  `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"` //视频的点赞总数
	CommentCount  *int64  `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,omitempty"`    //视频的评论总数
	IsFavorite    *bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,omitempty"`          //true-已点赞 false-未点赞
	Title         *string `protobuf:"bytes,8,req,name=title" json:"title,omitempty"`                                       //视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_publish_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_publish_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_douyin_publish_proto_rawDescGZIP(), []int{4}
}

func (x *Video) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil && x.PlayUrl != nil {
		return *x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil && x.CoverUrl != nil {
		return *x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil && x.FavoriteCount != nil {
		return *x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil && x.CommentCount != nil {
		return *x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil && x.IsFavorite != nil {
		return *x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`                                            //用户id
	Name          *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`                                         //用户名称
	FollowCount   *int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count,omitempty"`       //关注总数
	FollowerCount *int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count,omitempty"` //粉丝数量
	IsFollow      *bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow" json:"is_follow,omitempty"`                //true-已关注，false-未关注
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_publish_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_publish_proto_msgTypes[5]
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
	return file_douyin_publish_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil && x.FollowCount != nil {
		return *x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil && x.FollowerCount != nil {
		return *x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil && x.IsFollow != nil {
		return *x.IsFollow
	}
	return false
}

var File_douyin_publish_proto protoreflect.FileDescriptor

var file_douyin_publish_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x22,
	0x5f, 0x0a, 0x1d, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x60, 0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x22, 0x4c, 0x0a, 0x1b, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x8d, 0x01, 0x0a, 0x1c, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x2d, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0xf9, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x08, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x91, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x42, 0x24, 0x5a, 0x22, 0x6d, 0x79, 0x64, 0x6f, 0x75, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
}

var (
	file_douyin_publish_proto_rawDescOnce sync.Once
	file_douyin_publish_proto_rawDescData = file_douyin_publish_proto_rawDesc
)

func file_douyin_publish_proto_rawDescGZIP() []byte {
	file_douyin_publish_proto_rawDescOnce.Do(func() {
		file_douyin_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_douyin_publish_proto_rawDescData)
	})
	return file_douyin_publish_proto_rawDescData
}

var file_douyin_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_douyin_publish_proto_goTypes = []interface{}{
	(*DouyinPublishActionRequest)(nil),  // 0: publish.douyin_publish_action_request
	(*DouyinPublishActionResponse)(nil), // 1: publish.douyin_publish_action_response
	(*DouyinPublishListRequest)(nil),    // 2: publish.douyin_publish_list_request
	(*DouyinPublishListResponse)(nil),   // 3: publish.douyin_publish_list_response
	(*Video)(nil),                       // 4: publish.Video
	(*User)(nil),                        // 5: publish.User
}
var file_douyin_publish_proto_depIdxs = []int32{
	4, // 0: publish.douyin_publish_list_response.video_list:type_name -> publish.Video
	5, // 1: publish.Video.author:type_name -> publish.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_douyin_publish_proto_init() }
func file_douyin_publish_proto_init() {
	if File_douyin_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_douyin_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishActionRequest); i {
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
		file_douyin_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishActionResponse); i {
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
		file_douyin_publish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishListRequest); i {
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
		file_douyin_publish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishListResponse); i {
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
		file_douyin_publish_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_douyin_publish_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_douyin_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_douyin_publish_proto_goTypes,
		DependencyIndexes: file_douyin_publish_proto_depIdxs,
		MessageInfos:      file_douyin_publish_proto_msgTypes,
	}.Build()
	File_douyin_publish_proto = out.File
	file_douyin_publish_proto_rawDesc = nil
	file_douyin_publish_proto_goTypes = nil
	file_douyin_publish_proto_depIdxs = nil
}

var _ context.Context

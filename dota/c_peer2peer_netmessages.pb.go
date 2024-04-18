// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.5
// source: c_peer2peer_netmessages.proto

package dota

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

type P2P_Messages int32

const (
	P2P_Messages_p2p_TextMessage          P2P_Messages = 256
	P2P_Messages_p2p_Voice                P2P_Messages = 257
	P2P_Messages_p2p_Ping                 P2P_Messages = 258
	P2P_Messages_p2p_VRAvatarPosition     P2P_Messages = 259
	P2P_Messages_p2p_WatchSynchronization P2P_Messages = 260
)

// Enum value maps for P2P_Messages.
var (
	P2P_Messages_name = map[int32]string{
		256: "p2p_TextMessage",
		257: "p2p_Voice",
		258: "p2p_Ping",
		259: "p2p_VRAvatarPosition",
		260: "p2p_WatchSynchronization",
	}
	P2P_Messages_value = map[string]int32{
		"p2p_TextMessage":          256,
		"p2p_Voice":                257,
		"p2p_Ping":                 258,
		"p2p_VRAvatarPosition":     259,
		"p2p_WatchSynchronization": 260,
	}
)

func (x P2P_Messages) Enum() *P2P_Messages {
	p := new(P2P_Messages)
	*p = x
	return p
}

func (x P2P_Messages) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (P2P_Messages) Descriptor() protoreflect.EnumDescriptor {
	return file_c_peer2peer_netmessages_proto_enumTypes[0].Descriptor()
}

func (P2P_Messages) Type() protoreflect.EnumType {
	return &file_c_peer2peer_netmessages_proto_enumTypes[0]
}

func (x P2P_Messages) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *P2P_Messages) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = P2P_Messages(num)
	return nil
}

// Deprecated: Use P2P_Messages.Descriptor instead.
func (P2P_Messages) EnumDescriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{0}
}

type CP2P_Voice_Handler_Flags int32

const (
	CP2P_Voice_Played_Audio CP2P_Voice_Handler_Flags = 1
)

// Enum value maps for CP2P_Voice_Handler_Flags.
var (
	CP2P_Voice_Handler_Flags_name = map[int32]string{
		1: "Played_Audio",
	}
	CP2P_Voice_Handler_Flags_value = map[string]int32{
		"Played_Audio": 1,
	}
)

func (x CP2P_Voice_Handler_Flags) Enum() *CP2P_Voice_Handler_Flags {
	p := new(CP2P_Voice_Handler_Flags)
	*p = x
	return p
}

func (x CP2P_Voice_Handler_Flags) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CP2P_Voice_Handler_Flags) Descriptor() protoreflect.EnumDescriptor {
	return file_c_peer2peer_netmessages_proto_enumTypes[1].Descriptor()
}

func (CP2P_Voice_Handler_Flags) Type() protoreflect.EnumType {
	return &file_c_peer2peer_netmessages_proto_enumTypes[1]
}

func (x CP2P_Voice_Handler_Flags) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CP2P_Voice_Handler_Flags) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CP2P_Voice_Handler_Flags(num)
	return nil
}

// Deprecated: Use CP2P_Voice_Handler_Flags.Descriptor instead.
func (CP2P_Voice_Handler_Flags) EnumDescriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{2, 0}
}

type CP2P_TextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text []byte `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (x *CP2P_TextMessage) Reset() {
	*x = CP2P_TextMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c_peer2peer_netmessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CP2P_TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CP2P_TextMessage) ProtoMessage() {}

func (x *CP2P_TextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_c_peer2peer_netmessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CP2P_TextMessage.ProtoReflect.Descriptor instead.
func (*CP2P_TextMessage) Descriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{0}
}

func (x *CP2P_TextMessage) GetText() []byte {
	if x != nil {
		return x.Text
	}
	return nil
}

type CSteam_Voice_Encoding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceData []byte `protobuf:"bytes,1,opt,name=voice_data,json=voiceData" json:"voice_data,omitempty"`
}

func (x *CSteam_Voice_Encoding) Reset() {
	*x = CSteam_Voice_Encoding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c_peer2peer_netmessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteam_Voice_Encoding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteam_Voice_Encoding) ProtoMessage() {}

func (x *CSteam_Voice_Encoding) ProtoReflect() protoreflect.Message {
	mi := &file_c_peer2peer_netmessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteam_Voice_Encoding.ProtoReflect.Descriptor instead.
func (*CSteam_Voice_Encoding) Descriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{1}
}

func (x *CSteam_Voice_Encoding) GetVoiceData() []byte {
	if x != nil {
		return x.VoiceData
	}
	return nil
}

type CP2P_Voice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audio          *CMsgVoiceAudio `protobuf:"bytes,1,opt,name=audio" json:"audio,omitempty"`
	BroadcastGroup *uint32         `protobuf:"varint,2,opt,name=broadcast_group,json=broadcastGroup" json:"broadcast_group,omitempty"`
}

func (x *CP2P_Voice) Reset() {
	*x = CP2P_Voice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c_peer2peer_netmessages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CP2P_Voice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CP2P_Voice) ProtoMessage() {}

func (x *CP2P_Voice) ProtoReflect() protoreflect.Message {
	mi := &file_c_peer2peer_netmessages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CP2P_Voice.ProtoReflect.Descriptor instead.
func (*CP2P_Voice) Descriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{2}
}

func (x *CP2P_Voice) GetAudio() *CMsgVoiceAudio {
	if x != nil {
		return x.Audio
	}
	return nil
}

func (x *CP2P_Voice) GetBroadcastGroup() uint32 {
	if x != nil && x.BroadcastGroup != nil {
		return *x.BroadcastGroup
	}
	return 0
}

type CP2P_Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendTime *uint64 `protobuf:"varint,1,req,name=send_time,json=sendTime" json:"send_time,omitempty"`
	IsReply  *bool   `protobuf:"varint,2,req,name=is_reply,json=isReply" json:"is_reply,omitempty"`
}

func (x *CP2P_Ping) Reset() {
	*x = CP2P_Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c_peer2peer_netmessages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CP2P_Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CP2P_Ping) ProtoMessage() {}

func (x *CP2P_Ping) ProtoReflect() protoreflect.Message {
	mi := &file_c_peer2peer_netmessages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CP2P_Ping.ProtoReflect.Descriptor instead.
func (*CP2P_Ping) Descriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{3}
}

func (x *CP2P_Ping) GetSendTime() uint64 {
	if x != nil && x.SendTime != nil {
		return *x.SendTime
	}
	return 0
}

func (x *CP2P_Ping) GetIsReply() bool {
	if x != nil && x.IsReply != nil {
		return *x.IsReply
	}
	return false
}

type CP2P_VRAvatarPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BodyParts  []*CP2P_VRAvatarPosition_COrientation `protobuf:"bytes,1,rep,name=body_parts,json=bodyParts" json:"body_parts,omitempty"`
	HatId      *int32                                `protobuf:"varint,2,opt,name=hat_id,json=hatId" json:"hat_id,omitempty"`
	SceneId    *int32                                `protobuf:"varint,3,opt,name=scene_id,json=sceneId" json:"scene_id,omitempty"`
	WorldScale *int32                                `protobuf:"varint,4,opt,name=world_scale,json=worldScale" json:"world_scale,omitempty"`
}

func (x *CP2P_VRAvatarPosition) Reset() {
	*x = CP2P_VRAvatarPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c_peer2peer_netmessages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CP2P_VRAvatarPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CP2P_VRAvatarPosition) ProtoMessage() {}

func (x *CP2P_VRAvatarPosition) ProtoReflect() protoreflect.Message {
	mi := &file_c_peer2peer_netmessages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CP2P_VRAvatarPosition.ProtoReflect.Descriptor instead.
func (*CP2P_VRAvatarPosition) Descriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{4}
}

func (x *CP2P_VRAvatarPosition) GetBodyParts() []*CP2P_VRAvatarPosition_COrientation {
	if x != nil {
		return x.BodyParts
	}
	return nil
}

func (x *CP2P_VRAvatarPosition) GetHatId() int32 {
	if x != nil && x.HatId != nil {
		return *x.HatId
	}
	return 0
}

func (x *CP2P_VRAvatarPosition) GetSceneId() int32 {
	if x != nil && x.SceneId != nil {
		return *x.SceneId
	}
	return 0
}

func (x *CP2P_VRAvatarPosition) GetWorldScale() int32 {
	if x != nil && x.WorldScale != nil {
		return *x.WorldScale
	}
	return 0
}

type CP2P_WatchSynchronization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DemoTick                         *int32  `protobuf:"varint,1,opt,name=demo_tick,json=demoTick" json:"demo_tick,omitempty"`
	Paused                           *bool   `protobuf:"varint,2,opt,name=paused" json:"paused,omitempty"`
	TvListenVoiceIndices             *uint64 `protobuf:"varint,3,opt,name=tv_listen_voice_indices,json=tvListenVoiceIndices" json:"tv_listen_voice_indices,omitempty"`
	DotaSpectatorMode                *int32  `protobuf:"varint,4,opt,name=dota_spectator_mode,json=dotaSpectatorMode" json:"dota_spectator_mode,omitempty"`
	DotaSpectatorWatchingBroadcaster *bool   `protobuf:"varint,5,opt,name=dota_spectator_watching_broadcaster,json=dotaSpectatorWatchingBroadcaster" json:"dota_spectator_watching_broadcaster,omitempty"`
	DotaSpectatorHeroIndex           *int32  `protobuf:"varint,6,opt,name=dota_spectator_hero_index,json=dotaSpectatorHeroIndex" json:"dota_spectator_hero_index,omitempty"`
	DotaSpectatorAutospeed           *int32  `protobuf:"varint,7,opt,name=dota_spectator_autospeed,json=dotaSpectatorAutospeed" json:"dota_spectator_autospeed,omitempty"`
	DotaReplaySpeed                  *int32  `protobuf:"varint,8,opt,name=dota_replay_speed,json=dotaReplaySpeed" json:"dota_replay_speed,omitempty"`
}

func (x *CP2P_WatchSynchronization) Reset() {
	*x = CP2P_WatchSynchronization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c_peer2peer_netmessages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CP2P_WatchSynchronization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CP2P_WatchSynchronization) ProtoMessage() {}

func (x *CP2P_WatchSynchronization) ProtoReflect() protoreflect.Message {
	mi := &file_c_peer2peer_netmessages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CP2P_WatchSynchronization.ProtoReflect.Descriptor instead.
func (*CP2P_WatchSynchronization) Descriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{5}
}

func (x *CP2P_WatchSynchronization) GetDemoTick() int32 {
	if x != nil && x.DemoTick != nil {
		return *x.DemoTick
	}
	return 0
}

func (x *CP2P_WatchSynchronization) GetPaused() bool {
	if x != nil && x.Paused != nil {
		return *x.Paused
	}
	return false
}

func (x *CP2P_WatchSynchronization) GetTvListenVoiceIndices() uint64 {
	if x != nil && x.TvListenVoiceIndices != nil {
		return *x.TvListenVoiceIndices
	}
	return 0
}

func (x *CP2P_WatchSynchronization) GetDotaSpectatorMode() int32 {
	if x != nil && x.DotaSpectatorMode != nil {
		return *x.DotaSpectatorMode
	}
	return 0
}

func (x *CP2P_WatchSynchronization) GetDotaSpectatorWatchingBroadcaster() bool {
	if x != nil && x.DotaSpectatorWatchingBroadcaster != nil {
		return *x.DotaSpectatorWatchingBroadcaster
	}
	return false
}

func (x *CP2P_WatchSynchronization) GetDotaSpectatorHeroIndex() int32 {
	if x != nil && x.DotaSpectatorHeroIndex != nil {
		return *x.DotaSpectatorHeroIndex
	}
	return 0
}

func (x *CP2P_WatchSynchronization) GetDotaSpectatorAutospeed() int32 {
	if x != nil && x.DotaSpectatorAutospeed != nil {
		return *x.DotaSpectatorAutospeed
	}
	return 0
}

func (x *CP2P_WatchSynchronization) GetDotaReplaySpeed() int32 {
	if x != nil && x.DotaReplaySpeed != nil {
		return *x.DotaReplaySpeed
	}
	return 0
}

type CP2P_VRAvatarPosition_COrientation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos *CMsgVector `protobuf:"bytes,1,opt,name=pos" json:"pos,omitempty"`
	Ang *CMsgQAngle `protobuf:"bytes,2,opt,name=ang" json:"ang,omitempty"`
}

func (x *CP2P_VRAvatarPosition_COrientation) Reset() {
	*x = CP2P_VRAvatarPosition_COrientation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_c_peer2peer_netmessages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CP2P_VRAvatarPosition_COrientation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CP2P_VRAvatarPosition_COrientation) ProtoMessage() {}

func (x *CP2P_VRAvatarPosition_COrientation) ProtoReflect() protoreflect.Message {
	mi := &file_c_peer2peer_netmessages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CP2P_VRAvatarPosition_COrientation.ProtoReflect.Descriptor instead.
func (*CP2P_VRAvatarPosition_COrientation) Descriptor() ([]byte, []int) {
	return file_c_peer2peer_netmessages_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CP2P_VRAvatarPosition_COrientation) GetPos() *CMsgVector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *CP2P_VRAvatarPosition_COrientation) GetAng() *CMsgQAngle {
	if x != nil {
		return x.Ang
	}
	return nil
}

var File_c_peer2peer_netmessages_proto protoreflect.FileDescriptor

var file_c_peer2peer_netmessages_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x32, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x6e, 0x65,
	0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x64, 0x6f, 0x74, 0x61, 0x1a, 0x11, 0x6e, 0x65, 0x74, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x26, 0x0a, 0x10, 0x43, 0x50, 0x32, 0x50, 0x5f, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x36, 0x0a, 0x15, 0x43, 0x53, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x84, 0x01, 0x0a, 0x0a, 0x43, 0x50, 0x32, 0x50, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x21, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x5f,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x01, 0x22, 0x43, 0x0a, 0x09, 0x43, 0x50, 0x32, 0x50, 0x5f,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x8b, 0x02, 0x0a,
	0x15, 0x43, 0x50, 0x32, 0x50, 0x5f, 0x56, 0x52, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x6f, 0x74,
	0x61, 0x2e, 0x43, 0x50, 0x32, 0x50, 0x5f, 0x56, 0x52, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x50, 0x61, 0x72, 0x74, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x1a, 0x56, 0x0a, 0x0c, 0x43, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x22, 0x0a, 0x03, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x51,
	0x41, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x03, 0x61, 0x6e, 0x67, 0x22, 0xa7, 0x03, 0x0a, 0x19, 0x43,
	0x50, 0x32, 0x50, 0x5f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x6d, 0x6f,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x6d,
	0x6f, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x12, 0x35, 0x0a,
	0x17, 0x74, 0x76, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14,
	0x74, 0x76, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x64, 0x6f, 0x74, 0x61, 0x53, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x4d, 0x0a, 0x23, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x20, 0x64, 0x6f, 0x74, 0x61, 0x53, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f, 0x72,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x19, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x64, 0x6f, 0x74, 0x61, 0x53, 0x70, 0x65, 0x63,
	0x74, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x38,
	0x0a, 0x18, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x16, 0x64, 0x6f, 0x74, 0x61, 0x53, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x41,
	0x75, 0x74, 0x6f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x6f, 0x74, 0x61,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x2a, 0x7d, 0x0a, 0x0c, 0x50, 0x32, 0x50, 0x5f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x0f, 0x70, 0x32, 0x70, 0x5f, 0x54, 0x65, 0x78, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x80, 0x02, 0x12, 0x0e, 0x0a, 0x09, 0x70, 0x32,
	0x70, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x10, 0x81, 0x02, 0x12, 0x0d, 0x0a, 0x08, 0x70, 0x32,
	0x70, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x82, 0x02, 0x12, 0x19, 0x0a, 0x14, 0x70, 0x32, 0x70,
	0x5f, 0x56, 0x52, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0x83, 0x02, 0x12, 0x1d, 0x0a, 0x18, 0x70, 0x32, 0x70, 0x5f, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x84, 0x02, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x6f, 0x74, 0x61, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x61,
	0x2f, 0x64, 0x6f, 0x74, 0x61, 0x3b, 0x64, 0x6f, 0x74, 0x61,
}

var (
	file_c_peer2peer_netmessages_proto_rawDescOnce sync.Once
	file_c_peer2peer_netmessages_proto_rawDescData = file_c_peer2peer_netmessages_proto_rawDesc
)

func file_c_peer2peer_netmessages_proto_rawDescGZIP() []byte {
	file_c_peer2peer_netmessages_proto_rawDescOnce.Do(func() {
		file_c_peer2peer_netmessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_c_peer2peer_netmessages_proto_rawDescData)
	})
	return file_c_peer2peer_netmessages_proto_rawDescData
}

var file_c_peer2peer_netmessages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_c_peer2peer_netmessages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_c_peer2peer_netmessages_proto_goTypes = []interface{}{
	(P2P_Messages)(0),                          // 0: dota.P2P_Messages
	(CP2P_Voice_Handler_Flags)(0),              // 1: dota.CP2P_Voice.Handler_Flags
	(*CP2P_TextMessage)(nil),                   // 2: dota.CP2P_TextMessage
	(*CSteam_Voice_Encoding)(nil),              // 3: dota.CSteam_Voice_Encoding
	(*CP2P_Voice)(nil),                         // 4: dota.CP2P_Voice
	(*CP2P_Ping)(nil),                          // 5: dota.CP2P_Ping
	(*CP2P_VRAvatarPosition)(nil),              // 6: dota.CP2P_VRAvatarPosition
	(*CP2P_WatchSynchronization)(nil),          // 7: dota.CP2P_WatchSynchronization
	(*CP2P_VRAvatarPosition_COrientation)(nil), // 8: dota.CP2P_VRAvatarPosition.COrientation
	(*CMsgVoiceAudio)(nil),                     // 9: dota.CMsgVoiceAudio
	(*CMsgVector)(nil),                         // 10: dota.CMsgVector
	(*CMsgQAngle)(nil),                         // 11: dota.CMsgQAngle
}
var file_c_peer2peer_netmessages_proto_depIdxs = []int32{
	9,  // 0: dota.CP2P_Voice.audio:type_name -> dota.CMsgVoiceAudio
	8,  // 1: dota.CP2P_VRAvatarPosition.body_parts:type_name -> dota.CP2P_VRAvatarPosition.COrientation
	10, // 2: dota.CP2P_VRAvatarPosition.COrientation.pos:type_name -> dota.CMsgVector
	11, // 3: dota.CP2P_VRAvatarPosition.COrientation.ang:type_name -> dota.CMsgQAngle
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_c_peer2peer_netmessages_proto_init() }
func file_c_peer2peer_netmessages_proto_init() {
	if File_c_peer2peer_netmessages_proto != nil {
		return
	}
	file_netmessages_proto_init()
	file_networkbasetypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_c_peer2peer_netmessages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CP2P_TextMessage); i {
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
		file_c_peer2peer_netmessages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSteam_Voice_Encoding); i {
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
		file_c_peer2peer_netmessages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CP2P_Voice); i {
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
		file_c_peer2peer_netmessages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CP2P_Ping); i {
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
		file_c_peer2peer_netmessages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CP2P_VRAvatarPosition); i {
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
		file_c_peer2peer_netmessages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CP2P_WatchSynchronization); i {
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
		file_c_peer2peer_netmessages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CP2P_VRAvatarPosition_COrientation); i {
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
			RawDescriptor: file_c_peer2peer_netmessages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_c_peer2peer_netmessages_proto_goTypes,
		DependencyIndexes: file_c_peer2peer_netmessages_proto_depIdxs,
		EnumInfos:         file_c_peer2peer_netmessages_proto_enumTypes,
		MessageInfos:      file_c_peer2peer_netmessages_proto_msgTypes,
	}.Build()
	File_c_peer2peer_netmessages_proto = out.File
	file_c_peer2peer_netmessages_proto_rawDesc = nil
	file_c_peer2peer_netmessages_proto_goTypes = nil
	file_c_peer2peer_netmessages_proto_depIdxs = nil
}

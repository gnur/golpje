// Code generated by protoc-gen-go.
// source: golpje.proto
// DO NOT EDIT!

/*
Package golpje is a generated protocol buffer package.

It is generated from these files:
	golpje.proto

It has these top-level messages:
	Empty
	ProtoShow
	ProtoEvent
	ProtoEpisode
	ProtoEvents
	EventRequest
	ProtoShows
	ProtoEpisodes
	ShowRequest
	AddShowResponse
	EpisodeRequest
	AddEpisodeResponse
	SyncShowResponse
	SyncShowRequest
*/
package golpje

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProtoShow struct {
	ID       string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Regexp   string `protobuf:"bytes,3,opt,name=Regexp" json:"Regexp,omitempty"`
	Active   bool   `protobuf:"varint,4,opt,name=Active" json:"Active,omitempty"`
	Seasonal bool   `protobuf:"varint,5,opt,name=Seasonal" json:"Seasonal,omitempty"`
	Minimal  int64  `protobuf:"varint,6,opt,name=Minimal" json:"Minimal,omitempty"`
}

func (m *ProtoShow) Reset()                    { *m = ProtoShow{} }
func (m *ProtoShow) String() string            { return proto.CompactTextString(m) }
func (*ProtoShow) ProtoMessage()               {}
func (*ProtoShow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProtoShow) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ProtoShow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoShow) GetRegexp() string {
	if m != nil {
		return m.Regexp
	}
	return ""
}

func (m *ProtoShow) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *ProtoShow) GetSeasonal() bool {
	if m != nil {
		return m.Seasonal
	}
	return false
}

func (m *ProtoShow) GetMinimal() int64 {
	if m != nil {
		return m.Minimal
	}
	return 0
}

type ProtoEvent struct {
	ID        string   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Timestamp int64    `protobuf:"varint,2,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Related   []string `protobuf:"bytes,3,rep,name=related" json:"related,omitempty"`
	Data      string   `protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
}

func (m *ProtoEvent) Reset()                    { *m = ProtoEvent{} }
func (m *ProtoEvent) String() string            { return proto.CompactTextString(m) }
func (*ProtoEvent) ProtoMessage()               {}
func (*ProtoEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProtoEvent) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ProtoEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ProtoEvent) GetRelated() []string {
	if m != nil {
		return m.Related
	}
	return nil
}

func (m *ProtoEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ProtoEpisode struct {
	ID          string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=Title" json:"Title,omitempty"`
	Showid      string `protobuf:"bytes,3,opt,name=Showid" json:"Showid,omitempty"`
	Added       int64  `protobuf:"varint,4,opt,name=Added" json:"Added,omitempty"`
	Magnetlink  string `protobuf:"bytes,5,opt,name=Magnetlink" json:"Magnetlink,omitempty"`
	Episodeid   string `protobuf:"bytes,6,opt,name=Episodeid" json:"Episodeid,omitempty"`
	Downloaded  bool   `protobuf:"varint,7,opt,name=Downloaded" json:"Downloaded,omitempty"`
	Downloading bool   `protobuf:"varint,8,opt,name=Downloading" json:"Downloading,omitempty"`
}

func (m *ProtoEpisode) Reset()                    { *m = ProtoEpisode{} }
func (m *ProtoEpisode) String() string            { return proto.CompactTextString(m) }
func (*ProtoEpisode) ProtoMessage()               {}
func (*ProtoEpisode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProtoEpisode) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ProtoEpisode) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProtoEpisode) GetShowid() string {
	if m != nil {
		return m.Showid
	}
	return ""
}

func (m *ProtoEpisode) GetAdded() int64 {
	if m != nil {
		return m.Added
	}
	return 0
}

func (m *ProtoEpisode) GetMagnetlink() string {
	if m != nil {
		return m.Magnetlink
	}
	return ""
}

func (m *ProtoEpisode) GetEpisodeid() string {
	if m != nil {
		return m.Episodeid
	}
	return ""
}

func (m *ProtoEpisode) GetDownloaded() bool {
	if m != nil {
		return m.Downloaded
	}
	return false
}

func (m *ProtoEpisode) GetDownloading() bool {
	if m != nil {
		return m.Downloading
	}
	return false
}

type ProtoEvents struct {
	Events []*ProtoEvent `protobuf:"bytes,1,rep,name=Events" json:"Events,omitempty"`
}

func (m *ProtoEvents) Reset()                    { *m = ProtoEvents{} }
func (m *ProtoEvents) String() string            { return proto.CompactTextString(m) }
func (*ProtoEvents) ProtoMessage()               {}
func (*ProtoEvents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ProtoEvents) GetEvents() []*ProtoEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type EventRequest struct {
	All   bool  `protobuf:"varint,1,opt,name=All" json:"All,omitempty"`
	Since int64 `protobuf:"varint,2,opt,name=Since" json:"Since,omitempty"`
}

func (m *EventRequest) Reset()                    { *m = EventRequest{} }
func (m *EventRequest) String() string            { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()               {}
func (*EventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EventRequest) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *EventRequest) GetSince() int64 {
	if m != nil {
		return m.Since
	}
	return 0
}

type ProtoShows struct {
	Shows []*ProtoShow `protobuf:"bytes,1,rep,name=Shows" json:"Shows,omitempty"`
}

func (m *ProtoShows) Reset()                    { *m = ProtoShows{} }
func (m *ProtoShows) String() string            { return proto.CompactTextString(m) }
func (*ProtoShows) ProtoMessage()               {}
func (*ProtoShows) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ProtoShows) GetShows() []*ProtoShow {
	if m != nil {
		return m.Shows
	}
	return nil
}

type ProtoEpisodes struct {
	Episodes []*ProtoEpisode `protobuf:"bytes,1,rep,name=Episodes" json:"Episodes,omitempty"`
}

func (m *ProtoEpisodes) Reset()                    { *m = ProtoEpisodes{} }
func (m *ProtoEpisodes) String() string            { return proto.CompactTextString(m) }
func (*ProtoEpisodes) ProtoMessage()               {}
func (*ProtoEpisodes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ProtoEpisodes) GetEpisodes() []*ProtoEpisode {
	if m != nil {
		return m.Episodes
	}
	return nil
}

type ShowRequest struct {
	Onlyactive bool   `protobuf:"varint,1,opt,name=Onlyactive" json:"Onlyactive,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *ShowRequest) Reset()                    { *m = ShowRequest{} }
func (m *ShowRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()               {}
func (*ShowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ShowRequest) GetOnlyactive() bool {
	if m != nil {
		return m.Onlyactive
	}
	return false
}

func (m *ShowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddShowResponse struct {
	Show    *ProtoShow `protobuf:"bytes,1,opt,name=Show" json:"Show,omitempty"`
	Success bool       `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	Error   string     `protobuf:"bytes,3,opt,name=Error" json:"Error,omitempty"`
}

func (m *AddShowResponse) Reset()                    { *m = AddShowResponse{} }
func (m *AddShowResponse) String() string            { return proto.CompactTextString(m) }
func (*AddShowResponse) ProtoMessage()               {}
func (*AddShowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddShowResponse) GetShow() *ProtoShow {
	if m != nil {
		return m.Show
	}
	return nil
}

func (m *AddShowResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AddShowResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EpisodeRequest struct {
	Showid string `protobuf:"bytes,1,opt,name=Showid" json:"Showid,omitempty"`
	Season uint32 `protobuf:"varint,2,opt,name=Season" json:"Season,omitempty"`
}

func (m *EpisodeRequest) Reset()                    { *m = EpisodeRequest{} }
func (m *EpisodeRequest) String() string            { return proto.CompactTextString(m) }
func (*EpisodeRequest) ProtoMessage()               {}
func (*EpisodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *EpisodeRequest) GetShowid() string {
	if m != nil {
		return m.Showid
	}
	return ""
}

func (m *EpisodeRequest) GetSeason() uint32 {
	if m != nil {
		return m.Season
	}
	return 0
}

type AddEpisodeResponse struct {
	Episode *ProtoEpisode `protobuf:"bytes,1,opt,name=Episode" json:"Episode,omitempty"`
	Success bool          `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	Error   string        `protobuf:"bytes,3,opt,name=Error" json:"Error,omitempty"`
}

func (m *AddEpisodeResponse) Reset()                    { *m = AddEpisodeResponse{} }
func (m *AddEpisodeResponse) String() string            { return proto.CompactTextString(m) }
func (*AddEpisodeResponse) ProtoMessage()               {}
func (*AddEpisodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AddEpisodeResponse) GetEpisode() *ProtoEpisode {
	if m != nil {
		return m.Episode
	}
	return nil
}

func (m *AddEpisodeResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AddEpisodeResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SyncShowResponse struct {
	FoundEpisodes int64  `protobuf:"varint,1,opt,name=FoundEpisodes" json:"FoundEpisodes,omitempty"`
	Success       bool   `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	Error         string `protobuf:"bytes,3,opt,name=Error" json:"Error,omitempty"`
}

func (m *SyncShowResponse) Reset()                    { *m = SyncShowResponse{} }
func (m *SyncShowResponse) String() string            { return proto.CompactTextString(m) }
func (*SyncShowResponse) ProtoMessage()               {}
func (*SyncShowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SyncShowResponse) GetFoundEpisodes() int64 {
	if m != nil {
		return m.FoundEpisodes
	}
	return 0
}

func (m *SyncShowResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SyncShowResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SyncShowRequest struct {
	ShowID string `protobuf:"bytes,1,opt,name=ShowID" json:"ShowID,omitempty"`
}

func (m *SyncShowRequest) Reset()                    { *m = SyncShowRequest{} }
func (m *SyncShowRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncShowRequest) ProtoMessage()               {}
func (*SyncShowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SyncShowRequest) GetShowID() string {
	if m != nil {
		return m.ShowID
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "golpje.Empty")
	proto.RegisterType((*ProtoShow)(nil), "golpje.ProtoShow")
	proto.RegisterType((*ProtoEvent)(nil), "golpje.ProtoEvent")
	proto.RegisterType((*ProtoEpisode)(nil), "golpje.ProtoEpisode")
	proto.RegisterType((*ProtoEvents)(nil), "golpje.ProtoEvents")
	proto.RegisterType((*EventRequest)(nil), "golpje.EventRequest")
	proto.RegisterType((*ProtoShows)(nil), "golpje.ProtoShows")
	proto.RegisterType((*ProtoEpisodes)(nil), "golpje.ProtoEpisodes")
	proto.RegisterType((*ShowRequest)(nil), "golpje.ShowRequest")
	proto.RegisterType((*AddShowResponse)(nil), "golpje.AddShowResponse")
	proto.RegisterType((*EpisodeRequest)(nil), "golpje.EpisodeRequest")
	proto.RegisterType((*AddEpisodeResponse)(nil), "golpje.AddEpisodeResponse")
	proto.RegisterType((*SyncShowResponse)(nil), "golpje.SyncShowResponse")
	proto.RegisterType((*SyncShowRequest)(nil), "golpje.SyncShowRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Golpje service

type GolpjeClient interface {
	GetEvents(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*ProtoEvents, error)
	GetShows(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ProtoShows, error)
	AddShow(ctx context.Context, in *ProtoShow, opts ...grpc.CallOption) (*AddShowResponse, error)
	DelShow(ctx context.Context, in *ProtoShow, opts ...grpc.CallOption) (*AddShowResponse, error)
	SyncShow(ctx context.Context, in *SyncShowRequest, opts ...grpc.CallOption) (*SyncShowResponse, error)
	GetEpisodes(ctx context.Context, in *EpisodeRequest, opts ...grpc.CallOption) (*ProtoEpisodes, error)
	AddEpisode(ctx context.Context, in *ProtoEpisode, opts ...grpc.CallOption) (*AddEpisodeResponse, error)
}

type golpjeClient struct {
	cc *grpc.ClientConn
}

func NewGolpjeClient(cc *grpc.ClientConn) GolpjeClient {
	return &golpjeClient{cc}
}

func (c *golpjeClient) GetEvents(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*ProtoEvents, error) {
	out := new(ProtoEvents)
	err := grpc.Invoke(ctx, "/golpje.Golpje/GetEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golpjeClient) GetShows(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ProtoShows, error) {
	out := new(ProtoShows)
	err := grpc.Invoke(ctx, "/golpje.Golpje/GetShows", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golpjeClient) AddShow(ctx context.Context, in *ProtoShow, opts ...grpc.CallOption) (*AddShowResponse, error) {
	out := new(AddShowResponse)
	err := grpc.Invoke(ctx, "/golpje.Golpje/AddShow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golpjeClient) DelShow(ctx context.Context, in *ProtoShow, opts ...grpc.CallOption) (*AddShowResponse, error) {
	out := new(AddShowResponse)
	err := grpc.Invoke(ctx, "/golpje.Golpje/DelShow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golpjeClient) SyncShow(ctx context.Context, in *SyncShowRequest, opts ...grpc.CallOption) (*SyncShowResponse, error) {
	out := new(SyncShowResponse)
	err := grpc.Invoke(ctx, "/golpje.Golpje/SyncShow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golpjeClient) GetEpisodes(ctx context.Context, in *EpisodeRequest, opts ...grpc.CallOption) (*ProtoEpisodes, error) {
	out := new(ProtoEpisodes)
	err := grpc.Invoke(ctx, "/golpje.Golpje/GetEpisodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *golpjeClient) AddEpisode(ctx context.Context, in *ProtoEpisode, opts ...grpc.CallOption) (*AddEpisodeResponse, error) {
	out := new(AddEpisodeResponse)
	err := grpc.Invoke(ctx, "/golpje.Golpje/AddEpisode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Golpje service

type GolpjeServer interface {
	GetEvents(context.Context, *EventRequest) (*ProtoEvents, error)
	GetShows(context.Context, *ShowRequest) (*ProtoShows, error)
	AddShow(context.Context, *ProtoShow) (*AddShowResponse, error)
	DelShow(context.Context, *ProtoShow) (*AddShowResponse, error)
	SyncShow(context.Context, *SyncShowRequest) (*SyncShowResponse, error)
	GetEpisodes(context.Context, *EpisodeRequest) (*ProtoEpisodes, error)
	AddEpisode(context.Context, *ProtoEpisode) (*AddEpisodeResponse, error)
}

func RegisterGolpjeServer(s *grpc.Server, srv GolpjeServer) {
	s.RegisterService(&_Golpje_serviceDesc, srv)
}

func _Golpje_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolpjeServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golpje.Golpje/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolpjeServer).GetEvents(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Golpje_GetShows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolpjeServer).GetShows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golpje.Golpje/GetShows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolpjeServer).GetShows(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Golpje_AddShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtoShow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolpjeServer).AddShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golpje.Golpje/AddShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolpjeServer).AddShow(ctx, req.(*ProtoShow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Golpje_DelShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtoShow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolpjeServer).DelShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golpje.Golpje/DelShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolpjeServer).DelShow(ctx, req.(*ProtoShow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Golpje_SyncShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolpjeServer).SyncShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golpje.Golpje/SyncShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolpjeServer).SyncShow(ctx, req.(*SyncShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Golpje_GetEpisodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolpjeServer).GetEpisodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golpje.Golpje/GetEpisodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolpjeServer).GetEpisodes(ctx, req.(*EpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Golpje_AddEpisode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtoEpisode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GolpjeServer).AddEpisode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golpje.Golpje/AddEpisode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GolpjeServer).AddEpisode(ctx, req.(*ProtoEpisode))
	}
	return interceptor(ctx, in, info, handler)
}

var _Golpje_serviceDesc = grpc.ServiceDesc{
	ServiceName: "golpje.Golpje",
	HandlerType: (*GolpjeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _Golpje_GetEvents_Handler,
		},
		{
			MethodName: "GetShows",
			Handler:    _Golpje_GetShows_Handler,
		},
		{
			MethodName: "AddShow",
			Handler:    _Golpje_AddShow_Handler,
		},
		{
			MethodName: "DelShow",
			Handler:    _Golpje_DelShow_Handler,
		},
		{
			MethodName: "SyncShow",
			Handler:    _Golpje_SyncShow_Handler,
		},
		{
			MethodName: "GetEpisodes",
			Handler:    _Golpje_GetEpisodes_Handler,
		},
		{
			MethodName: "AddEpisode",
			Handler:    _Golpje_AddEpisode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golpje.proto",
}

func init() { proto.RegisterFile("golpje.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5b, 0x6e, 0xd3, 0x40,
	0x14, 0xad, 0xeb, 0x26, 0xb1, 0x6f, 0xfa, 0x62, 0x28, 0xc5, 0xb2, 0x50, 0x15, 0x8d, 0x40, 0x04,
	0x3e, 0x2a, 0x54, 0x54, 0x1e, 0x3f, 0x50, 0x4b, 0x29, 0x55, 0x3f, 0x0a, 0x68, 0xd2, 0x0d, 0x98,
	0xcc, 0x28, 0x31, 0x38, 0x63, 0x93, 0x99, 0xb6, 0x74, 0x21, 0x6c, 0x83, 0x45, 0xb1, 0x12, 0x34,
	0x2f, 0x3f, 0x9a, 0xfc, 0x94, 0xbf, 0x39, 0xe7, 0xde, 0xb9, 0x8f, 0x33, 0xc7, 0x32, 0x6c, 0x4e,
	0x8b, 0xbc, 0xfc, 0xce, 0x0e, 0xcb, 0x45, 0x21, 0x0b, 0xd4, 0x35, 0x08, 0xf7, 0xa0, 0x73, 0x3a,
	0x2f, 0xe5, 0x2d, 0xfe, 0xed, 0x41, 0xf8, 0x55, 0x85, 0xc6, 0xb3, 0xe2, 0x06, 0x6d, 0xc3, 0xfa,
	0xf9, 0x28, 0xf2, 0x06, 0xde, 0x30, 0x24, 0xeb, 0xe7, 0x23, 0x84, 0x60, 0xe3, 0x73, 0x3a, 0x67,
	0xd1, 0xba, 0x66, 0xf4, 0x19, 0xed, 0x43, 0x97, 0xb0, 0x29, 0xfb, 0x55, 0x46, 0xbe, 0x66, 0x2d,
	0x52, 0x7c, 0x32, 0x91, 0xd9, 0x35, 0x8b, 0x36, 0x06, 0xde, 0x30, 0x20, 0x16, 0xa1, 0x18, 0x82,
	0x31, 0x4b, 0x45, 0xc1, 0xd3, 0x3c, 0xea, 0xe8, 0x48, 0x85, 0x51, 0x04, 0xbd, 0x8b, 0x8c, 0x67,
	0xf3, 0x34, 0x8f, 0xba, 0x03, 0x6f, 0xe8, 0x13, 0x07, 0xf1, 0x0c, 0x40, 0x8f, 0x75, 0x7a, 0xcd,
	0xb8, 0x5c, 0x9a, 0xeb, 0x09, 0x84, 0x97, 0xd9, 0x9c, 0x09, 0x99, 0xce, 0x4b, 0x3d, 0x9c, 0x4f,
	0x6a, 0x42, 0x55, 0x5d, 0xb0, 0x3c, 0x95, 0x8c, 0x46, 0xfe, 0xc0, 0x1f, 0x86, 0xc4, 0x41, 0xb5,
	0xcf, 0x28, 0x95, 0xa9, 0x9e, 0x30, 0x24, 0xfa, 0x8c, 0xff, 0x7a, 0xb0, 0x69, 0x5a, 0x95, 0x99,
	0x28, 0x28, 0x5b, 0x6a, 0xb6, 0x07, 0x9d, 0xcb, 0x4c, 0xe6, 0x4e, 0x05, 0x03, 0xd4, 0xba, 0x4a,
	0xb2, 0x8c, 0x3a, 0x19, 0x0c, 0x52, 0xd9, 0x09, 0xa5, 0x8c, 0xea, 0x1e, 0x3e, 0x31, 0x00, 0x1d,
	0x00, 0x5c, 0xa4, 0x53, 0xce, 0x64, 0x9e, 0xf1, 0x1f, 0x5a, 0x86, 0x90, 0x34, 0x18, 0xb5, 0x90,
	0x6d, 0x9f, 0x51, 0x2d, 0x45, 0x48, 0x6a, 0x42, 0xdd, 0x1e, 0x15, 0x37, 0x3c, 0x2f, 0x52, 0x55,
	0xb8, 0xa7, 0x45, 0x6c, 0x30, 0x68, 0x00, 0x7d, 0x87, 0x32, 0x3e, 0x8d, 0x02, 0x9d, 0xd0, 0xa4,
	0xf0, 0x7b, 0xe8, 0xd7, 0x72, 0x0a, 0xf4, 0x12, 0xba, 0xe6, 0x14, 0x79, 0x03, 0x7f, 0xd8, 0x3f,
	0x42, 0x87, 0xd6, 0x25, 0x75, 0x12, 0xb1, 0x19, 0xf8, 0x0d, 0x6c, 0x1a, 0x82, 0xfd, 0xbc, 0x62,
	0x42, 0xa2, 0x5d, 0xf0, 0x93, 0x3c, 0xd7, 0xfa, 0x04, 0x44, 0x1d, 0xd5, 0xca, 0xe3, 0x8c, 0x4f,
	0x98, 0x7d, 0x09, 0x03, 0xf0, 0xb1, 0x7d, 0x41, 0xa5, 0x8b, 0x40, 0xcf, 0xa1, 0xa3, 0x0f, 0xb6,
	0xe1, 0x83, 0x56, 0x43, 0x15, 0x21, 0x26, 0x8e, 0x13, 0xd8, 0x6a, 0xbe, 0x86, 0x40, 0xaf, 0x20,
	0x70, 0x67, 0x7b, 0x79, 0xaf, 0x3d, 0xad, 0x09, 0x92, 0x2a, 0x0b, 0x27, 0xd0, 0xd7, 0x15, 0xed,
	0xc0, 0x07, 0x00, 0x5f, 0x78, 0x7e, 0x9b, 0x1a, 0x73, 0x9a, 0xb9, 0x1b, 0xcc, 0x2a, 0x93, 0xe3,
	0x19, 0xec, 0x24, 0x94, 0x9a, 0x2a, 0xa2, 0x2c, 0xb8, 0x60, 0xe8, 0x19, 0x6c, 0x28, 0xac, 0x0b,
	0xac, 0x5c, 0x40, 0x87, 0x95, 0xf9, 0xc6, 0x57, 0x93, 0x09, 0x13, 0x42, 0x17, 0x0c, 0x88, 0x83,
	0x4a, 0xa6, 0xd3, 0xc5, 0xa2, 0x58, 0x58, 0xc3, 0x18, 0x80, 0x4f, 0x60, 0xdb, 0x6d, 0x60, 0xe7,
	0xad, 0x9d, 0xe5, 0xb5, 0x9c, 0xa5, 0x78, 0xfd, 0xe1, 0xe8, 0xc2, 0x5b, 0xc4, 0x22, 0x2c, 0x01,
	0x25, 0x94, 0x56, 0x45, 0xec, 0xb8, 0x87, 0xd0, 0xb3, 0x94, 0x9d, 0x78, 0xb5, 0x6a, 0x2e, 0xe9,
	0xde, 0x73, 0xcf, 0x60, 0x77, 0x7c, 0xcb, 0x27, 0x2d, 0x89, 0x9e, 0xc2, 0xd6, 0xa7, 0xe2, 0x8a,
	0xd3, 0xc6, 0x7b, 0x29, 0x43, 0xb4, 0xc9, 0x7b, 0x77, 0x7a, 0x01, 0x3b, 0x75, 0xa7, 0x96, 0x44,
	0xd5, 0x67, 0x6a, 0xd1, 0xd1, 0x1f, 0x1f, 0xba, 0x67, 0x7a, 0x4b, 0xf4, 0x0e, 0xc2, 0x33, 0x26,
	0xad, 0xdf, 0xab, 0xdd, 0x9b, 0x4e, 0x8e, 0x1f, 0x2e, 0xbb, 0x5e, 0xe0, 0x35, 0x74, 0x0c, 0xc1,
	0x19, 0x93, 0xc6, 0xb6, 0x55, 0x4a, 0xa3, 0x7b, 0x8c, 0x96, 0xde, 0x5e, 0x5d, 0x7b, 0x0b, 0x3d,
	0x6b, 0x19, 0xb4, 0x6c, 0x8e, 0xf8, 0xb1, 0xa3, 0xee, 0xd8, 0xca, 0x5c, 0x1c, 0xb1, 0xfc, 0x3f,
	0x2e, 0x7e, 0x84, 0xc0, 0x09, 0x83, 0xaa, 0xb4, 0x3b, 0x52, 0xc5, 0xd1, 0x72, 0xa0, 0x2a, 0xf0,
	0x01, 0xfa, 0x4a, 0x23, 0xf7, 0x30, 0xfb, 0x95, 0x4a, 0x2d, 0x43, 0xc6, 0x8f, 0x56, 0x39, 0x47,
	0xad, 0x7c, 0x02, 0x50, 0x3b, 0x0f, 0xad, 0x34, 0x58, 0x1c, 0x37, 0xe6, 0xbf, 0xe3, 0x51, 0xbc,
	0xf6, 0xad, 0xab, 0x7f, 0x4b, 0xaf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x27, 0xc4, 0x64,
	0xa6, 0x06, 0x00, 0x00,
}

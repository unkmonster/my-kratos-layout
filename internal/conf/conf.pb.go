// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.0--dev
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Env int32

const (
	Env_UNSPECIFIED Env = 0
	Env_DEV         Env = 1
	Env_STAGING     Env = 2
	Env_PROD        Env = 3
)

// Enum value maps for Env.
var (
	Env_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "DEV",
		2: "STAGING",
		3: "PROD",
	}
	Env_value = map[string]int32{
		"UNSPECIFIED": 0,
		"DEV":         1,
		"STAGING":     2,
		"PROD":        3,
	}
)

func (x Env) Enum() *Env {
	p := new(Env)
	*p = x
	return p
}

func (x Env) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Env) Descriptor() protoreflect.EnumDescriptor {
	return file_conf_conf_proto_enumTypes[0].Descriptor()
}

func (Env) Type() protoreflect.EnumType {
	return &file_conf_conf_proto_enumTypes[0]
}

func (x Env) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Env.Descriptor instead.
func (Env) EnumDescriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

// TODO 服务注册 & 发现
type Bootstrap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Env           Env                    `protobuf:"varint,1,opt,name=env,proto3,enum=kratos.api.Env" json:"env,omitempty"`
	Server        *Server                `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Data          *Data                  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Observability *Observability         `protobuf:"bytes,4,opt,name=observability,proto3" json:"observability,omitempty"`
	Client        *Client                `protobuf:"bytes,5,opt,name=client,proto3" json:"client,omitempty"`
	App           *App                   `protobuf:"bytes,15,opt,name=app,proto3" json:"app,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_conf_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetEnv() Env {
	if x != nil {
		return x.Env
	}
	return Env_UNSPECIFIED
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetObservability() *Observability {
	if x != nil {
		return x.Observability
	}
	return nil
}

func (x *Bootstrap) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Bootstrap) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_conf_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis         *Data_Redis            `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_conf_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

type Observability struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Trace         *Observability_Trace   `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	Metrics       *Observability_Metrics `protobuf:"bytes,2,opt,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Observability) Reset() {
	*x = Observability{}
	mi := &file_conf_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Observability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observability) ProtoMessage() {}

func (x *Observability) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observability.ProtoReflect.Descriptor instead.
func (*Observability) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Observability) GetTrace() *Observability_Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Observability) GetMetrics() *Observability_Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Client_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Client) Reset() {
	*x = Client{}
	mi := &file_conf_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Client) GetHttp() *Client_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

type App struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *App) Reset() {
	*x = App{}
	mi := &file_conf_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{5}
}

type Server_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_conf_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_conf_conf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Protocol      string                 `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Schema        string                 `protobuf:"bytes,6,opt,name=schema,proto3" json:"schema,omitempty"`
	Params        string                 `protobuf:"bytes,7,opt,name=params,proto3" json:"params,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	mi := &file_conf_conf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Database) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Data_Database) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Database) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Data_Database) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Data_Database) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout   *durationpb.Duration   `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout  *durationpb.Duration   `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	Username      string                 `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Db            int32                  `protobuf:"varint,7,opt,name=db,proto3" json:"db,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	mi := &file_conf_conf_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Data_Redis) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

type Observability_Trace struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// host:port grpc endpoint
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// 生产环境采样率，仅生产环境
	ProductionSampleRate float32 `protobuf:"fixed32,2,opt,name=production_sample_rate,json=productionSampleRate,proto3" json:"production_sample_rate,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Observability_Trace) Reset() {
	*x = Observability_Trace{}
	mi := &file_conf_conf_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Observability_Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observability_Trace) ProtoMessage() {}

func (x *Observability_Trace) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observability_Trace.ProtoReflect.Descriptor instead.
func (*Observability_Trace) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Observability_Trace) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Observability_Trace) GetProductionSampleRate() float32 {
	if x != nil {
		return x.ProductionSampleRate
	}
	return 0
}

type Observability_Metrics struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// grpc endpoint
	Endpoint      string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Observability_Metrics) Reset() {
	*x = Observability_Metrics{}
	mi := &file_conf_conf_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Observability_Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observability_Metrics) ProtoMessage() {}

func (x *Observability_Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observability_Metrics.ProtoReflect.Descriptor instead.
func (*Observability_Metrics) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Observability_Metrics) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type Client_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RetryCount    int32                  `protobuf:"varint,1,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Client_HTTP) Reset() {
	*x = Client_HTTP{}
	mi := &file_conf_conf_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Client_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_HTTP) ProtoMessage() {}

func (x *Client_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_HTTP.ProtoReflect.Descriptor instead.
func (*Client_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Client_HTTP) GetRetryCount() int32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

var File_conf_conf_proto protoreflect.FileDescriptor

const file_conf_conf_proto_rawDesc = "" +
	"\n" +
	"\x0fconf/conf.proto\x12\n" +
	"kratos.api\x1a\x1egoogle/protobuf/duration.proto\"\x90\x02\n" +
	"\tBootstrap\x12!\n" +
	"\x03env\x18\x01 \x01(\x0e2\x0f.kratos.api.EnvR\x03env\x12*\n" +
	"\x06server\x18\x02 \x01(\v2\x12.kratos.api.ServerR\x06server\x12$\n" +
	"\x04data\x18\x03 \x01(\v2\x10.kratos.api.DataR\x04data\x12?\n" +
	"\robservability\x18\x04 \x01(\v2\x19.kratos.api.ObservabilityR\robservability\x12*\n" +
	"\x06client\x18\x05 \x01(\v2\x12.kratos.api.ClientR\x06client\x12!\n" +
	"\x03app\x18\x0f \x01(\v2\x0f.kratos.api.AppR\x03app\"\xb8\x02\n" +
	"\x06Server\x12+\n" +
	"\x04http\x18\x01 \x01(\v2\x17.kratos.api.Server.HTTPR\x04http\x12+\n" +
	"\x04grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\x04grpc\x1ai\n" +
	"\x04HTTP\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\x1ai\n" +
	"\x04GRPC\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x123\n" +
	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\"\xa6\x04\n" +
	"\x04Data\x125\n" +
	"\bdatabase\x18\x01 \x01(\v2\x19.kratos.api.Data.DatabaseR\bdatabase\x12,\n" +
	"\x05redis\x18\x02 \x01(\v2\x16.kratos.api.Data.RedisR\x05redis\x1a\xba\x01\n" +
	"\bDatabase\x12\x16\n" +
	"\x06driver\x18\x01 \x01(\tR\x06driver\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x12\x1a\n" +
	"\busername\x18\x03 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\x12\x1a\n" +
	"\bprotocol\x18\x05 \x01(\tR\bprotocol\x12\x16\n" +
	"\x06schema\x18\x06 \x01(\tR\x06schema\x12\x16\n" +
	"\x06params\x18\a \x01(\tR\x06params\x1a\xfb\x01\n" +
	"\x05Redis\x12\x18\n" +
	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x12\n" +
	"\x04addr\x18\x02 \x01(\tR\x04addr\x12<\n" +
	"\fread_timeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\vreadTimeout\x12>\n" +
	"\rwrite_timeout\x18\x04 \x01(\v2\x19.google.protobuf.DurationR\fwriteTimeout\x12\x1a\n" +
	"\busername\x18\x05 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x06 \x01(\tR\bpassword\x12\x0e\n" +
	"\x02db\x18\a \x01(\x05R\x02db\"\x85\x02\n" +
	"\rObservability\x125\n" +
	"\x05trace\x18\x01 \x01(\v2\x1f.kratos.api.Observability.TraceR\x05trace\x12;\n" +
	"\ametrics\x18\x02 \x01(\v2!.kratos.api.Observability.MetricsR\ametrics\x1aY\n" +
	"\x05Trace\x12\x1a\n" +
	"\bendpoint\x18\x01 \x01(\tR\bendpoint\x124\n" +
	"\x16production_sample_rate\x18\x02 \x01(\x02R\x14productionSampleRate\x1a%\n" +
	"\aMetrics\x12\x1a\n" +
	"\bendpoint\x18\x01 \x01(\tR\bendpoint\"^\n" +
	"\x06Client\x12+\n" +
	"\x04http\x18\x01 \x01(\v2\x17.kratos.api.Client.HTTPR\x04http\x1a'\n" +
	"\x04HTTP\x12\x1f\n" +
	"\vretry_count\x18\x01 \x01(\x05R\n" +
	"retryCount\"\x05\n" +
	"\x03App*6\n" +
	"\x03Env\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\a\n" +
	"\x03DEV\x10\x01\x12\v\n" +
	"\aSTAGING\x10\x02\x12\b\n" +
	"\x04PROD\x10\x03B7Z5github.com/go-kratos/kratos-layout/internal/conf;confb\x06proto3"

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData []byte
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)))
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_conf_conf_proto_goTypes = []any{
	(Env)(0),                      // 0: kratos.api.Env
	(*Bootstrap)(nil),             // 1: kratos.api.Bootstrap
	(*Server)(nil),                // 2: kratos.api.Server
	(*Data)(nil),                  // 3: kratos.api.Data
	(*Observability)(nil),         // 4: kratos.api.Observability
	(*Client)(nil),                // 5: kratos.api.Client
	(*App)(nil),                   // 6: kratos.api.App
	(*Server_HTTP)(nil),           // 7: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),           // 8: kratos.api.Server.GRPC
	(*Data_Database)(nil),         // 9: kratos.api.Data.Database
	(*Data_Redis)(nil),            // 10: kratos.api.Data.Redis
	(*Observability_Trace)(nil),   // 11: kratos.api.Observability.Trace
	(*Observability_Metrics)(nil), // 12: kratos.api.Observability.Metrics
	(*Client_HTTP)(nil),           // 13: kratos.api.Client.HTTP
	(*durationpb.Duration)(nil),   // 14: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	0,  // 0: kratos.api.Bootstrap.env:type_name -> kratos.api.Env
	2,  // 1: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	3,  // 2: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	4,  // 3: kratos.api.Bootstrap.observability:type_name -> kratos.api.Observability
	5,  // 4: kratos.api.Bootstrap.client:type_name -> kratos.api.Client
	6,  // 5: kratos.api.Bootstrap.app:type_name -> kratos.api.App
	7,  // 6: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	8,  // 7: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	9,  // 8: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	10, // 9: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	11, // 10: kratos.api.Observability.trace:type_name -> kratos.api.Observability.Trace
	12, // 11: kratos.api.Observability.metrics:type_name -> kratos.api.Observability.Metrics
	13, // 12: kratos.api.Client.http:type_name -> kratos.api.Client.HTTP
	14, // 13: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	14, // 14: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	14, // 15: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	14, // 16: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		EnumInfos:         file_conf_conf_proto_enumTypes,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}

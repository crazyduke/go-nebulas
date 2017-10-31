// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package nebletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	P2PConfig
	RPCConfig
	PowConfig
	AccountConfig
*/
package nebletpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Neblet global configurations.
type Config struct {
	// P2P network config.
	P2P *P2PConfig `protobuf:"bytes,1,opt,name=p2p" json:"p2p,omitempty"`
	// RPC server config.
	Rpc *RPCConfig `protobuf:"bytes,2,opt,name=rpc" json:"rpc,omitempty"`
	// Pow config
	Pow *PowConfig `protobuf:"bytes,3,opt,name=pow" json:"pow,omitempty"`
	// Account manager config
	Account *AccountConfig `protobuf:"bytes,4,opt,name=account" json:"account,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetP2P() *P2PConfig {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetPow() *PowConfig {
	if m != nil {
		return m.Pow
	}
	return nil
}

func (m *Config) GetAccount() *AccountConfig {
	if m != nil {
		return m.Account
	}
	return nil
}

type P2PConfig struct {
	// P2P seed node addresses.
	Seed string `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	// P2P node port number.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// P2P node chainID.
	ChainId uint32 `protobuf:"varint,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// P2P node version.
	Version uint32 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *P2PConfig) Reset()                    { *m = P2PConfig{} }
func (m *P2PConfig) String() string            { return proto.CompactTextString(m) }
func (*P2PConfig) ProtoMessage()               {}
func (*P2PConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *P2PConfig) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *P2PConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *P2PConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *P2PConfig) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type RPCConfig struct {
	// API RPC server listening port number.
	ApiPort uint32 `protobuf:"varint,1,opt,name=api_port,json=apiPort,proto3" json:"api_port,omitempty"`
	// Management RPC server listening port number.
	ManagementPort uint32 `protobuf:"varint,2,opt,name=management_port,json=managementPort,proto3" json:"management_port,omitempty"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *RPCConfig) GetApiPort() uint32 {
	if m != nil {
		return m.ApiPort
	}
	return 0
}

func (m *RPCConfig) GetManagementPort() uint32 {
	if m != nil {
		return m.ManagementPort
	}
	return 0
}

type PowConfig struct {
	// pow miner's coinbase
	Coinbase string `protobuf:"bytes,1,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
}

func (m *PowConfig) Reset()                    { *m = PowConfig{} }
func (m *PowConfig) String() string            { return proto.CompactTextString(m) }
func (*PowConfig) ProtoMessage()               {}
func (*PowConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *PowConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

type AccountConfig struct {
	// Account transaction sign algorithm type
	Signature uint32 `protobuf:"varint,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Account addr encrypt algorithm type
	Encrypt uint32 `protobuf:"varint,2,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	// Account key file directory
	KeyDir string `protobuf:"bytes,3,opt,name=key_dir,json=keyDir,proto3" json:"key_dir,omitempty"`
	// Account test keyfile's passphrase
	TestPassphrase string `protobuf:"bytes,4,opt,name=test_passphrase,json=testPassphrase,proto3" json:"test_passphrase,omitempty"`
}

func (m *AccountConfig) Reset()                    { *m = AccountConfig{} }
func (m *AccountConfig) String() string            { return proto.CompactTextString(m) }
func (*AccountConfig) ProtoMessage()               {}
func (*AccountConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *AccountConfig) GetSignature() uint32 {
	if m != nil {
		return m.Signature
	}
	return 0
}

func (m *AccountConfig) GetEncrypt() uint32 {
	if m != nil {
		return m.Encrypt
	}
	return 0
}

func (m *AccountConfig) GetKeyDir() string {
	if m != nil {
		return m.KeyDir
	}
	return ""
}

func (m *AccountConfig) GetTestPassphrase() string {
	if m != nil {
		return m.TestPassphrase
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "nebletpb.Config")
	proto.RegisterType((*P2PConfig)(nil), "nebletpb.P2PConfig")
	proto.RegisterType((*RPCConfig)(nil), "nebletpb.RPCConfig")
	proto.RegisterType((*PowConfig)(nil), "nebletpb.PowConfig")
	proto.RegisterType((*AccountConfig)(nil), "nebletpb.AccountConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xf1, 0x26, 0xc4, 0xf1, 0xec, 0x3a, 0x0b, 0xda, 0x43, 0xbc, 0xa5, 0x87, 0x62, 0x28,
	0xe9, 0x29, 0xd0, 0xf4, 0x09, 0x4a, 0x7a, 0xe9, 0xa9, 0x46, 0x2f, 0x60, 0x64, 0x79, 0x9a, 0x88,
	0x34, 0x92, 0x90, 0x94, 0x86, 0x3c, 0x42, 0xdf, 0xa6, 0x8f, 0x58, 0x34, 0x89, 0x9d, 0xfe, 0xb9,
	0x69, 0xbe, 0xf9, 0xf1, 0xcd, 0xf7, 0x81, 0xe0, 0x8f, 0x34, 0xfa, 0x59, 0xad, 0xe6, 0xd6, 0x99,
	0x60, 0xd8, 0x58, 0x63, 0xf3, 0x82, 0xc1, 0x36, 0xe5, 0x7b, 0x02, 0xa3, 0x25, 0xad, 0xd8, 0x35,
	0x0c, 0xec, 0xc2, 0x16, 0xc9, 0x55, 0x72, 0xf3, 0x7b, 0xf1, 0x6f, 0xde, 0x21, 0xf3, 0x6a, 0x51,
	0x1d, 0x09, 0x1e, 0xf7, 0x11, 0x73, 0x56, 0x16, 0xbf, 0xbe, 0x63, 0xbc, 0x5a, 0x76, 0x98, 0xb3,
	0x92, 0xdc, 0xcc, 0xbe, 0x18, 0xfc, 0x70, 0x33, 0xfb, 0xde, 0xcd, 0xec, 0xd9, 0x2d, 0xa4, 0x42,
	0x4a, 0xb3, 0xd3, 0xa1, 0x18, 0x12, 0x3a, 0x3d, 0xa3, 0xf7, 0xc7, 0xc5, 0x09, 0xef, 0xb8, 0x72,
	0x0d, 0x59, 0x1f, 0x89, 0x31, 0x18, 0x7a, 0xc4, 0x96, 0x52, 0x67, 0x9c, 0xde, 0x51, 0xb3, 0xc6,
	0x05, 0x8a, 0x98, 0x73, 0x7a, 0xb3, 0xff, 0x30, 0x96, 0x6b, 0xa1, 0x74, 0xad, 0x5a, 0xca, 0x94,
	0xf3, 0x94, 0xe6, 0xc7, 0x96, 0x15, 0x90, 0xbe, 0xa2, 0xf3, 0xca, 0x68, 0x8a, 0x90, 0xf3, 0x6e,
	0x2c, 0x9f, 0x20, 0xeb, 0x5b, 0x45, 0x07, 0x61, 0x55, 0x4d, 0xce, 0xc9, 0x91, 0x13, 0x56, 0x55,
	0xd1, 0x7c, 0x06, 0x7f, 0xb7, 0x42, 0x8b, 0x15, 0x6e, 0x51, 0x87, 0xfa, 0xd3, 0xed, 0xc9, 0x59,
	0x8e, 0x60, 0x39, 0x83, 0xac, 0xef, 0xcf, 0x2e, 0x60, 0x2c, 0x8d, 0xd2, 0x8d, 0xf0, 0x78, 0x8a,
	0xdf, 0xcf, 0xe5, 0x5b, 0x02, 0xf9, 0x97, 0xfa, 0xec, 0x12, 0x32, 0xaf, 0x56, 0x5a, 0x84, 0x9d,
	0xc3, 0xd3, 0xfd, 0xb3, 0x10, 0x3b, 0xa0, 0x96, 0xee, 0x60, 0xbb, 0xcb, 0xdd, 0xc8, 0xa6, 0x90,
	0x6e, 0xf0, 0x50, 0xb7, 0xca, 0x51, 0xef, 0x8c, 0x8f, 0x36, 0x78, 0x78, 0x50, 0x2e, 0x86, 0x0e,
	0xe8, 0x43, 0x6d, 0x85, 0xf7, 0x76, 0xed, 0x62, 0x8a, 0x21, 0x01, 0x93, 0x28, 0x57, 0xbd, 0xda,
	0x8c, 0xe8, 0xcf, 0xdc, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x30, 0x17, 0xf3, 0x6d, 0x43, 0x02,
	0x00, 0x00,
}

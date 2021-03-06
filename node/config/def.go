package config

import (
	"encoding"
	"time"
)

// Common is common config between full node and miner
type Common struct {
	API    API
}

// API contains configs for API endpoint
type API struct {
	ListenAddress       string
	RemoteListenAddress string
	Timeout             Duration
}

// FullNode is a full node config
type FullNode struct {
	Common
}

type MinerConfig struct {
	Common

	BlockRecord string
}

func defCommon() Common {
	return Common{
		API: API{
			ListenAddress: "/ip4/127.0.0.1/tcp/1234/http",
			Timeout:       Duration(30 * time.Second),
		},
	}
}

// DefaultFullNode returns the default config
func DefaultFullNode() *FullNode {
	return &FullNode{
		Common: defCommon(),
	}
}

func DefaultMinerConfig() *MinerConfig {
	minerCfg := &MinerConfig{
		Common:      defCommon(),
		BlockRecord: "localdb",
	}

	minerCfg.Common.API.ListenAddress = "/ip4/0.0.0.0/tcp/12308/http" //change default address

	return minerCfg
}

var _ encoding.TextMarshaler = (*Duration)(nil)
var _ encoding.TextUnmarshaler = (*Duration)(nil)

// Duration is a wrapper type for time.Duration
// for decoding and encoding from/to TOML
type Duration time.Duration

// UnmarshalText implements interface for TOML decoding
func (dur *Duration) UnmarshalText(text []byte) error {
	d, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*dur = Duration(d)
	return err
}

func (dur Duration) MarshalText() ([]byte, error) {
	d := time.Duration(dur)
	return []byte(d.String()), nil
}

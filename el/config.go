package el

import (
	"errors"

	"gopkg.in/gcfg.v1"
)

// Server Config
type ElServerConfig struct {
	HopStart    int
	HopEnd      int
	ListenAddr  string
	Addr        string
	MTU         int
	Key         string
	FixMSS      bool
	MorphMethod string
	PeerTimeout int
	Up          string
	Down        string
}

// Client Config
type ElClientConfig struct {
	Server             string
	HopStart           int
	HopEnd             int
	Key                string
	MTU                int
	FixMSS             bool
	Local              bool
	MorphMethod        string
	Redirect_gateway   bool
	Net_gateway        []string // Deprecated
	Up                 string
	Down               string
	Heartbeat_interval int
}

type ElConfig struct {
	Default struct {
		Mode string
	}
	Server ElServerConfig
	Client ElClientConfig
}

func ParseElConfig(filename string) (interface{}, error) {
	cfg := new(ElConfig)
	err := gcfg.ReadFileInto(cfg, filename)
	if err != nil {
		return nil, err
	}
	switch cfg.Default.Mode {
	case "server":
		return cfg.Server, nil
	case "client":
		return cfg.Client, nil
	default:
		return nil, errors.New("Wrong mode")
	}
}

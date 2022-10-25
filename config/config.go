package config

import (
	"chain-api-imgo/resource"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	cfg     *Config = nil
	cfgFile string
)

type ListenConfig struct {
	Addr string `yaml:"addr"`
}

type LogConfig struct {
	Path        string `yaml:"path"`
	Level       string `yaml:"level"`
	StdErrLevel string `yaml:"std_err_level"`
}
type RedisConfig struct {
	Addr    string        `yaml:"addr,flow"`
	PoolNum int           `yaml:"pool_num,flow"`
	Timeout time.Duration `yaml:"timeout,flow"`
}
type MysqlConfig struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Max  int    `yaml:"max"`
}

type RemoteNodeConfig struct {
	Address string `yaml:"address"`
	Ca      string `yaml:"ca_path"`
}

type ChainNode struct {
	ChainId string             `yaml:"chain_id"`
	OrgId   string             `yaml:"org_id"`
	Tls     bool               `yaml:"tls"`
	TlsHost string             `yaml:"tls_host"`
	ConnCnt int                `yaml:"conn_cnt"`
	Remotes []RemoteNodeConfig `yaml:"remotes"`
	//UserConf *CertConf `yaml:"user"`
	//Pkcs11                sdk.Pkcs11Config  `yaml:"pkcs11"`
}
type ChainAdminCrts struct {
	TlsKeyPath string `yaml:"tls_key_path"`
	TlsCrtPath string `yaml:"tls_crt_path"`
}

type ConnectClient struct {
	Key     string `yaml:"key"`
	Crt     string `yaml:"crt"`
	SignKey string `yaml:"sign_key"`
	SignCrt string `yaml:"sign_crt"`
}

type Config struct {
	Env            string           `yaml:"env"`
	Log            LogConfig        `yaml:"log,flow"`
	Redis          RedisConfig      `yaml:"cluster_redis"`
	Mysql          []MysqlConfig    `yaml:"mysql"`
	ChainNode      ChainNode        `yaml:"chain_node"`
	ChainAdminCrts []ChainAdminCrts `yaml:"chain_admin_crts"` // 链的root证书列表
	Client         ConnectClient    `yaml:"client"`
	// CaBaseConfig CaBaseConfig `yaml:"ca_base_config"`
	// ChainAdminName      []string              `yaml:"chain_admin_name"`
	//ChainUser           map[string]*ChainUser `yaml:"chain_user"`
	// RootCaConf          *CaConfig             `yaml:"root_config"`
	// IntermediateCaConfs []*ImCaConfig         `yaml:"intermediate_config"`
	// AccessControlConfs  []*AccessControlConf  `yaml:"access_control_config"`
}

/*
type ChainUser struct {
	TlsKeyPath  string `yaml:"tls_key_path"`
	TlsCrtPath  string `yaml:"tls_crt_path"`
	SignKeyPath string `yaml:"sign_key_path"`
	SignCrtPath string `yaml:"sign_crt_path"`
}








type AccessControlConf struct {
	AppRole string `yaml:"app_role"`
	AppId   string `yaml:"app_id"`
	AppKey  string `yaml:"app_key"`
}

type ChainNodeConfig struct {
	NodeAddr       string   `yaml:"node_addr"`
	ConnCnt        int      `yaml:"conn_cnt"`
	EnableTls      bool     `yaml:"enable_tls"`
	TrustRootPaths []string `yaml:"trust_root_paths"`
	TlsHostName    string   `yaml:"tls_host_name"`
}
*/

func (l *LogConfig) GetLogPath() string {
	return l.Path
}
func (l *LogConfig) GetLogLevel() string {
	return l.Level
}

func (l *LogConfig) GetLogErrLevel() string {
	return l.StdErrLevel
}

func (p *Config) GetLogPath() string {
	return p.Log.Path
}

func (p *Config) GetLogLevel() string {
	return p.Log.Level
}

func LoadAndSet(configFile string) (config *Config, err error) {
	cfg := new(Config)
	err = loadConfig(configFile, cfg)
	if err != nil {
		return cfg, err
	}
	cfgFile = configFile
	print(cfg.Log.Path + "\n")
	return cfg, err
}

func GetConfig() *Config {
	return cfg
}

func SetConfig(conf *Config) {
	cfg = conf
}

func Reload() {
	c := new(Config)
	err := loadConfig(cfgFile, c)
	if err != nil {
		return
	}
	cfg = c
}

func loadConfig(configFile string, config interface{}) (err error) {
	data, err := resource.Get(configFile)
	if err != nil {
		if b, err := ioutil.ReadFile(configFile); err == nil {
			data = b
		}
		return
	}
	return loadConfigFromBytes(data, config)
}

func loadConfigFromBytes(data []byte, config interface{}) error {
	return yaml.Unmarshal(data, config)
}

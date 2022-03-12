package conf

type Config struct {
	Address  string      `json:"address"`
	Port     int         `json:"port"`
	Database Database    `json:"database"`
	Cache    CacheConfig `json:"cache"`
	TempDir  string      `json:"temp_dir"`
}

type CacheConfig struct {
	Expiration      int64 `json:"expiration"`
	CleanupInterval int64 `json:"cleanup_interval"`
}

type Database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
	DBFile      string `json:"db_file"`
	SslMode     string `json:"ssl_mode"`
}

func DefaultConfig() *Config {
	return &Config{
		Address: "0.0.0.0",
		Port:    7777,
		TempDir: "data/temp",
		Database: Database{
			Type:        "sqlite3",
			Port:        0,
			TablePrefix: "x_",
			DBFile:      "data/data.db",
			SslMode:     "disable",
		},
		Cache: CacheConfig{
			Expiration:      60,
			CleanupInterval: 120,
		},
	}
}

package config

// DatabaseConf conf about MongoDB
type Database struct {
	Host     string `toml:"host"`
	DB       string `toml:"database"`
	HTTPS    bool   `toml:"https"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

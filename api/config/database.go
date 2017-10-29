package config

// Database conf about MongoDB
type Database struct {
	Host     string `toml:"host"`
	DB       string `toml:"database"`
	SSL      bool   `toml:"SSL"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

package config

type WebhookChannel struct {
	Name    string `toml:"channel"`
	Pattern string `toml:"pattern"`
}

type Webhook struct {
	URL      string           `toml:"url"`
	Icon     string           `toml:"icon"`
	Channels []WebhookChannel `toml:"channel"`
}

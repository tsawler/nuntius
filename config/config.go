package config

// File config file
type File struct {
	Host      string        `yaml:"host"` // The host, eg: :8080 will start on 0.0.0.0:8080
	SSL       SSL           `yaml:"ssl"`
	Profiling bool          `yaml:"profiling"`
	Apps      []Application `yaml:"apps"`
}

// SSL related configuration options
type SSL struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	KeyFile  string `yaml:"key_file"`
	CertFile string `yaml:"cert_file"`
}

// Application related configuration options
type Application struct {
	Name       string   `yaml:"name"`
	AppID      string   `yaml:"app_id"`
	Key        string   `yaml:"key"`
	Secret     string   `yaml:"secret"`
	OnlySSL    bool     `yaml:"only_ssl"`
	Enabled    bool     `yaml:"enabled"`
	UserEvents bool     `yaml:"user_events"`
	WebHooks   Webhooks `yaml:"webhooks"`
}

// Webhooks related configuration options
type Webhooks struct {
	Enabled bool   `yaml:"enabled"`
	URL     string `yaml:"url"`
}

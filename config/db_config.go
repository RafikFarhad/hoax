package config

type DbConfig struct {
	Agent    string `ini:"db_agent"`
	Path     string `ini:"db_path"` // required for sqlite
	Host     string `ini:"db_url"`
	Port     string `ini:"db_port"`
	Name     string `ini:"db_name"`
	User     string `ini:"db_user"`
	Password string `ini:"db_password"`
	Log      bool   `ini:"db_log"`
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		Log: false,
	}
}

package config

type LogConfig struct {
	Agent string `ini:"log_agent"`
	Level string `ini:"log_level"`
	// file logger
	FilePath string `ini:"log_file_path"`
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Agent: "std",
	}
}

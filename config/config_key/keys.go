package config_key

type Key string

const (
	Verbose  Key = "verbose"
	Config   Key = "config"
	LogLevel Key = "log_level"
	LogFile  Key = "log_file"
	Port     Key = "port"
)

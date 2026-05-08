package config_keys

type Key int

const (
	InvalidKey Key = iota
	Port
	Verbose
	ConfigurationFile
	LogLevel
	LogFile
)

func (k Key) String() string {
	switch k {
	case Port:
		return "Port"
	case Verbose:
		return "Verbose"
	case ConfigurationFile:
		return "ConfigurationFile"
	case LogLevel:
		return "LogLevel"
	case LogFile:
		return "LogFile"
	default:
		return "InvalidKey"
	}
}

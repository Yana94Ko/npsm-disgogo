package config

import (
	"os"
)

type EnvVars struct {
	Token          string
	DefaultChannel string
	ChanWoongBlog  string
	YanaBlog       string
}

func LoadEnvVars() EnvVars {
	return EnvVars{
		Token:          getEnv("DISCORD_TOKEN", ""),
		DefaultChannel: getEnv("DEFAULT_CHANNEL", ""),
		ChanWoongBlog:  getEnv("CHANWOONG_BLOG", ""),
		YanaBlog:       getEnv("YANA_BLOG", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

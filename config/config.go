package config

import (
	"fmt"
	"strings"

	svccfg "source.golabs.io/go-libs/service_commons/config"
)

var cfg *svccfg.Config

func Load() {
	// tempcfg var has been used intentionally to avoid reads on cfg when LoadWithOptions might be running
	// avoids concurrent map read and write panic
	tempcfg := &svccfg.Config{}
	tempcfg.LoadWithOptions(map[string]interface{}{"configPath": "$GOPATH/src/github.com/mahendrakariya/add"})
	cfg = tempcfg
}

func Get() *svccfg.Config {
	return cfg
}

func AppNodeIP() string {
	return cfg.GetValue("APP_NODE_IP")
}

func AppNodeID() string {
	return cfg.GetValue("APP_NODE_ID")
}

func ConsulAddress() string {
	return fmt.Sprintf("%s:%d",
		cfg.GetOptionalValue("CONSUL_CLIENT_HOST", "localhost"),
		cfg.GetOptionalIntValue("CONSUL_CLIENT_PORT", 8500))
}

func ConsulTags() []string {
	return strings.Split(cfg.GetOptionalValue("CONSUL_TAGS", "grpc"), ",")
}

func ConsulCheckURL() string {
	return fmt.Sprintf("http://%s:%d/ping", AppNodeIP(), Port())
}

func ConsulServiceCheckInterval() string {
	return cfg.GetOptionalValue("CONSUL_SERVICE_CHECK_INTERVAL", "3s")
}

func Port() int {
	return cfg.GetOptionalIntValue("PORT", 3000)
}

func LogLevel() string {
	return cfg.GetOptionalValue("LOG_LEVEL", "INFO")
}

func AppName() string {
	return cfg.GetOptionalValue("APP_NAME", "add-grpc")
}

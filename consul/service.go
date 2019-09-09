package consul

import (
	"github.com/mahendrakariya/add/config"

	"source.golabs.io/go-libs/consul_client/clients/consul"
	consulcfg "source.golabs.io/go-libs/consul_client/config"
	"source.golabs.io/go-libs/service_commons/logger"
)

var Client consul.Client

func InitializeClient() error {
	var err error
	Client, err = consul.NewConsulClient(GetConfig())
	if err != nil {
		logger.Get().WithField("error", err.Error()).Error("Error connnecting consul agent")
		return err
	}
	return nil
}

func Register() error {
	err := Client.RegisterToConsul()
	if err != nil {
		logger.Get().WithField("error", err.Error()).Error("Error registering the sevice")
		return err
	}
	return nil
}

func DeRegister() error {
	err := Client.DeRegisterFromConsul(config.AppNodeID())
	if err != nil {
		logger.Get().WithField("error", err.Error()).Error("Error deregistering service from consul")
	}
	return nil
}

func GetConfig() *consulcfg.Config {
	c := consulcfg.NewConfig()
	c.ConsulAddress = config.ConsulAddress()
	c.AppName = config.AppName()
	c.Port = config.Port()
	c.NodeID = config.AppNodeID()
	c.NodeIP = config.AppNodeIP()
	c.ConsulTags = config.ConsulTags()
	c.ConsulCheckURL = config.ConsulCheckURL()
	c.ConsulCheckInterval = config.ConsulServiceCheckInterval()
	return c
}

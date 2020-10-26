package kobe

import (
	"github.com/KubeOperator/kobe/api"
	kobeClient "github.com/KubeOperator/kobe/pkg/client"
	"github.com/spf13/viper"
	"io"
)

type Interface interface {
	RunPlaybook(name string) (string, error)
	Watch(writer io.Writer, taskId string) error
	GetResult(taskId string) (*api.Result, error)
	SetVar(key string, value string)
}

type Config struct {
	Inventory api.Inventory
}

type Kobe struct {
	Project   string
	Inventory api.Inventory
	client    *kobeClient.KobeClient
}

func NewAnsible(c *Config) *Kobe {
	c.Inventory.Vars = map[string]string{}
	host := viper.GetString("kobe.host")
	port := viper.GetInt("kobe.port")
	return &Kobe{
		Project:   "ko",
		Inventory: c.Inventory,
		client:    kobeClient.NewKobeClient(host, port),
	}
}

func (k *Kobe) RunPlaybook(name string) (string, error) {
	result, err := k.client.RunPlaybook(k.Project, name, k.Inventory)
	if err != nil {
		return "", err
	}
	return result.Id, nil
}

func (k *Kobe) SetVar(key string, value string) {
	k.Inventory.Vars[key] = value
}

func (k *Kobe) RunAdhoc(pattern, module, param string) (string, error) {
	result, err := k.client.RunAdhoc(pattern, module, param, k.Inventory)
	if err != nil {
		return "", nil
	}
	return result.Id, nil
}

func (k *Kobe) Watch(writer io.Writer, taskId string) error {
	err := k.client.WatchRun(taskId, writer)
	if err != nil {
		return err
	}
	return nil
}

func (a *Kobe) GetResult(taskId string) (*api.Result, error) {
	return a.client.GetResult(taskId)
}

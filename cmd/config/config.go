package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	GreenLightEnvelope Envelope[GreenLight]
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./cmd/config")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("viper读取配置文件错误")
	}
	GreenLightEnvelope = GetServiceConfig("green_light", viper.GetViper())
	return nil
}

type Envelope[T any] struct {
	ServiceName string
	Data        T
	Err         error
}

func GetServiceConfig(serviceName string, v *viper.Viper) Envelope[GreenLight] {
	if serviceName == "green_light" {
		g, err := GetGreenLightConfig(v)
		return Envelope[GreenLight]{
			ServiceName: "green_light",
			Data:        g,
			Err:         err,
		}
	}

	return Envelope[GreenLight]{
		ServiceName: "Not Found",
		Data:        GreenLight{},
		Err:         fmt.Errorf("Service Not Found"),
	}
}

type GreenLight struct {
	Port int
	Env  string
}

func GetGreenLightConfig(v *viper.Viper) (GreenLight, error) {
	greenLightConfig := v.Sub("services").Sub("green_light")
	if greenLightConfig == nil {
		return GreenLight{}, fmt.Errorf("找不到对应的greenLight配置")
	}
	port := greenLightConfig.GetInt("port")
	env := greenLightConfig.GetString("env")
	return GreenLight{
		Port: port,
		Env:  env,
	}, nil
}

package main

import (
	"fmt"
	"github.com/looksaw/go_greenlight/cmd/config"
	"github.com/looksaw/go_greenlight/internal/router"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	//GreenLight的配置
	GreenLightConfig config.Envelope[config.GreenLight]
}

type application struct {
	config Config
	logger zerolog.Logger
}

func main() {
	//初始化Config
	config.InitConfig()
	config := Config{
		GreenLightConfig: config.GreenLightEnvelope,
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger()
	application := application{
		config: config,
		logger: logger,
	}

	log.Logger = logger

	r := router.SetupRouter()
	rPost := fmt.Sprintf(":%d", application.config.GreenLightConfig.Data.Port)
	r.Run(rPost)

}

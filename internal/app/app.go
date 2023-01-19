package app

import (
	"fmt"
	"github.com/miromax42/dots-backend/config"
)

func Run(cfg *config.Config) {
	fmt.Printf("Hello, %s!", cfg.App.Name)
}

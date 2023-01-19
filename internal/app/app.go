package app

import (
	"fmt"
	"github.com/miromax42/dots-backend/config"
	"github.com/miromax42/dots-backend/pkg/game"
)

func Run(cfg *config.Config) {
	fmt.Printf("Hello, %s!\n", cfg.App.Name)

	g := game.New()

	fmt.Println(g.Render())
}

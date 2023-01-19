package app

import (
	"bufio"
	"fmt"
	"github.com/miromax42/dots-backend/config"
	"github.com/miromax42/dots-backend/pkg/game"
	"os"
	"strings"
)

func Run(cfg *config.Config) {
	fmt.Printf("Hello, %s!\n", cfg.App.Name)

	g := game.New()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(g.Render())

		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		s = strings.Trim(s, "\r\n")

		err = g.Move(s)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("exit")
}

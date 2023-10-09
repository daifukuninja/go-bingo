package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daifukuninja/go-bingo/bingo/elm/model"
)

func main() {
	if _, err := tea.NewProgram(model.Tag{}).Run(); err != nil {
		fmt.Println("uh oh:", err)
		os.Exit(1)
	}
}

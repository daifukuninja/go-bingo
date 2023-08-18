package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daifukuninja/go-bingo/bingo/contents"
)

type model struct {
	tag int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
		m.tag++
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	producer := contents.NewProducer()
	return producer.GetContents(m.tag)
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Println("uh oh:", err)
		os.Exit(1)
	}
}

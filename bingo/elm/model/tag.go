package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daifukuninja/go-bingo/bingo/contents"
)

type Tag struct {
	Tag int
}

func (m Tag) Init() tea.Cmd {
	return nil
}

func (m Tag) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
		m.Tag++
		return m, nil
	}

	return m, nil
}

func (m Tag) View() string {
	producer := contents.NewProducer()
	return producer.GetContents(m.Tag)
}

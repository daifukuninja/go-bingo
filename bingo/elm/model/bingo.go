package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daifukuninja/go-bingo/bingo/contents"
)

type Bingo struct {
	Tag int
}

func (m Bingo) Init() tea.Cmd {
	return nil
}

func (m Bingo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Bingo) View() string {
	producer := contents.NewProducer()
	return producer.GetContents(m.Tag)
}

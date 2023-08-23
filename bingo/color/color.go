package color

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

const (
	LightGreen = "119"
	DeepPink4  = "89"
	Magenta3   = "164"
)

// TODO: termenvをuninstallする
var (
	lbingoOn  = lipgloss.NewStyle().Foreground(lipgloss.Color(Magenta3))
	lbingoOff = lipgloss.NewStyle().Foreground(lipgloss.Color(DeepPink4))
)

func BingoOn(tag int) string {
	return lbingoOn.Render(strconv.Itoa(tag))
}

func BingoOFf(tag int) string {
	return lbingoOff.Render(strconv.Itoa(tag))
}

package color

import (
	"strconv"

	"github.com/muesli/termenv"
)

const (
	LightGreen = "119"
	DeepPink4  = "89"
	Magenta3   = "164"
)

var (
	colorProf = termenv.EnvColorProfile().Color
	bingoOn   = termenv.Style{}.Foreground(colorProf(Magenta3)).Styled
	bingoOff  = termenv.Style{}.Foreground(colorProf(DeepPink4)).Styled
)

func BingoOn(tag int) string {
	return bingoOn(strconv.Itoa(tag))
}

func BingoOFf(tag int) string {
	return bingoOff(strconv.Itoa(tag))
}

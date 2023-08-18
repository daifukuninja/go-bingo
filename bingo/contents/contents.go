package contents

import (
	_ "embed"
	"strings"
	"text/template"

	c "github.com/daifukuninja/go-bingo/bingo/color"
)

type Producer struct {
	tmpl *template.Template
}

//go:embed template/view.tmpl
var templateString string

func NewProducer() *Producer {
	tmpl, _ := template.New("view.tmpl").Parse(templateString)
	return &Producer{
		tmpl: tmpl,
	}
}

func (p *Producer) GetContents(tag int) string {
	type Data struct {
		Tag string
	}
	writer := new(strings.Builder)
	data := &Data{
		Tag: c.BingoOn(tag),
	}
	p.tmpl.Execute(writer, data)

	return writer.String()
}

package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"text/template"
)

var (
	ErrNilConfig = errors.New("nil config")
)

type TextCVGenerator struct {
	config *Config
}

func NewTextCVGenerator(config *Config) (*TextCVGenerator, error) {
	if config == nil {
		return nil, ErrNilConfig
	}
	return &TextCVGenerator{
		config: config,
	}, nil
}

func (g *TextCVGenerator) readTpl() (TextTemplate, error) {
	tplName := "templates/default.md"
	if len(g.config.Style.Template) > 0 {
		tplName = g.config.Style.Template
	}
	data, err := os.ReadFile(tplName)
	if err != nil {
		return "", fmt.Errorf("read template: %v", err)
	}
	tpl := string(data)
	return TextTemplate(tpl), nil
}

func (g *TextCVGenerator) replace() {

}

func (g *TextCVGenerator) Generate() error {
	content, err := g.readTpl()
	if err != nil {
		return err
	}
	t, err := template.New(g.config.Style.Template).Parse(string(content))
	if err != nil {
		return err
	}
	outpath := path.Join("build", "output.txt")
	outfile, err := os.Create(outpath)
	if err != nil {
		return err
	}
	return t.Execute(outfile, g.config.Data)
}

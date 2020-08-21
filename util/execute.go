package util

import (
	"fmt"
	"strings"
	"text/template"
)

func Execute(fileName string, value interface{}) (string, error) {
	tpl, err := template.ParseFiles(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to parse template file %s with cause %w", fileName, err)
	}
	var buf strings.Builder
	err = tpl.Execute(&buf, value)
	if err != nil {
		return "", fmt.Errorf("failed to render value to template with cause %w", err)
	}
	return buf.String(), nil
}

type Render struct {
	tpl *template.Template
}

func NewRender(fileName string) (*Render, error) {
	tpl, err := template.ParseGlob(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create template render with cause %w", err)
	}
	r := &Render{
		tpl:tpl,
	}
	return r, nil
}

func (r *Render) Rend(value interface{}) (string, error) {
	var buf strings.Builder
	err := r.tpl.Execute(&buf, value)
	if err != nil {
		return "", fmt.Errorf("failed to render value(%v) to templte with cause %w", value, err)
	}
	return buf.String(), nil
}

package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "template"
}

func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string {
	return "world"
}

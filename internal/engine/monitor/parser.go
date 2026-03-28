package monitor

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/CFF4HA/Staircase/internal/engine/parser"
	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/antlr4-go/antlr/v4"
)

const (
	regexTemplate = `/^(.*){{- range . -}}({{- printf "%s" . -}}){{- end -}}(.*)$/`
)

type StaircaseProcessor struct {
	listener *staircaseListener
}

func NewStaircaseProcessor() *StaircaseProcessor {
	return &StaircaseProcessor{
		listener: &staircaseListener{},
	}
}

func (s *StaircaseProcessor) Process(input string) *types.RetrievalQuery {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewstaircaseLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewstaircaseParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(s.listener, p.Job())
	return s.listener.query
}

// ---

type staircaseListener struct {
	*parser.BasestaircaseListener

	query     *types.RetrievalQuery
	regexTmpl *template.Template
}

func (s *staircaseListener) EnterJob(ctx *parser.JobContext) {
	s.query = new(types.RetrievalQuery)

	if t, err := template.New("regex").Parse(regexTemplate); err == nil {
		s.regexTmpl = t
	} else {
		panic(fmt.Errorf("failed to parse regex template, err=%v", err))
	}
}

func (s *staircaseListener) ExitRule_navigation(ctx *parser.Rule_navigationContext) {
	n := types.Navigation{}

	if ctx.DRILLDOWN() != nil {
		n.Keyword = "drilldown"
	} else if ctx.SIBLING() != nil {
		n.Keyword = "sibling"
	} else if ctx.CHILD() != nil {
		n.Keyword = "child"
	} else if ctx.PARENT() != nil {
		n.Keyword = "parent"
	}

	if ctx.INT() != nil {
		i, err := strconv.Atoi(ctx.INT().GetText())
		if err != nil {
			panic(err)
		}

		n.Quantifier = i
	}

	s.query.Navigations = append(s.query.Navigations, n)
}

func (s *staircaseListener) ExitRule_keyword_list(ctx *parser.Rule_keyword_listContext) {
	var builder strings.Builder
	var keywords []string

	for _, keyword := range ctx.AllSTRING() {
		keywords = append(keywords, strings.Trim(keyword.GetText(), "\""))
	}

	if err := s.regexTmpl.Execute(&builder, keywords); err != nil {
		panic(err)
	}
	s.query.Regex = builder.String()
}

func (s *staircaseListener) ExitRule_site(ctx *parser.Rule_siteContext) {
	if ctx.URL() != nil {
		s.query.Site = ctx.URL().GetText()
	}
}

func (s *staircaseListener) ExitRule_type_declaration(ctx *parser.Rule_type_declarationContext) {
	var action string

	if ctx.CLICK() != nil {
		action = "click"
	} else if ctx.COLLECT() != nil {
		action = "extract-text"
	}

	s.query.Action = types.Action{
		Type: action,
	}
}

func (s *staircaseListener) ExitRule_in_value(ctx *parser.Rule_in_valueContext) {
	action := "input"
	modifier := strings.Trim(ctx.STRING().GetText(), "\"")

	s.query.Action = types.Action{
		Type:  action,
		Value: modifier,
	}
}

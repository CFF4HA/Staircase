// Code generated from staircase/staircase.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // staircase

import "github.com/antlr4-go/antlr/v4"

// BasestaircaseListener is a complete listener for a parse tree produced by staircaseParser.
type BasestaircaseListener struct{}

var _ staircaseListener = &BasestaircaseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasestaircaseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasestaircaseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasestaircaseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasestaircaseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJob is called when production job is entered.
func (s *BasestaircaseListener) EnterJob(ctx *JobContext) {}

// ExitJob is called when production job is exited.
func (s *BasestaircaseListener) ExitJob(ctx *JobContext) {}

// EnterRule_navigation is called when production rule_navigation is entered.
func (s *BasestaircaseListener) EnterRule_navigation(ctx *Rule_navigationContext) {}

// ExitRule_navigation is called when production rule_navigation is exited.
func (s *BasestaircaseListener) ExitRule_navigation(ctx *Rule_navigationContext) {}

// EnterRule_navigation_sequence is called when production rule_navigation_sequence is entered.
func (s *BasestaircaseListener) EnterRule_navigation_sequence(ctx *Rule_navigation_sequenceContext) {}

// ExitRule_navigation_sequence is called when production rule_navigation_sequence is exited.
func (s *BasestaircaseListener) ExitRule_navigation_sequence(ctx *Rule_navigation_sequenceContext) {}

// EnterRule_keyword_list is called when production rule_keyword_list is entered.
func (s *BasestaircaseListener) EnterRule_keyword_list(ctx *Rule_keyword_listContext) {}

// ExitRule_keyword_list is called when production rule_keyword_list is exited.
func (s *BasestaircaseListener) ExitRule_keyword_list(ctx *Rule_keyword_listContext) {}

// EnterRule_site is called when production rule_site is entered.
func (s *BasestaircaseListener) EnterRule_site(ctx *Rule_siteContext) {}

// ExitRule_site is called when production rule_site is exited.
func (s *BasestaircaseListener) ExitRule_site(ctx *Rule_siteContext) {}

// EnterRule_type_declaration is called when production rule_type_declaration is entered.
func (s *BasestaircaseListener) EnterRule_type_declaration(ctx *Rule_type_declarationContext) {}

// ExitRule_type_declaration is called when production rule_type_declaration is exited.
func (s *BasestaircaseListener) ExitRule_type_declaration(ctx *Rule_type_declarationContext) {}

// EnterRule_in_value is called when production rule_in_value is entered.
func (s *BasestaircaseListener) EnterRule_in_value(ctx *Rule_in_valueContext) {}

// ExitRule_in_value is called when production rule_in_value is exited.
func (s *BasestaircaseListener) ExitRule_in_value(ctx *Rule_in_valueContext) {}

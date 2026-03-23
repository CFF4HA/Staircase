// Code generated from staircase/staircase.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // staircase

import "github.com/antlr4-go/antlr/v4"

// staircaseListener is a complete listener for a parse tree produced by staircaseParser.
type staircaseListener interface {
	antlr.ParseTreeListener

	// EnterJob is called when entering the job production.
	EnterJob(c *JobContext)

	// EnterRule_navigation is called when entering the rule_navigation production.
	EnterRule_navigation(c *Rule_navigationContext)

	// EnterRule_navigation_sequence is called when entering the rule_navigation_sequence production.
	EnterRule_navigation_sequence(c *Rule_navigation_sequenceContext)

	// EnterRule_keyword_list is called when entering the rule_keyword_list production.
	EnterRule_keyword_list(c *Rule_keyword_listContext)

	// EnterRule_site is called when entering the rule_site production.
	EnterRule_site(c *Rule_siteContext)

	// EnterRule_type_declaration is called when entering the rule_type_declaration production.
	EnterRule_type_declaration(c *Rule_type_declarationContext)

	// EnterRule_in_value is called when entering the rule_in_value production.
	EnterRule_in_value(c *Rule_in_valueContext)

	// ExitJob is called when exiting the job production.
	ExitJob(c *JobContext)

	// ExitRule_navigation is called when exiting the rule_navigation production.
	ExitRule_navigation(c *Rule_navigationContext)

	// ExitRule_navigation_sequence is called when exiting the rule_navigation_sequence production.
	ExitRule_navigation_sequence(c *Rule_navigation_sequenceContext)

	// ExitRule_keyword_list is called when exiting the rule_keyword_list production.
	ExitRule_keyword_list(c *Rule_keyword_listContext)

	// ExitRule_site is called when exiting the rule_site production.
	ExitRule_site(c *Rule_siteContext)

	// ExitRule_type_declaration is called when exiting the rule_type_declaration production.
	ExitRule_type_declaration(c *Rule_type_declarationContext)

	// ExitRule_in_value is called when exiting the rule_in_value production.
	ExitRule_in_value(c *Rule_in_valueContext)
}

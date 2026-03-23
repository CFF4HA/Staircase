// The following is the staircase grammar! We use staircase to navigate
// a website's DOM-Tree! It is based on the idea that you can select an anchor (using regex)
// and then navigate ('parent, child, sibling, drilldown'), and finally, perform 
// an action ('click', 'input', 'extract-text').
//
// Here's what it looks like:
// 
// // "Collection Job w/ Predefined Site"
// click @https://.../...
//     keywords "keyword1" "keyword2" // translates to a singular regex expression, treated as AND.
//     parent child[2] child[2] drilldown sibling
//
// Notes: 
//  - If there is no "site" provided, then a previous context will be used, this is handled
//  at the API level though, not at the parser level.

// Parser Rules

grammar staircase;

// Parser Rules
job: rule_type_declaration rule_site? KEYWORDS rule_keyword_list rule_navigation_sequence rule_in_value? EOF;

rule_navigation: (PARENT | CHILD | DRILLDOWN | SIBLING) (LBRACK INT RBRACK)?;

rule_navigation_sequence: rule_navigation+;
rule_keyword_list: STRING+;
rule_site: AT URL;

rule_type_declaration: CLICK | COLLECT | INPUT;
rule_in_value: STRING;

// Lexer Rules
CLICK: 'click';
COLLECT: 'collect';
INPUT: 'input';

PARENT: 'parent';
CHILD: 'child';
DRILLDOWN: 'drilldown';
SIBLING: 'sibling';

KEYWORDS: 'keywords';
LBRACK: '[';
RBRACK: ']';
AT: '@';

URL: 'http' 's'? '://' [-a-zA-Z0-9./?=_#]+;

// Supports escaped quotes inside the string
STRING: '"' ( '\\"' | ~('"'|'\\') )* '"';

INT: [0-9]+;
WS: [ \t\r\n]+ -> skip;

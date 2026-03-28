package chrome

const (
	engine = `// Based on the RetrievalQeuery, we will now return the relevant items for 
// the chromedp context. The way this will work is that we will leverage 
// the template library to compile it.
var logs = {
	Result: "",
	Logs: []
};

var node = undefined;
var root = undefined;

while (root == null || root == undefined) {
	root = getAnchor({{.Regex}});
}

{{ range.Navigations }}
node = root.navigate("{{.Keyword}}", {{.Quantifier }});
while (node == null || node == undefined) {
	node = root.navigate("{{.Keyword}}", {{.Quantifier }});
}
root = node;
node = undefined;
{{ end }}

{{ if eq .Action.Type "click" }}
root.node.click();
{{ else if eq .Action.Type "extract-text" }}
logs.Result = root.node.textContent.trim();
{{ else if eq .Action.Type "input" }}
root.node.value = "{{.Action.Value}}";
{{ end }}

logs;

`

	regex = `
// This file is a javascript file for a DOM parsing engine powering our chromedp based web scraper.
// It is based on an architecture where you a query has an "Anchor", and a "Target". 
//
// Terminology: 
//  - "Anchor": This is a node in the DOM relative to which we can reference other elements.
//  - "Target": This is a node in the DOM on which an action will occur (i.e., click, extractText, etc).
//
// Notes:
//   - An "Anchor" can itself be a "Target".
//   - The first "Anchor" will be selected through Regex.
//
//
// Note: 
//   - This file will be a gotemplate compiled with specific values like the 
//   selector to use for a given context and the attribute to test the regex on.


// The "Anchor" will have an attribute called "Node" which is of type "Element".
class Anchor {
  constructor(node) {
    this.node = node;
  }

  fromRegex(regex) {
    return getAnchor(regex);
  }

  // This function will be used to explore other nodes 
  // relative to the "Anchor". We will do so by navigating 
  // - up to the parent 
  // - or selecting one of the children.
  navigate(keyword, quantifier) {
    switch (keyword) {
      case "parent":
				if (this.node.parentElement == null) {
					return this;
				}
        var n = this.node.parentElement;
        for (var i = 0; i < quantifier - 1; i++) {
          n = n.parentElement;
        }

        return new Anchor(n);
      case "child":
				if (this.node.children.length == 0) {
				return this;
				}
        if (this.node.children.length < quantifier) {
          return null;
        }

				if (quantifier > 0)
					return new Anchor(this.node.children[quantifier - 1]);
				return new Anchor(this.node.children[0]);
      case "sibling":
				if (this.node.nextElementSibling == null) {
					return this;
				}
				while (this.node.nextElementSibling == null) {
					
				}
        return new Anchor(this.node.nextElementSibling);
      case "drilldown":
        // this case will iterate through all the children until 
        // it reaches the furthest one down, defaulting to the first 
        // child in the children array.
				if (this.node.children.length == 0) {
					return this;
				}
        var n = this.node.children[0];
        while (n.children.length > 0) {
          n = n.children[0];
        }

        return new Anchor(n);
      default:
        return null;
    }
  }
};

function getElementsByRegex(regex) {
  const elements = Array.from(document.querySelectorAll('*')).filter(el => regex.test(el.textContent));
  return elements;
};

function getAnchor(regex) {
  var elements = getElementsByRegex(regex);
	const MAX_OPERATION_COUNT = 20000 * 2;
	var count = 0;
	while (elements.length == 0 && count < MAX_OPERATION_COUNT) {
		elements = getElementsByRegex(regex);
		count += 1;
	}

  return new Anchor(elements[elements.length - 1]);
}
	`
)

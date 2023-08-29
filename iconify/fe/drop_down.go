package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feDropDown0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feDropDown1" fill="currentColor"><path id="feDropDown2" d="m5 8l7 8l7-8z"/></g></g>`),
		g.Group(children),
	)
}
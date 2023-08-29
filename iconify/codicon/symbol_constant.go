package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolConstant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4 6h8v1H4V6zm8 3H4v1h8V9z"/><path d="m1 4l1-1h12l1 1v8l-1 1H2l-1-1V4zm1 0v8h12V4H2z"/></g>`),
		g.Group(children),
	)
}
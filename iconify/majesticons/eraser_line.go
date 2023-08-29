package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#majesticonsEraserLine0)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.607 6.535l-7.779 7.779a2 2 0 0 0 0 2.828l1.536 1.536a2 2 0 0 0 1.414.585h6.243M11.607 6.535l2.12-2.12a2 2 0 0 1 2.83 0l4.242 4.242a2 2 0 0 1 0 2.828l-2.121 2.122m-7.071-7.072l7.07 7.072m0 0l-5.656 5.656m0 0H20"/></g><defs><clipPath id="majesticonsEraserLine0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}
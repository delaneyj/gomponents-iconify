package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weightlifting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-miterlimit="2" stroke-width="4"><path d="M24 27a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 9h40M4 4v10M44 4v10M11 9v17.1L24 34l13-8V9M24 34v10"/></g>`),
		g.Group(children),
	)
}
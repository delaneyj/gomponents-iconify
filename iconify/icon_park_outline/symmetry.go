package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symmetry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="m4 15l20-6v30L4 33V15Zm20-6l20 6v18l-20 6V9Z"/><path stroke-linecap="round" d="M24 4v40"/></g>`),
		g.Group(children),
	)
}
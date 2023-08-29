package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M20 5v6.5a3 3 0 0 1-3 3H5"/><path d="M7.5 18L4 14.5L7.5 11"/></g>`),
		g.Group(children),
	)
}
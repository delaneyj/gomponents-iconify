package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogicXnor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12h-2M2 9h4m-4 6h4m-1 4c1.778-4.667 1.778-9.333 0-14m3 0c10.667 2.1 10.667 12.6 0 14c1.806-4.667 1.806-9.333 0-14z"/><path d="M16 12a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geometry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 21l4-12m2 0l1.48 4.439m.949 2.847L17 21M10 7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-6 5c1.526 2.955 4.588 5 8 5c3.41 0 6.473-2.048 8-5m-8-7V3"/>`),
		g.Group(children),
	)
}
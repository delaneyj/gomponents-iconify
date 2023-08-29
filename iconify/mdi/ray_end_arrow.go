package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RayEndArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1 12l4 4v-3h12.17c.41 1.17 1.52 2 2.83 2a3 3 0 0 0 3-3a3 3 0 0 0-3-3c-1.31 0-2.42.83-2.83 2H5V8l-4 4Z"/>`),
		g.Group(children),
	)
}
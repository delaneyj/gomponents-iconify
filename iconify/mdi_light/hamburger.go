package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hamburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4h9a5 5 0 0 1 5 5H2a5 5 0 0 1 5-5m9 1H7C5.14 5 3.57 6.27 3.13 8h16.74c-.44-1.73-2.01-3-3.87-3m5 11a5 5 0 0 1-5 5H7a5 5 0 0 1-5-5h19M7 20h9c1.86 0 3.43-1.27 3.87-3H3.13c.44 1.73 2.01 3 3.87 3m5.5-10l2 2l2-2H19a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h8.5m2 3.41L12.09 11H4a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h15a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1h-2.09l-2.41 2.41Z"/>`),
		g.Group(children),
	)
}
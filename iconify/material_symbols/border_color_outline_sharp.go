package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderColorOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 24v-4h20v4H2Zm2-6v-3.75l9.05-9.05l3.75 3.75L7.75 18H4Zm2-2h.9L14 8.95L13.05 8L6 15.1v.9Zm11.925-8.15l-3.75-3.75l2.5-2.5l3.75 3.75l-2.5 2.5ZM6 16Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h2v6H3Zm3-2v-2h5q.425 0 .713-.288T12 16v-3H7q-.5 0-.8-.388t-.15-.887l2.4-8q.1-.325.35-.525t.6-.2h7.2q.35 0 .6.2t.35.525l2.4 8q.15.5-.15.888T19 13h-5v3q0 1.25-.875 2.125T11 19H6Z"/>`),
		g.Group(children),
	)
}
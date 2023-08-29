package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelicopterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13V7Q6.5 7 4.75 8.75T3 13h6Zm4 9H3v-2h10v2Zm2-3H1v-6q0-3.35 2.325-5.675T9 5h6v5h5l1-2h2v7l-8 .8V19Zm4-15H3V2h16v2Z"/>`),
		g.Group(children),
	)
}
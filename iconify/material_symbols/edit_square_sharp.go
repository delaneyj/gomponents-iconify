package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditSquareSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23.7v-18h10.925L7 12.625V19.7h7.075L21 12.75V23.7H3Zm6-6v-4.25l7.175-7.175l4.275 4.2l-7.2 7.225H9Zm12.875-8.65L17.6 4.85l2.525-2.525l4.2 4.275l-2.45 2.45Z"/>`),
		g.Group(children),
	)
}
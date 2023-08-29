package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeGMobiledataSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17v-2h5v-2H3v-2h5V9H3V7h7v10H3Zm18-6v6h-9V7h9v2h-7v6h5v-2h-2.5v-2H21Z"/>`),
		g.Group(children),
	)
}
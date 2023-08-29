package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkRingOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.35 14.85L16.9 13.4q.3-.275.463-.637t.162-.763q0-.4-.162-.763T16.9 10.6l1.45-1.45q.575.575.875 1.313t.3 1.537q0 .8-.3 1.538t-.875 1.312Zm2.45 2.45l-1.4-1.4q.775-.775 1.2-1.775T21.025 12q0-1.125-.425-2.125T19.4 8.1l1.4-1.4q1.075 1.05 1.65 2.425T23.025 12q0 1.5-.575 2.875T20.8 17.3ZM5 23V1h14v6h-2V6H7v12h10v-1h2v6H5Zm2-3v1h10v-1H7ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}
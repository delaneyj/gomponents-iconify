package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkRingOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM5 6.425l2 2V18h9.6l2.4 2.4V23H5V6.425ZM7 20v1h10v-1H7ZM7 4.175V4h10V3H7v1.175L5.15 2.3V1H19v6h-2V6H8.825L7 4.175ZM7 20v1v-1Zm11.35-5.15L16.9 13.4q.3-.275.463-.638t.162-.762q0-.4-.162-.763T16.9 10.6l1.45-1.45q.575.575.875 1.313t.3 1.537q0 .8-.3 1.538t-.875 1.312Zm2.45 2.45l-1.4-1.4q.775-.775 1.2-1.775T21.025 12q0-1.125-.425-2.125T19.4 8.1l1.4-1.4q1.075 1.05 1.65 2.425T23.025 12q0 1.5-.575 2.875T20.8 17.3ZM17 3v1v-1Z"/>`),
		g.Group(children),
	)
}
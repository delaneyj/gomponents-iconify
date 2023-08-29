package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayingCardsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.2 14.8l1.15-4.15L12.8 8.2l-1.15 4.15l3.55 2.45ZM4 18.825l-2.65-1.25L4 11.85v6.975ZM6 21v-8l2.9 8H6Zm5.25.575L5.425 5.55L16.75 1.425l5.825 16.025l-11.325 4.125Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeDownOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 15L1 11.5l1.05-1.05L3.8 12.2q-.15-.675-.225-1.35T3.5 9.5q0-2.05.675-3.975T6.125 2L7.2 3.075q-1.075 1.4-1.638 3.038T5 9.5q0 .65.075 1.288t.25 1.262l1.625-1.6L8 11.5L4.5 15Zm3.1 2.675l.95-2.075l2.9-.225l-3.15-8.6l1.875-.7l4.05 11.125l-2.5.175l3.675 1.7l6.175-2.25l-2.425-6.55l1.875-.7l3.125 8.45l-8.875 3.225L7.6 17.675Zm6.6-3.475l-1.7-4.7l1.875-.7l1.725 4.7l-1.9.7Zm2.825-1.025l-1.375-3.75l1.875-.7l1.375 3.75l-1.875.7Zm.125 1.975Z"/>`),
		g.Group(children),
	)
}
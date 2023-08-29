package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeUpOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.125 14q-1.275-1.6-1.95-3.525T3.5 6.5q0-.675.075-1.35T3.8 3.8L2.05 5.55L1 4.5L4.5 1L8 4.5L6.95 5.55l-1.625-1.6q-.175.625-.25 1.262T5 6.5q0 1.75.563 3.388T7.2 12.924L6.125 14ZM7.6 17.675l.95-2.075l2.9-.225l-3.15-8.6l1.875-.7l4.05 11.125l-2.5.175l3.675 1.7l6.175-2.25l-2.425-6.55l1.875-.7l3.125 8.45l-8.875 3.225L7.6 17.675Zm6.6-3.475l-1.7-4.7l1.875-.7l1.725 4.7l-1.9.7Zm2.825-1.025l-1.375-3.75l1.875-.7l1.375 3.75l-1.875.7Zm.125 1.975Z"/>`),
		g.Group(children),
	)
}
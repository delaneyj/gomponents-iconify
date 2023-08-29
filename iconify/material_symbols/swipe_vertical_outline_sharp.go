package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeVerticalOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-1.5h2.025q-1.475-1.8-2.25-3.975T1 12q0-2.35.775-4.525T4.025 3.5H2V2h5v5H5.5V4.1Q4.05 5.75 3.275 7.775T2.5 12q0 2.2.775 4.225T5.5 19.9V17H7v5H2Zm5.6-4.325l.95-2.075l2.9-.225l-3.15-8.6l1.875-.7l4.05 11.125l-2.5.175l3.675 1.7l6.175-2.25l-2.425-6.55l1.875-.7l3.125 8.45l-8.875 3.225L7.6 17.675Zm6.6-3.475l-1.7-4.7l1.875-.7l1.725 4.7l-1.9.7Zm2.825-1.025l-1.375-3.75l1.875-.7l1.375 3.75l-1.875.7Zm.125 1.975Z"/>`),
		g.Group(children),
	)
}
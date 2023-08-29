package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeRightOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.575 22L4.6 16l1.575-1.625l2.825.8V6h2v11.825l-2.425-.675l2.85 2.85H18v-7h2v9h-9.425ZM12 15v-5h2v5h-2Zm3 0v-4h2v4h-2ZM2 7q.75-2.65 3.55-4.325T12 1q2.35 0 4.525.775t3.975 2.25V2H22v5h-5V5.5h2.9q-1.65-1.45-3.675-2.225T12 2.5q-2.95 0-5.213 1.2T3.575 7H2Zm12.525 10Z"/>`),
		g.Group(children),
	)
}
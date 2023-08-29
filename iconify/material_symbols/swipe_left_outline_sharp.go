package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeLeftOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7V2h1.5v2.025q1.8-1.475 3.975-2.25T12 1q3.65 0 6.45 1.675T22 7h-1.575q-.95-2.1-3.213-3.3T12 2.5q-2.2 0-4.225.775T4.1 5.5H7V7H2Zm8.575 15L4.6 16l1.575-1.625l2.825.8V6h2v11.825l-2.425-.675l2.85 2.85H18v-7h2v9h-9.425ZM12 15v-5h2v5h-2Zm3 0v-4h2v4h-2Zm-.475 2Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GestureSelectSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12v-2h2v2H1Zm20 0v-2h2v2h-2ZM1 8V6h2v2H1Zm20 0V6h2v2h-2ZM1 4V2h2v2H1Zm4 8v-2h2v2H5Zm12 0v-2h2v2h-2Zm4-8V2h2v2h-2ZM5 4V2h2v2H5Zm4 0V2h2v2H9Zm4 0V2h2v2h-2Zm4 0V2h2v2h-2Zm-5.175 19q-.6 0-1.15-.225t-.975-.65L4.6 17l1.6-1.625l2.8.8V7h2v8h1v-4h2v4h1v-3h2v3h1v-1h2v5q0 1.65-1.175 2.825T16 23h-4.175Z"/>`),
		g.Group(children),
	)
}
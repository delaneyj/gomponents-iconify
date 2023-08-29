package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 6V3.55L3.55 9.5H6V11H1V6h1.5v2.45L8.45 2.5H6V1h5v5H9.5Zm5.325 17q-.6 0-1.15-.225t-.975-.65L7.6 17l.75-.775q.4-.4.938-.537t1.062.012l1.65.475V8q0-.425.288-.713T13 7q.425 0 .713.288T14 8v7h1v-3q0-.425.288-.713T16 11q.425 0 .713.288T17 12v3h1v-2q0-.425.288-.713T19 12q.425 0 .713.288T20 13v2h1q0-.425.288-.713T22 14q.425 0 .713.288T23 15v4q0 1.65-1.175 2.825T19 23h-4.175Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.825 22q-.6 0-1.15-.225t-.975-.65L4.6 16l.75-.775q.4-.4.938-.537t1.062.012l1.65.475V7q0-.425.288-.713T10 6q.425 0 .713.288T11 7v10.825l-2.425-.675l2.55 2.55q.125.125.313.213t.387.087H16q.825 0 1.413-.588T18 18v-4q0-.425.288-.713T19 13q.425 0 .713.288T20 14v4q0 1.65-1.175 2.825T16 22h-4.175ZM12 15v-4q0-.425.288-.713T13 10q.425 0 .713.288T14 11v4h-2Zm3 0v-3q0-.425.288-.713T16 11q.425 0 .713.288T17 12v3h-2ZM2 7V2h1.5v2.025q1.8-1.475 3.975-2.25T12 1q3.65 0 6.45 1.675T22 7h-1.575q-.95-2.1-3.213-3.3T12 2.5q-2.2 0-4.225.775T4.1 5.5H7V7H2Zm12.5 10Z"/>`),
		g.Group(children),
	)
}
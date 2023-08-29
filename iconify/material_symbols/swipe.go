package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swipe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 7h-5V5.5h2.9q-1.65-1.45-3.675-2.225T12 2.5q-2.2 0-4.225.775T4.1 5.5H7V7H2V2h1.5v2.025q1.8-1.475 3.975-2.25T12 1q2.35 0 4.525.775t3.975 2.25V2H22v5ZM11.825 22q-.6 0-1.15-.225t-.975-.65L4.6 16l.75-.775q.4-.4.938-.537t1.062.012l1.65.475V7q0-.425.288-.713T10 6q.425 0 .713.288T11 7v7h1v-3q0-.425.288-.713T13 10q.425 0 .713.288T14 11v3h1v-2q0-.425.288-.713T16 11q.425 0 .713.288T17 12v2h1q0-.425.288-.713T19 13q.425 0 .713.288T20 14v4q0 1.65-1.175 2.825T16 22h-4.175Z"/>`),
		g.Group(children),
	)
}
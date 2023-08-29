package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AvgTimeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 3q-.425 0-.713-.288T9 2q0-.425.288-.713T10 1h4q.425 0 .713.288T15 2q0 .425-.288.713T14 3h-4Zm0 11.75l-1.1-2.2q-.125-.275-.375-.413T8 12H3.05q.375-3.375 2.925-5.688T12 4q1.55 0 2.975.5t2.675 1.45l.7-.7q.275-.275.7-.288t.7.288q.275.275.275.7t-.275.7l-.7.7q.8 1.05 1.275 2.213T20.95 12h-4.325L14.9 8.55q-.275-.575-.9-.575t-.9.575l-3.1 6.2ZM12 22q-3.475 0-6.025-2.313T3.05 14h4.325L9.1 17.45q.275.575.9.575t.9-.575l3.1-6.2l1.1 2.2q.125.275.375.413T16 14h4.95q-.375 3.375-2.913 5.687T12 22Z"/>`),
		g.Group(children),
	)
}
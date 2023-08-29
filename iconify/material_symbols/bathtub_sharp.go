package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BathtubSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9q-.825 0-1.413-.588T5 7q0-.825.588-1.413T7 5q.825 0 1.413.588T9 7q0 .825-.588 1.413T7 9ZM4 22v-1H2v-8h3v-.75q0-.95.65-1.6t1.6-.65q.5 0 .925.2t.775.55l1.4 1.55q.2.2.388.375t.412.325H18V4.85q0-.35-.25-.6t-.6-.25q-.15 0-.288.063t-.262.187L15.35 5.5q.125.425.05.838t-.3.762l-2.75-2.8q.35-.225.75-.288t.8.088l1.25-1.25q.4-.4.913-.625T17.15 2q1.2 0 2.025.825T20 4.85V13h2v8h-2v1H4Z"/>`),
		g.Group(children),
	)
}
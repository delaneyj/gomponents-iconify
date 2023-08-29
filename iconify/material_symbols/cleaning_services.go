package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleaningServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23v-7q0-2.075 1.463-3.538T8 11h1V3q0-.825.588-1.413T11 1h2q.825 0 1.413.588T15 3v8h1q2.075 0 3.538 1.463T21 16v7H3Zm2-2h2v-3q0-.425.288-.713T8 17q.425 0 .713.288T9 18v3h2v-3q0-.425.288-.713T12 17q.425 0 .713.288T13 18v3h2v-3q0-.425.288-.713T16 17q.425 0 .713.288T17 18v3h2v-5q0-1.25-.875-2.125T16 13H8q-1.25 0-2.125.875T5 16v5Zm8-10V3h-2v8h2Z"/>`),
		g.Group(children),
	)
}
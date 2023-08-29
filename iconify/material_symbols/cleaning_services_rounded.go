package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CleaningServicesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-5q0-2.075 1.463-3.538T8 11h1V4q0-1.25.875-2.125T12 1q1.25 0 2.125.875T15 4v7h1q2.075 0 3.538 1.463T21 16v5q0 .825-.588 1.413T19 23H5q-.825 0-1.413-.588T3 21Zm2 0h2v-3q0-.425.288-.713T8 17q.425 0 .713.288T9 18v3h2v-3q0-.425.288-.713T12 17q.425 0 .713.288T13 18v3h2v-3q0-.425.288-.713T16 17q.425 0 .713.288T17 18v3h2v-5q0-1.25-.875-2.125T16 13H8q-1.25 0-2.125.875T5 16v5Z"/>`),
		g.Group(children),
	)
}
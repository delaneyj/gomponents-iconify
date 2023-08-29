package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalCafeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20q0-.425.288-.713T5 19h14q.425 0 .713.288T20 20q0 .425-.288.713T19 21H5Zm3-4q-1.65 0-2.825-1.175T4 13V5q0-.825.588-1.413T6 3h14q.825 0 1.413.588T22 5v3q0 .825-.588 1.413T20 10h-2v3q0 1.65-1.175 2.825T14 17H8Zm10-9h2V5h-2v3Z"/>`),
		g.Group(children),
	)
}
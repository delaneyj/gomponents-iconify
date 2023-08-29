package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CountertopsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.425 0-.713-.288T4 19v-7H3q-.425 0-.713-.288T2 11q0-.425.288-.713T3 10h3q-.825 0-1.413-.588T4 8V5q0-.425.288-.713T5 4h4q.425 0 .713.288T10 5v3q0 .825-.588 1.413T8 10h8V7q0-.425-.288-.713T15 6q-.275 0-.5.138t-.425.412q-.175.25-.387.35t-.413.1q-.65 0-.912-.488t.062-1.037q.4-.65 1.075-1.062T15 4q1.25 0 2.125.875T18 7v3h3q.425 0 .713.288T22 11q0 .425-.288.713T21 12h-1v7q0 .425-.288.713T19 20H5Zm6-2h2v-6h-2v6Z"/>`),
		g.Group(children),
	)
}
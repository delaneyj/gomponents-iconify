package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18v2q0 .425-.288.713T6 21q-.425 0-.713-.288T5 20v-6q0-.825.588-1.413T7 12h1v-2H7q-.825 0-1.413-.588T5 8V5q0-.825.588-1.413T7 3h10q.825 0 1.413.588T19 5v3q0 .825-.588 1.413T17 10h-1v2h1q.825 0 1.413.588T19 14v6q0 .425-.288.713T18 21q-.425 0-.713-.288T17 20v-2H7Zm3-6h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}
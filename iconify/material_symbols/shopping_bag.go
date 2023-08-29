package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V8q0-.825.588-1.413T6 6h2q0-1.65 1.175-2.825T12 2q1.65 0 2.825 1.175T16 6h2q.825 0 1.413.588T20 8v12q0 .825-.588 1.413T18 22H6Zm4-16h4q0-.825-.588-1.413T12 4q-.825 0-1.413.588T10 6Zm5 5q.425 0 .713-.288T16 10V8h-2v2q0 .425.288.713T15 11Zm-6 0q.425 0 .713-.288T10 10V8H8v2q0 .425.288.713T9 11Z"/>`),
		g.Group(children),
	)
}
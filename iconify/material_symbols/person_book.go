package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16q-1.725 0-3.225.525T6 18v2h12v-2q-1.275-.95-2.775-1.475T12 16Zm-6 6q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Zm6-8q1.45 0 2.475-1.025T15.5 10.5q0-1.45-1.025-2.475T12 7q-1.45 0-2.475 1.025T8.5 10.5q0 1.45 1.025 2.475T12 14Z"/>`),
		g.Group(children),
	)
}
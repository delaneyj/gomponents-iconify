package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18q-2.925 0-4.963-2.038T4 11V5q0-.825.588-1.413T6 3h12.5q1.45 0 2.475 1.025T22 6.5q0 1.45-1.025 2.475T18.5 10H18v1q0 2.925-2.038 4.963T11 18ZM6 8h10V5H6v3Zm12 0h.5q.625 0 1.063-.438T20 6.5q0-.625-.438-1.063T18.5 5H18v3ZM4 21v-2h16v2H4Z"/>`),
		g.Group(children),
	)
}
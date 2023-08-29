package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiAlertOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-9l2.45-7H8V4h3.3q-.3.975-.3 2t.3 2H5.85L4.8 11h8.3q.975.975 2.25 1.488T18 13q.5 0 .975-.063T20 12.75V22h-3v-2H5v2H2Zm2-9v5v-5Zm2.5 4q.625 0 1.063-.438T8 15.5q0-.625-.438-1.063T6.5 14q-.625 0-1.063.438T5 15.5q0 .625.438 1.063T6.5 17Zm9 0q.625 0 1.063-.438T17 15.5q0-.625-.438-1.063T15.5 14q-.625 0-1.063.438T14 15.5q0 .625.438 1.063T15.5 17Zm2.5-6q-2.075 0-3.538-1.463T13 6q0-2.075 1.463-3.538T18 1q2.075 0 3.538 1.463T23 6q0 2.075-1.463 3.538T18 11Zm-.5-4h1V3h-1v4ZM4 18h14v-5H4v5Zm13.5-9h1V8h-1v1Z"/>`),
		g.Group(children),
	)
}
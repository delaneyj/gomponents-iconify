package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecimalIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 22l-1.4-1.4l1.575-1.6H12v-2h6.175L16.6 15.4L18 14l4 4l-4 4ZM2 13v-3h3v3H2Zm7.5 0q-1.45 0-2.475-1.025T6 9.5v-4q0-1.45 1.025-2.475T9.5 2q1.45 0 2.475 1.025T13 5.5v4q0 1.45-1.025 2.475T9.5 13Zm9 0q-1.45 0-2.475-1.025T15 9.5v-4q0-1.45 1.025-2.475T18.5 2q1.45 0 2.475 1.025T22 5.5v4q0 1.45-1.025 2.475T18.5 13Zm-9-2q.625 0 1.063-.438T11 9.5v-4q0-.625-.438-1.063T9.5 4q-.625 0-1.063.438T8 5.5v4q0 .625.438 1.063T9.5 11Zm9 0q.625 0 1.063-.438T20 9.5v-4q0-.625-.438-1.063T18.5 4q-.625 0-1.063.438T17 5.5v4q0 .625.438 1.063T18.5 11Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecimalDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 22l-4-4l4-4l1.4 1.4l-1.575 1.6H22v2h-6.175l1.575 1.6L16 22ZM2 13v-3h3v3H2Zm7.5 0q-1.45 0-2.475-1.025T6 9.5v-4q0-1.45 1.025-2.475T9.5 2q1.45 0 2.475 1.025T13 5.5v4q0 1.45-1.025 2.475T9.5 13Zm0-2q.625 0 1.063-.438T11 9.5v-4q0-.625-.438-1.063T9.5 4q-.625 0-1.063.438T8 5.5v4q0 .625.438 1.063T9.5 11Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitEnterexitRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.15 15h3.35q.625 0 1.063.438T16 16.5q0 .625-.438 1.063T14.5 18H7q-.425 0-.713-.288T6 17V9.5q0-.625.438-1.063T7.5 8q.625 0 1.063.438T9 9.5v3.25l5.7-5.7q.45-.45 1.1-.45t1.1.45q.45.45.45 1.1t-.45 1.1L11.15 15Z"/>`),
		g.Group(children),
	)
}
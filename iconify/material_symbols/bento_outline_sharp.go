package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BentoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19V5h20v14H2Zm12-8h6V7h-6v4ZM4 17h8V7H4v10Zm4-3.5q-.625 0-1.063-.438T6.5 12q0-.625.438-1.063T8 10.5q.625 0 1.063.438T9.5 12q0 .625-.438 1.063T8 13.5Zm6 3.5h6v-4h-6v4Z"/>`),
		g.Group(children),
	)
}
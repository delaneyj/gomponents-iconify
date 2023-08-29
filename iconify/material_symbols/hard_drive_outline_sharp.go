package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17h16v-6H4v6Zm13-1.5q.625 0 1.063-.438T18.5 14q0-.625-.438-1.063T17 12.5q-.625 0-1.063.438T15.5 14q0 .625.438 1.063T17 15.5ZM22 9h-2.825l-2-2H6.825l-2 2H2l4-4h12l4 4ZM2 19V9h20v10H2Z"/>`),
		g.Group(children),
	)
}
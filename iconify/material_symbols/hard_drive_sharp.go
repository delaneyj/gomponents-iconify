package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16q.625 0 1.063-.438T18.5 14.5q0-.625-.438-1.063T17 13q-.625 0-1.063.438T15.5 14.5q0 .625.438 1.063T17 16Zm5-7H2l4-4h12l4 4ZM2 19v-8h20v8H2Z"/>`),
		g.Group(children),
	)
}
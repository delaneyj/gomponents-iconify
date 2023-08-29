package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 21q-.625 0-1.063-.438T7 19.5V13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h4V4.5q0-.625.438-1.063T8.5 3q.625 0 1.063.438T10 4.5V11h4V7.5q0-.625.438-1.063T15.5 6q.625 0 1.063.438T17 7.5V11h4q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-4v3.5q0 .625-.438 1.063T15.5 18q-.625 0-1.063-.438T14 16.5V13h-4v6.5q0 .625-.438 1.063T8.5 21Z"/>`),
		g.Group(children),
	)
}
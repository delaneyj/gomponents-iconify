package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-4H7.5q-.625 0-1.063-.438T6 15.5q0-.625.438-1.063T7.5 14H11v-4H4.5q-.625 0-1.063-.438T3 8.5q0-.625.438-1.063T4.5 7H11V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v4h6.5q.625 0 1.063.438T21 8.5q0 .625-.438 1.063T19.5 10H13v4h3.5q.625 0 1.063.438T18 15.5q0 .625-.438 1.063T16.5 17H13v4q0 .425-.288.713T12 22q-.425 0-.713-.288T11 21Z"/>`),
		g.Group(children),
	)
}
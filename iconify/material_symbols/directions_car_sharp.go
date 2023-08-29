package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsCarSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19v2H3v-9l2.45-7h13.1L21 12v9h-3v-2H6Zm-.2-9h12.4l-1.05-3H6.85L5.8 10Zm1.7 6q.625 0 1.063-.438T9 14.5q0-.625-.438-1.063T7.5 13q-.625 0-1.063.438T6 14.5q0 .625.438 1.063T7.5 16Zm9 0q.625 0 1.063-.438T18 14.5q0-.625-.438-1.063T16.5 13q-.625 0-1.063.438T15 14.5q0 .625.438 1.063T16.5 16Z"/>`),
		g.Group(children),
	)
}
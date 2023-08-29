package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackgroundDotLargeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 17q.625 0 1.063-.438T10 15.5q0-.625-.438-1.063T8.5 14q-.625 0-1.063.438T7 15.5q0 .625.438 1.063T8.5 17Zm0-7q.625 0 1.063-.438T10 8.5q0-.625-.438-1.063T8.5 7q-.625 0-1.063.438T7 8.5q0 .625.438 1.063T8.5 10Zm7 7q.625 0 1.063-.438T17 15.5q0-.625-.438-1.063T15.5 14q-.625 0-1.063.438T14 15.5q0 .625.438 1.063T15.5 17Zm0-7q.625 0 1.063-.438T17 8.5q0-.625-.438-1.063T15.5 7q-.625 0-1.063.438T14 8.5q0 .625.438 1.063T15.5 10ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}
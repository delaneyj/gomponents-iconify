package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CasinoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 18q.625 0 1.063-.438T9 16.5q0-.625-.438-1.063T7.5 15q-.625 0-1.063.438T6 16.5q0 .625.438 1.063T7.5 18Zm0-9q.625 0 1.063-.438T9 7.5q0-.625-.438-1.063T7.5 6q-.625 0-1.063.438T6 7.5q0 .625.438 1.063T7.5 9Zm4.5 4.5q.625 0 1.063-.438T13.5 12q0-.625-.438-1.063T12 10.5q-.625 0-1.063.438T10.5 12q0 .625.438 1.063T12 13.5Zm4.5 4.5q.625 0 1.063-.438T18 16.5q0-.625-.438-1.063T16.5 15q-.625 0-1.063.438T15 16.5q0 .625.438 1.063T16.5 18Zm0-9q.625 0 1.063-.438T18 7.5q0-.625-.438-1.063T16.5 6q-.625 0-1.063.438T15 7.5q0 .625.438 1.063T16.5 9ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}
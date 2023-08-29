package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsCricketOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 17.7L1.625 7.825l4.2-4.2L15.7 13.5l-4.2 4.2Zm0-2.8l1.4-1.4l-7.1-7.1l-1.4 1.4l7.1 7.1Zm7.1 7.1l-4.25-4.25l1.4-1.4L20 20.6L18.6 22Zm-.1-13q-1.45 0-2.475-1.025T15 5.5q0-1.45 1.025-2.475T18.5 2q1.45 0 2.475 1.025T22 5.5q0 1.45-1.025 2.475T18.5 9Zm0-2q.625 0 1.063-.438T20 5.5q0-.625-.438-1.063T18.5 4q-.625 0-1.063.438T17 5.5q0 .625.438 1.063T18.5 7Zm0-1.5Zm-9.85 5.15Z"/>`),
		g.Group(children),
	)
}
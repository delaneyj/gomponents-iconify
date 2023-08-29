package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallArtSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V6h6l4-4l4 4h6v16H2Zm4-4h12l-3.75-5l-3 4L9 14l-3 4Zm11.5-5q.625 0 1.063-.438T19 11.5q0-.625-.438-1.063T17.5 10q-.625 0-1.063.438T16 11.5q0 .625.438 1.063T17.5 13Zm-7.4-7h3.8L12 4.1L10.1 6Z"/>`),
		g.Group(children),
	)
}
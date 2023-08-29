package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallArt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V8q0-.825.588-1.413T4 6h4l4-4l4 4h4q.825 0 1.413.588T22 8v12q0 .825-.588 1.413T20 22H4Zm2-4h12l-3.75-5l-3 4L9 14l-3 4Zm11.5-5q.625 0 1.063-.438T19 11.5q0-.625-.438-1.063T17.5 10q-.625 0-1.063.438T16 11.5q0 .625.438 1.063T17.5 13Zm-7.4-7h3.8L12 4.1L10.1 6Z"/>`),
		g.Group(children),
	)
}
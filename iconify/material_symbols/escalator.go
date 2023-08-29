package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escalator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18h3.3l5-9H17q.625 0 1.063-.438T18.5 7.5q0-.625-.438-1.063T17 6h-3.3l-5 9H7q-.625 0-1.062.438T5.5 16.5q0 .625.438 1.063T7 18Zm-2 3q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
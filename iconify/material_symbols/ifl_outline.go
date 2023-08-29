package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IflOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 17.5q.625 0 1.063-.438T17.5 16q0-.625-.438-1.063T16 14.5q-.625 0-1.063.438T14.5 16q0 .625.438 1.063T16 17.5Zm-4-4q.625 0 1.063-.438T13.5 12q0-.625-.438-1.063T12 10.5q-.625 0-1.063.438T10.5 12q0 .625.438 1.063T12 13.5Zm-4-4q.625 0 1.063-.438T9.5 8q0-.625-.438-1.063T8 6.5q-.625 0-1.063.438T6.5 8q0 .625.438 1.063T8 9.5ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
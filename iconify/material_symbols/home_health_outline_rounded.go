package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHealthOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 14.5v1q0 .625.438 1.063T12 17q.625 0 1.063-.438T13.5 15.5v-1h1q.625 0 1.063-.438T16 13q0-.625-.438-1.063T14.5 11.5h-1v-1q0-.625-.438-1.063T12 9q-.625 0-1.063.438T10.5 10.5v1h-1q-.625 0-1.063.438T8 13q0 .625.438 1.063T9.5 14.5h1ZM6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Zm0-2h12v-9l-6-4.5L6 10v9Zm6-6.75Z"/>`),
		g.Group(children),
	)
}
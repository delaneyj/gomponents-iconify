package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17.5V8q0-2.425 2.125-3.175T11 4l.75-1.5H7V1h10v1.5h-3.25L13 4q2.975.075 4.988.813T20 8v9.5q0 1.475-1.012 2.488T16.5 21l1.5 1.5v.5h-2l-2-2h-4l-2 2H6v-.5L7.5 21q-1.475 0-2.488-1.012T4 17.5Zm8 .5q.625 0 1.063-.438T13.5 16.5q0-.625-.438-1.063T12 15q-.625 0-1.063.438T10.5 16.5q0 .625.438 1.063T12 18Zm-6-6h12V9H6v3Z"/>`),
		g.Group(children),
	)
}
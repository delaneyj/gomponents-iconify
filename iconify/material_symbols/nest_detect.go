package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestDetect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 22q-1.45 0-2.475-1.025T6 18.5V4q0-.825.588-1.413T8 2h3q.825 0 1.413.588T13 4v14.5q0 1.45-1.025 2.475T9.5 22Zm0-2q.625 0 1.063-.438T11 18.5q0-.625-.438-1.063T9.5 17q-.625 0-1.063.438T8 18.5q0 .625.438 1.063T9.5 20Zm7.5-2q-.825 0-1.413-.588T15 16V8q0-.825.588-1.413T17 6h1q.825 0 1.413.588T20 8v8q0 .825-.588 1.413T18 18h-1Z"/>`),
		g.Group(children),
	)
}
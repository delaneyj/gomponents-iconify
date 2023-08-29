package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 21l-7-7l-2.35 2.35q.2.375.275.8T10 18q0 1.65-1.175 2.825T6 22q-1.65 0-2.825-1.175T2 18q0-1.65 1.175-2.825T6 14q.425 0 .85.075t.8.275L10 12L7.65 9.65q-.375.2-.8.275T6 10q-1.65 0-2.825-1.175T2 6q0-1.65 1.175-2.825T6 2q1.65 0 2.825 1.175T10 6q0 .425-.075.85t-.275.8L22 20v1h-3Zm-4-10l-2-2l6-6h3v1l-7 7ZM6 8q.825 0 1.413-.588T8 6q0-.825-.588-1.413T6 4q-.825 0-1.413.588T4 6q0 .825.588 1.413T6 8Zm6 4.5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T12 11.5q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15ZM6 20q.825 0 1.413-.588T8 18q0-.825-.588-1.413T6 16q-.825 0-1.413.588T4 18q0 .825.588 1.413T6 20Z"/>`),
		g.Group(children),
	)
}
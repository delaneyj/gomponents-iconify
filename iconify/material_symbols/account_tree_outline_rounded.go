package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountTreeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21q-.825 0-1.413-.588T15 19v-1h-2q-.825 0-1.413-.588T11 16V8H9v1q0 .825-.588 1.413T7 11H4q-.825 0-1.413-.588T2 9V5q0-.825.588-1.413T4 3h3q.825 0 1.413.588T9 5v1h6V5q0-.825.588-1.413T17 3h3q.825 0 1.413.588T22 5v4q0 .825-.588 1.413T20 11h-3q-.825 0-1.413-.588T15 9V8h-2v8h2v-1q0-.825.588-1.413T17 13h3q.825 0 1.413.588T22 15v4q0 .825-.588 1.413T20 21h-3ZM4 5v4v-4Zm13 10v4v-4Zm0-10v4v-4Zm0 4h3V5h-3v4Zm0 10h3v-4h-3v4ZM4 9h3V5H4v4Z"/>`),
		g.Group(children),
	)
}
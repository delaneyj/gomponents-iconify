package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LtePlusMobiledataBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16h4v-2H6V8H4v8Zm5.5 0h2v-6H13V8H8v2h1.5v6Zm4.5 0h4v-2h-2v-1h1.5v-2H16v-1h2V8h-4v8ZM3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v6.25h-1.25V10h-1.5v1.25H19v1.5h1.25V14h1.5v-1.25H23V19q0 .825-.588 1.413T21 21H3Z"/>`),
		g.Group(children),
	)
}
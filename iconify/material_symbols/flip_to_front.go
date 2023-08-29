package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19h2v2Zm-2-4v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm4 12v-2h2v2H7Zm2-4q-.825 0-1.413-.588T7 15V5q0-.825.588-1.413T9 3h10q.825 0 1.413.588T21 5v10q0 .825-.588 1.413T19 17H9Zm0-2h10V5H9v10Zm2 6v-2h2v2h-2Zm4 0v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}
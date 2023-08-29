package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gradient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm6-10v2h2v-2h-2Zm-4 0v2h2v-2H7Zm2 2v2h2v-2H9Zm4 0v2h2v-2h-2Zm-8 0v2h2v-2H5Zm10-2v2h2v2h2v-2h-2v-2h-2Zm-8 4v2H5v2h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h-2v-2h-2v2h-2v-2h-2v2H9v-2H7Zm12-4v2v-2Zm0 4v2v-2Z"/>`),
		g.Group(children),
	)
}
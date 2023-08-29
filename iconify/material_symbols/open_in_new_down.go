package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInNewDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v7h-2V5H5v14h7v2H5Zm9 0v-2h3.6L8.3 9.7l1.4-1.4l9.3 9.275V14h2v7h-7Z"/>`),
		g.Group(children),
	)
}
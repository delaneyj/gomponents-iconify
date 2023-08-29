package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Output(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v2h-2V5H5v14h14v-2h2v2q0 .825-.588 1.413T19 21H5Zm12-4l-1.4-1.4l2.575-2.6H9v-2h9.175L15.6 8.4L17 7l5 5l-5 5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrokenImageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm1-8.425l4-4l4 4l4-4l1 1V5H5v6.575l1 1ZM5 19h14v-6.6l-1-1l-4 4l-4-4l-4 4l-1-1V19Zm0 0v-6.6v2V5v14Z"/>`),
		g.Group(children),
	)
}
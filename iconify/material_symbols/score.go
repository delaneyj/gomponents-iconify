package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Score(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7-9h1.5V6H12v6Zm3.5 0h1.7l-2-3l2-3h-1.7l-2 3l2 3ZM7 12h4v-1.5H8.5v-.75H11V6H7v1.5h2.5v.75H7V12Zm12 1v-2.5l-6 6l-4-4l-4 4V19l4-4l4 4l6-6Z"/>`),
		g.Group(children),
	)
}
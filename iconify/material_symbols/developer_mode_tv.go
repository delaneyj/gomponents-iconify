package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperModeTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2H4q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-4v2H8ZM4.8 11l2.6-2.6L6 7l-4 4l4 4l1.4-1.4L4.8 11Zm14.4 0l-2.6 2.6L18 15l4-4l-4-4l-1.4 1.4l2.6 2.6Z"/>`),
		g.Group(children),
	)
}
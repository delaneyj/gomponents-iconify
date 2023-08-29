package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkEraseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.4 16L13 14.6l2.6-2.6L13 9.4L14.4 8l2.6 2.6L19.6 8L21 9.4L18.4 12l2.6 2.6l-1.4 1.4l-2.6-2.6l-2.6 2.6ZM6 23q-.825 0-1.413-.588T4 21V3q0-.825.588-1.413T6 1h10q.825 0 1.413.588T18 3v4h-2V6H6v12h10v-1h2v4q0 .825-.588 1.413T16 23H6Zm0-3v1h10v-1H6ZM6 4h10V3H6v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}
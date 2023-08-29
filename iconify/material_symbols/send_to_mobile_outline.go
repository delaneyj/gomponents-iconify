package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendToMobileOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 16l-1.4-1.4l1.55-1.6H12v-2h5.15L15.6 9.4L17 8l4 4l-4 4ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v12h10v-1h2v4q0 .825-.588 1.413T17 23H7Zm0-3v1h10v-1H7ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}
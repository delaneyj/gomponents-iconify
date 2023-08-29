package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInPhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23q-.825 0-1.413-.588T5 21v-6h2v3h10V6H7v3H5V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm3-7l-1.4-1.4l1.55-1.6H2v-2h8.15L8.6 9.4L10 8l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}
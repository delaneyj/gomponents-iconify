package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluorescentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15.05v-6h14v6H5ZM11 5V2h2v3h-2Zm7.7 2.8l-1.4-1.4l1.8-1.8L20.5 6l-1.8 1.8ZM11 22v-3h2v3h-2Zm8.1-2.5l-1.8-1.8l1.4-1.4l1.8 1.8l-1.4 1.4ZM5.3 7.8L3.5 6l1.4-1.4l1.8 1.8l-1.4 1.4Zm-.4 11.7l-1.4-1.4l1.8-1.8l1.4 1.4l-1.8 1.8ZM7 13.05h10v-2H7v2Zm0 0v-2v2Z"/>`),
		g.Group(children),
	)
}
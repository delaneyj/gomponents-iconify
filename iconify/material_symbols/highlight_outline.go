package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.65 8L3.5 5.9l1.4-1.45L7.05 6.6L5.65 8ZM11 5V2h2v3h-2Zm7.4 3l-1.45-1.4l2.15-2.1l1.4 1.4L18.4 8ZM9 22v-5l-3-3V9h12v5l-3 3v5H9Zm2-2h2v-3.825l3-3V11H8v2.175l3 3V20Zm1-4.5Z"/>`),
		g.Group(children),
	)
}
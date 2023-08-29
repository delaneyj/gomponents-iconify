package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessAutoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.8 16h1.6l.8-2.3h3.65l.8 2.3h1.55l-3.4-9h-1.6l-3.4 9Zm2.85-3.6l1.3-3.75h.1l1.3 3.75h-2.7ZM12 23.3L8.65 20H4v-4.65L.7 12L4 8.65V4h4.65L12 .7L15.35 4H20v4.65L23.3 12L20 15.35V20h-4.65L12 23.3ZM12 12Zm0 8.5l2.5-2.5H18v-3.5l2.5-2.5L18 9.5V6h-3.5L12 3.5L9.5 6H6v3.5L3.5 12L6 14.5V18h3.5l2.5 2.5Z"/>`),
		g.Group(children),
	)
}
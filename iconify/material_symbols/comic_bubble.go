package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComicBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20.9L8.1 18H4v-4.1L1.1 11L4 8.1V4h4.1L11 1.1L13.9 4H18v4.1l2.9 2.9l-2.9 2.9l2.875 5.65q.175.325.1.638t-.275.512q-.2.2-.512.275t-.638-.1L13.9 18L11 20.9Z"/>`),
		g.Group(children),
	)
}
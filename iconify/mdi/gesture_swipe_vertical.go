package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GestureSwipeVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3L1 6h2v3H1l3 3l3-3H5V6h2L4 3m7 5a1 1 0 0 0-1 1v10l-3.2-1.72h-.22c-.28 0-.55.11-.74.32l-.74.77l4.9 4.2c.26.28.62.43 1 .43h6.5a1.5 1.5 0 0 0 1.5-1.5v-4.36c0-.58-.32-1.11-.85-1.35l-4.94-2.19l-1.21-.13V9a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}
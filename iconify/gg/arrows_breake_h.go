package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsBreakeH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.243 7h2v4h.005v2h-.005v4h-2v-4H4.828l1.829 1.828l-1.414 1.415L1 12l4.243-4.243l1.414 1.415L4.828 11h4.415V7Zm6.01 0h-2v4h-.005v2h.005v4h2v-4h4.414l-1.829 1.829l1.415 1.414L23.495 12l-4.242-4.243l-1.415 1.415L19.668 11h-4.414V7Z"/>`),
		g.Group(children),
	)
}
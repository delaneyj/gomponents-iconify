package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardCloseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l4-4H8M4 3c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2H4m0 2h16v10H4V5m1 1v2h2V6H5m3 0v2h2V6H8m3 0v2h2V6h-2m3 0v2h2V6h-2m3 0v2h2V6h-2M5 9v2h2V9H5m3 0v2h2V9H8m3 0v2h2V9h-2m3 0v2h2V9h-2m3 0v2h2V9h-2m-9 3v2h8v-2H8Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapHorizontalVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 6l4 4V7h8a2 2 0 0 1 2 2a2 2 0 0 1-2 2H8a4 4 0 0 0-4 4a4 4 0 0 0 4 4h8v3l4-4l-4-4v3H8a2 2 0 0 1-2-2a2 2 0 0 1 2-2h8a4 4 0 0 0 4-4a4 4 0 0 0-4-4H8V2L4 6Z"/>`),
		g.Group(children),
	)
}
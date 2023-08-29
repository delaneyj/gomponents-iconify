package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChoppingBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 8h28.011A1.99 1.99 0 0 1 44 10v28c0 1.105-.883 2-1.987 2H14c-4 0-10-2-10-16S11 8 14 8Zm-2 12v8"/>`),
		g.Group(children),
	)
}
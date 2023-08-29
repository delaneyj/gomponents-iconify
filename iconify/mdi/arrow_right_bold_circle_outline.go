package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightBoldCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 12l-5 5v-3H8v-4h4V7l5 5M2 12A10 10 0 0 1 12 2a10 10 0 0 1 10 10a10 10 0 0 1-10 10A10 10 0 0 1 2 12m2 0a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8a8 8 0 0 0-8 8Z"/>`),
		g.Group(children),
	)
}
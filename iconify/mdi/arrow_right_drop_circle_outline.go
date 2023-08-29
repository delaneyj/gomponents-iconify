package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightDropCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2A10 10 0 0 0 2 12m2 0a8 8 0 0 1 8-8a8 8 0 0 1 8 8a8 8 0 0 1-8 8a8 8 0 0 1-8-8m6 5l5-5l-5-5v10Z"/>`),
		g.Group(children),
	)
}
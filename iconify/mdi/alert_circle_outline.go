package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2v2h-2v-2m0-8h2v6h-2V7m1-5C6.47 2 2 6.5 2 12a10 10 0 0 0 10 10a10 10 0 0 0 10-10A10 10 0 0 0 12 2m0 18a8 8 0 0 1-8-8a8 8 0 0 1 8-8a8 8 0 0 1 8 8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}
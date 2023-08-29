package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForumMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6h-2v6.1c1.2.2 2.2.7 3 1.4V7c0-.5-.5-1-1-1M6 17c0 .5.5 1 1 1h5c0-1.1.3-2.1.8-3H6v2M16 2H3c-.5 0-1 .5-1 1v14l4-4h8.7c.7-.5 1.5-.8 2.3-.9V3c0-.5-.5-1-1-1m6 15v2h-8v-2h8Z"/>`),
		g.Group(children),
	)
}
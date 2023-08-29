package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6H4c-.55 0-1 .45-1 1v1c0 .55.45 1 1 1h1v9h14V9h1c.55 0 1-.45 1-1V7c0-.55-.45-1-1-1m-3 10H7V9h10v7Z"/>`),
		g.Group(children),
	)
}
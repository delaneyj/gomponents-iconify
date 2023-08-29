package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumericFourBoxMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5v16h16v2H3a2 2 0 0 1-2-2V5h2m18-4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h14m-6 14h2V5h-2v4h-2V5h-2v6h4v4Z"/>`),
		g.Group(children),
	)
}
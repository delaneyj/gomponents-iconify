package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Advertisements(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 7c-1.1 0-2 .9-2 2v2c0 1.1.9 2 2 2h2v2h-4v2h4c1.1 0 2-.9 2-2v-2c0-1.1-.9-2-2-2h-2V9h4V7h-4M9 7v10h4c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2H9m2 2h2v6h-2V9M3 7c-1.1 0-2 .9-2 2v8h2v-4h2v4h2V9c0-1.1-.9-2-2-2H3m0 2h2v2H3V9Z"/>`),
		g.Group(children),
	)
}
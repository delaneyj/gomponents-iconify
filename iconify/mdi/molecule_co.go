package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoleculeCo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7c-1.1 0-2 .9-2 2v6a2 2 0 0 0 2 2h3v-2H8V9h3V7H8m6 0c-1.1 0-2 .9-2 2v6a2 2 0 0 0 2 2h2c1.11 0 2-.89 2-2V9a2 2 0 0 0-2-2h-2m0 2h2v6h-2V9"/>`),
		g.Group(children),
	)
}
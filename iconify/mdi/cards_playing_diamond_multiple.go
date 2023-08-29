package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsPlayingDiamondMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4v18c0 1.1.9 2 2 2h12v-2H5V4H3m16-4H9C7.9 0 7 .9 7 2v16c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V2c0-1.1-.9-2-2-2m-5 15l-3-5l3-5l3 5l-3 5Z"/>`),
		g.Group(children),
	)
}
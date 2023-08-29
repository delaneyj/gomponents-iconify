package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseRollingLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M102 88v96a6 6 0 0 1-12 0V88a6 6 0 0 1 12 0Zm26-6a6 6 0 0 0-6 6v96a6 6 0 0 0 12 0V88a6 6 0 0 0-6-6Zm32 0a6 6 0 0 0-6 6v96a6 6 0 0 0 12 0V88a6 6 0 0 0-6-6Zm46-18v144a14 14 0 0 1-14 14h-18v18a6 6 0 0 1-12 0v-18H94v18a6 6 0 0 1-12 0v-18H64a14 14 0 0 1-14-14V64a14 14 0 0 1 14-14h26V24a22 22 0 0 1 22-22h32a22 22 0 0 1 22 22v26h26a14 14 0 0 1 14 14ZM102 50h52V24a10 10 0 0 0-10-10h-32a10 10 0 0 0-10 10Zm92 14a2 2 0 0 0-2-2H64a2 2 0 0 0-2 2v144a2 2 0 0 0 2 2h128a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}
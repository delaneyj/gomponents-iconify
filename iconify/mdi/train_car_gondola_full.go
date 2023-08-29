package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarGondolaFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10c-1.96-1.81-5.26-3-9-3s-7.04 1.19-9 3H1v7h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1v-7h-2m0 5h-2v-2h-2v2h-2v-2h-2v2h-2v-2H9v2H7v-2H5v2H3v-3h18v3Z"/>`),
		g.Group(children),
	)
}
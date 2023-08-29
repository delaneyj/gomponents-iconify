package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarGondola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 10v7h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1v-7H1m20 5h-2v-2h-2v2h-2v-2h-2v2h-2v-2H9v2H7v-2H5v2H3v-3h18v3Z"/>`),
		g.Group(children),
	)
}
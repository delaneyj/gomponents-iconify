package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarContainer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 6v11h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1V6H1m20 9h-2V9h-2v6h-2V9h-2v6h-2V9H9v6H7V9H5v6H3V8h18v7Z"/>`),
		g.Group(children),
	)
}
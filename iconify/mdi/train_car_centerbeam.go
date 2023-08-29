package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarCenterbeam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 6v11h-1c0 1.11-.89 2-2 2s-2-.89-2-2H6c0 1.11-.89 2-2 2s-2-.89-2-2H1V6h2v9h18V6h2Z"/>`),
		g.Group(children),
	)
}
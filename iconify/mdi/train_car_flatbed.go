package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarFlatbed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 15v2h-1c0 1.11-.89 2-2 2s-2-.89-2-2H6c0 1.11-.89 2-2 2s-2-.89-2-2H1v-2h22Z"/>`),
		g.Group(children),
	)
}
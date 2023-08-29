package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarTank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 15v2h-1c0 1.11-.89 2-2 2s-2-.89-2-2H6c0 1.11-.89 2-2 2s-2-.89-2-2H1v-2h22m-2-7h-6V7H9v1H3c-1.1 0-2 .9-2 2v2a2 2 0 0 0 2 2h18c1.11 0 2-.89 2-2v-2a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}
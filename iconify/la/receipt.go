package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Receipt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 5.375V13H3v14h26V13h-4V5.375l-3 1.5l-2-1l-2 1l-2-1l-2 1l-2-1l-2 1zm5 2.75l2 1l2-1l2 1l2-1l2 1l1-.5V17H9V8.625l1 .5zM5 15h2v4h18v-4h2v10H5zm4 6v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2zm4 0v2h2v-2z"/>`),
		g.Group(children),
	)
}
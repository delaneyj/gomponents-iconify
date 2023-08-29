package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v14h18V5h-2v6.5l-3-2.25l-3 2.25V5H3Zm12 0v2.5l1-.75l1 .75V5h-2ZM5 11h6v2H5v-2Zm0 4h14v2H5v-2Z"/>`),
		g.Group(children),
	)
}
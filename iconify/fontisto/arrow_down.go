package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.6 8.034L1.634 8l10.633 10.608L22.901 8l.034.034v5.319L12.268 24L1.602 13.353z"/><path fill="currentColor" d="M14.044 0v18.666h-3.555V0z"/>`),
		g.Group(children),
	)
}
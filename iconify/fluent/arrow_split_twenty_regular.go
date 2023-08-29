package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSplitTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3a.5.5 0 0 1 .5.5V8H13a2 2 0 0 1 2 2v5.293l1.646-1.647a.5.5 0 0 1 .708.708l-2.5 2.5a.5.5 0 0 1-.708 0l-2.5-2.5a.5.5 0 0 1 .708-.708L14 15.293V10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v5.293l1.646-1.647a.5.5 0 0 1 .708.708l-2.5 2.5a.5.5 0 0 1-.708 0l-2.5-2.5a.5.5 0 0 1 .708-.708L5 15.293V10a2 2 0 0 1 2-2h2.5V3.5A.5.5 0 0 1 10 3Z"/>`),
		g.Group(children),
	)
}
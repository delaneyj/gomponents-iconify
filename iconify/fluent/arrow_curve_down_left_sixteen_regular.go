package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveDownLeftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.934 1.252a.5.5 0 0 1-.186.682C8.213 3.383 8 6.002 8 8v5.293l2.646-2.646a.5.5 0 0 1 .708.707l-3.5 3.5a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.707L7 13.293V8c0-2.002.187-5.183 3.252-6.934a.5.5 0 0 1 .682.186Z"/>`),
		g.Group(children),
	)
}
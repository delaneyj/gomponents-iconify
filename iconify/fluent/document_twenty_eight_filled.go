package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M14 2v8a2 2 0 0 0 2 2h7.999l.001.078V23.6a2.4 2.4 0 0 1-2.4 2.4H6.4A2.4 2.4 0 0 1 4 23.6V4.4A2.4 2.4 0 0 1 6.4 2H14Z"/><path d="M15.5 2.475V10a.5.5 0 0 0 .5.5h7.502a2.739 2.739 0 0 0-.307-.366l-7.431-7.431a2.401 2.401 0 0 0-.264-.228Z"/></g>`),
		g.Group(children),
	)
}
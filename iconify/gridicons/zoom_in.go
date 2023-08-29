package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.8 13.8c.7-1.1 1.2-2.4 1.2-3.8c0-3.9-3.1-7-7-7s-7 3.1-7 7s3.1 7 7 7c1.4 0 2.7-.4 3.8-1.2L19 21l2-2l-5.2-5.2zM10 15c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5z"/><path fill="currentColor" d="M11 7H9v2H7v2h2v2h2v-2h2V9h-2z"/>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 8h-.68a2.64 2.64 0 0 0-2.112 1.056l-4.416 5.888A2.64 2.64 0 0 1 7.68 16H7"/><path d="M8 8h.368c1 0 1.915.565 2.362 1.46l2.54 5.08A2.64 2.64 0 0 0 15.632 16H16"/></g>`),
		g.Group(children),
	)
}
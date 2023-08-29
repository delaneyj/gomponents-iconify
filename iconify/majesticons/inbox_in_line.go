package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxInLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.293 7.293a1 1 0 0 1 1.414 0L11 8.586V3a1 1 0 1 1 2 0v5.586l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 0-1.414z"/><path d="M9 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-3v2h3a1 1 0 0 1 1 1v7h-3a2 2 0 0 0-2 2h-4a2 2 0 0 0-2-2H5V6a1 1 0 0 1 1-1h3V3zm7 12h3v3a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-3h3a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2z"/></g>`),
		g.Group(children),
	)
}
package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 13h3a2 2 0 0 1 2 2h4a2 2 0 0 1 2-2h3V6a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v7zm14 2h-3a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2H5v3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-3zM3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6z"/></g>`),
		g.Group(children),
	)
}
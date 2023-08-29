package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M6 9a6 6 0 0 1 12 0v6a6 6 0 0 1-12 0V9Z"/><path stroke-linecap="round" d="M12 7v4"/></g>`),
		g.Group(children),
	)
}
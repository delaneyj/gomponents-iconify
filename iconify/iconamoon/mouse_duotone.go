package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 9a6 6 0 1 1 12 0v6a6 6 0 0 1-12 0V9Z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 9a6 6 0 0 1 12 0v6a6 6 0 0 1-12 0V9Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7v4"/></g>`),
		g.Group(children),
	)
}
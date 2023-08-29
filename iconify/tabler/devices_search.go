package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 13V9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2.5"/><path d="M18 8V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h7m4 0a3 3 0 1 0 6 0a3 3 0 1 0-6 0m5.2 2.2L22 22M16 9h2"/></g>`),
		g.Group(children),
	)
}
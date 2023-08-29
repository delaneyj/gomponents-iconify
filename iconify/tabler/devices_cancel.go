package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesCancel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 15.5V9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v3.5"/><path d="M18 8V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8m4 1a3 3 0 1 0 6 0a3 3 0 1 0-6 0m1 2l4-4m-5-8h2"/></g>`),
		g.Group(children),
	)
}
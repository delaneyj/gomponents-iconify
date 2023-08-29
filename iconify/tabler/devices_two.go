package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 15H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h6m3 0a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1zM7 19h3m7-11v.01"/><path d="M16 16a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-7-1v4"/></g>`),
		g.Group(children),
	)
}
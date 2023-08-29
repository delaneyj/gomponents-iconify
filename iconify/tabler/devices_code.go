package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 15.5V9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v4m0 6a1 1 0 0 1-1 1"/><path d="M18 8V5a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h7m9 3l2-2l-2-2m-3 0l-2 2l2 2M16 9h2"/></g>`),
		g.Group(children),
	)
}
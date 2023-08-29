package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17 15a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z" opacity=".16"/><circle cx="12" cy="15" r="5" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="7" r="3" fill="currentColor" opacity=".16"/><circle cx="12" cy="7" r="3" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/></g>`),
		g.Group(children),
	)
}
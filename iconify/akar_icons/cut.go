package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="18" r="3"/><path d="M15 15L7 3m2 12l3-4.5M17 3l-3 4.5"/><circle cx="17" cy="18" r="3"/></g>`),
		g.Group(children),
	)
}
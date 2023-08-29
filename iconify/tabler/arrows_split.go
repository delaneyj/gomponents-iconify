package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 17h-8l-3.5-5H3m18-5h-8l-3.495 5"/><path d="m18 10l3-3l-3-3m0 16l3-3l-3-3"/></g>`),
		g.Group(children),
	)
}
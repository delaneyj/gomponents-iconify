package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabasePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6c0 1.657 3.582 3 8 3s8-1.343 8-3s-3.582-3-8-3s-8 1.343-8 3"/><path d="M4 6v6c0 1.657 3.582 3 8 3c1.075 0 2.1-.08 3.037-.224M20 12V6"/><path d="M4 12v6c0 1.657 3.582 3 8 3c.166 0 .331-.002.495-.006M16 19h6m-3-3v6"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6c0 1.657 3.582 3 8 3s8-1.343 8-3s-3.582-3-8-3s-8 1.343-8 3"/><path d="M4 6v6c0 1.657 3.582 3 8 3c.537 0 1.062-.02 1.57-.058M20 13.5V6"/><path d="M4 12v6c0 1.657 3.582 3 8 3c.384 0 .762-.01 1.132-.03M22 22l-5-5m0 5l5-5"/></g>`),
		g.Group(children),
	)
}
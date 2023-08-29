package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3c0 2 1 3 4 3c2 0 3 1 3 3m11-6c0 2-1 3-4 3c-2 0-3 .333-3 3m-2 9c-1 0-4-3-4-6c0-2 1.333-3 4-3s4 1 4 3c0 3-3 6-4 6"/><path d="m15.185 14.889l.095-.18a4 4 0 1 1-6.56 0M17 3c0 1.333-.333 2.333-1 3M7 3c0 1.333.333 2.333 1 3M7 6c-2.667.667-4.333 1.667-5 3m15-3c2.667.667 4.333 1.667 5 3M8.5 10L7 9m8.5 1L17 9m-5 6h.01"/></g>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819L4 9Z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819L4 9Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8 11V8a4 4 0 1 1 8 0v3"/></g>`),
		g.Group(children),
	)
}
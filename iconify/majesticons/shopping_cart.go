package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M18 15H7L5.5 6H21l-3 9z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.5 3m0 0L7 15h11l3-9H5.5z"/><circle cx="8" cy="20" r="1" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="17" cy="20" r="1" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`),
		g.Group(children),
	)
}
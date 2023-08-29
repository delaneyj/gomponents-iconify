package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3h2l.5 3m0 0L7 15h11l3-9H5.5z"/><circle cx="8" cy="20" r="1"/><circle cx="17" cy="20" r="1"/></g>`),
		g.Group(children),
	)
}
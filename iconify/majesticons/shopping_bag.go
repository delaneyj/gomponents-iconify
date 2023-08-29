package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 7a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3H6zm6-3a3 3 0 0 0-3 3v4a1 1 0 1 1-2 0V7a5 5 0 0 1 10 0v4a1 1 0 1 1-2 0V7a3 3 0 0 0-3-3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
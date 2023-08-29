package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3 10a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-9zm3-1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1H6z"/><path d="M12 4a3 3 0 0 0-3 3v4a1 1 0 1 1-2 0V7a5 5 0 0 1 10 0v4a1 1 0 1 1-2 0V7a3 3 0 0 0-3-3z"/></g>`),
		g.Group(children),
	)
}
package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feShoppingBag0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feShoppingBag1" fill="currentColor" fill-rule="nonzero"><path id="feShoppingBag2" d="M8 8V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v3h3a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h3Zm-3 2v9h14v-9h-3.002v3H14v-3h-4v3H8v-3H5Zm5-2h4V5h-4v3Z"/></g></g>`),
		g.Group(children),
	)
}
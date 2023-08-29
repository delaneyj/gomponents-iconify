package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBagOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaShoppingBagOutline0"><g id="evaShoppingBagOutline1"><g id="evaShoppingBagOutline2" fill="currentColor"><path d="m20.12 6.71l-2.83-2.83A3 3 0 0 0 15.17 3H8.83a3 3 0 0 0-2.12.88L3.88 6.71A3 3 0 0 0 3 8.83V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.83a3 3 0 0 0-.88-2.12Zm-12-1.42A1.05 1.05 0 0 1 8.83 5h6.34a1.05 1.05 0 0 1 .71.29L17.59 7H6.41ZM18 19H6a1 1 0 0 1-1-1V9h14v9a1 1 0 0 1-1 1Z"/><path d="M15 11a1 1 0 0 0-1 1a2 2 0 0 1-4 0a1 1 0 0 0-2 0a4 4 0 0 0 8 0a1 1 0 0 0-1-1Z"/></g></g></g>`),
		g.Group(children),
	)
}
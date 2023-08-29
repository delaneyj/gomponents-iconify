package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m13-2a2 2 0 1 0 2 2"/><path d="M17 17H6V6m3.239-.769L20 6l-1 7h-2m-4 0H6M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
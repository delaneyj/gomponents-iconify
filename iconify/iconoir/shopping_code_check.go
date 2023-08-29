package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCodeCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 5v2m4-2v6m8-6v1M6 10v6m0 2.5v.5m4-.5v.5m0-5v2m4-3v2m0-10v5m4-1v4m-3 6l2 2l4-4"/>`),
		g.Group(children),
	)
}
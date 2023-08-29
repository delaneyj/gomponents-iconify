package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFortress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h1a1 1 0 0 0 1-1v-1h0a3 3 0 0 1 6 0m3 2h1a1 1 0 0 0 1-1V5l-3-2l-3 2v6h-4V5L7 3L4 5v15a1 1 0 0 0 1 1h2m8-2v1a1 1 0 0 0 1 1h2M7 7h0v.01M7 10h0v.01M7 13h0v.01M17 7h0v.01M17 10h0v.01M17 13h0v.01"/>`),
		g.Group(children),
	)
}
package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPointOfSale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5v6h14V2zm-2 4H7V4h10v2zm5 16H2v-3h20v3zM18 9H6l-4 9h20l-4-9zm-8 7H8v-1h2v1zm0-2H8v-1h2v1zm0-2H8v-1h2v1zm3 4h-2v-1h2v1zm0-2h-2v-1h2v1zm0-2h-2v-1h2v1zm3 4h-2v-1h2v1zm0-2h-2v-1h2v1zm0-2h-2v-1h2v1z"/>`),
		g.Group(children),
	)
}
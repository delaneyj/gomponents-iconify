package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVisa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 15l-1-6l-2.5 6M9 15l1-6M3 9h1v6h.5L7 9m9 .5a.5.5 0 0 0-.5-.5h-.75c-.721 0-1.337.521-1.455 1.233l-.09.534A1.059 1.059 0 0 0 14.25 12a1.059 1.059 0 0 1 1.045 1.233l-.09.534A1.476 1.476 0 0 1 13.75 15H13a.5.5 0 0 1-.5-.5M18 14h2.7"/>`),
		g.Group(children),
	)
}
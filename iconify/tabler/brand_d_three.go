package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4h1.8C8.776 4 12 7.582 12 12s-3.224 8-7.2 8H3m9-16h5.472C19.42 4 21 5.79 21 8s-1.58 4-3.528 4m0 0H15m2.472 0H15.12m2.352 0C19.42 12 21 13.79 21 16s-1.58 4-3.528 4H12"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandBehance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 18V6h4.5a3 3 0 0 1 0 6a3 3 0 0 1 0 6H3m0-6h4.5m6.5 1h7a3.5 3.5 0 0 0-7 0v2a3.5 3.5 0 0 0 6.64 1M16 6h3"/>`),
		g.Group(children),
	)
}
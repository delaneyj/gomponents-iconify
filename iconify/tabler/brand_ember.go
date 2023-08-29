package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandEmber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12.958c8.466 1.647 11.112-1.196 12.17-2.294c2.116-2.196 0-6.589-2.646-5.49C9.88 6.27 6.174 12.86 9.35 17.252C11.466 20.18 15.35 19.43 21 15"/>`),
		g.Group(children),
	)
}
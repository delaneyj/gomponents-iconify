package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandGraphql(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.308 7.265l5.385-3.029m2.615-.001l5.384 3.03M20 9.5v5m-1.307 2.236l-5.385 3.029m-2.616 0l-5.384-3.03M4 14.5v-5m8.772-4.714l6.121 10.202M18.5 16h-13m-.393-1.012l6.122-10.201M10.5 3.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m0 17a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0M2.5 8a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m0 8a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m16 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0m0-8a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0"/>`),
		g.Group(children),
	)
}
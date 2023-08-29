package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonalPrism(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20.792 6.996l-3.775 2.643A2.005 2.005 0 0 1 15.87 10H8.13c-.41 0-.81-.126-1.146-.362L3.21 6.997M8 10v11m8-11v11"/><path d="m3.853 18.274l3.367 2.363A2 2 0 0 0 8.367 21h7.265c.41 0 .811-.126 1.147-.363l3.367-2.363c.536-.375.854-.99.854-1.643V7.369c0-.655-.318-1.268-.853-1.643L16.78 3.363A2 2 0 0 0 15.633 3H8.367c-.41 0-.811.126-1.147.363L3.853 5.726A2.006 2.006 0 0 0 3 7.37v9.261c0 .655.318 1.269.853 1.644z"/></g>`),
		g.Group(children),
	)
}
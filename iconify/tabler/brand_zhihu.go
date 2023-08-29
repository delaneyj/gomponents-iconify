package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandZhihu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6h6v12h-2l-2 2l-1-2h-1zM4 12h6.5m0-6h-5M6 4c-.5 2.5-1.5 3.5-2.5 4.5M8 6v7c0 4.5-2 5.5-4 7m7-2l-3-5"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandZoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.011 9.385v5.128L21 18V6zM3.887 6h10.08C15.435 6 17 7.203 17 8.803v8.196a.991.991 0 0 1-.975 1H5.652c-1.667 0-2.652-1.5-2.652-3l.01-8a.882.882 0 0 1 .208-.71a.841.841 0 0 1 .67-.287z"/>`),
		g.Group(children),
	)
}
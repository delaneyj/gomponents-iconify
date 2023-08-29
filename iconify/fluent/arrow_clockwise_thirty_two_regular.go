package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowClockwiseThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M27 16c0 6.075-4.925 11-11 11S5 22.075 5 16S9.925 5 16 5c2.923 0 5.58 1.14 7.55 3H21a1 1 0 1 0 0 2h5a1 1 0 0 0 1-1V4.5a1 1 0 1 0-2 0v2.12A12.956 12.956 0 0 0 16 3C8.82 3 3 8.82 3 16s5.82 13 13 13s13-5.82 13-13c0-.391-.017-.779-.051-1.161c-.043-.485-.46-.839-.947-.839c-.616 0-1.078.559-1.033 1.173c.02.273.031.549.031.827Z"/>`),
		g.Group(children),
	)
}
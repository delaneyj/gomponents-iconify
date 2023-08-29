package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterclockwiseThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 16c0 6.075 4.925 11 11 11s11-4.925 11-11S22.075 5 16 5c-2.923 0-5.58 1.14-7.55 3H11a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1V4.5a1 1 0 0 1 2 0v2.12A12.956 12.956 0 0 1 16 3c7.18 0 13 5.82 13 13s-5.82 13-13 13S3 23.18 3 16c0-.391.017-.779.051-1.161c.043-.485.46-.839.947-.839c.616 0 1.078.559 1.033 1.173c-.02.273-.031.549-.031.827Z"/>`),
		g.Group(children),
	)
}
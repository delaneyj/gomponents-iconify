package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUpThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 16c0 6.627 5.373 12 12 12s12-5.373 12-12S22.627 4 16 4S4 9.373 4 16Zm12 14C8.268 30 2 23.732 2 16S8.268 2 16 2s14 6.268 14 14s-6.268 14-14 14ZM9.293 18.707a1 1 0 0 1 0-1.414l5.878-5.878c.052-.052.06-.06.11-.102c.235-.196.46-.296.719-.293a.997.997 0 0 1 .717.293c.036.036.072.066.112.101l5.878 5.879a1 1 0 0 1-1.414 1.414L16 13.414l-5.293 5.293a1 1 0 0 1-1.414 0Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleOneThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 16C4 9.373 9.373 4 16 4s12 5.373 12 12s-5.373 12-12 12S4 22.627 4 16ZM16 2C8.268 2 2 8.268 2 16s6.268 14 14 14s14-6.268 14-14S23.732 2 16 2Zm2 7a1 1 0 0 0-1.96-.277l-.001.001l-.006.02a7.117 7.117 0 0 1-.164.455c-.126.314-.324.75-.605 1.228c-.57.967-1.445 2.033-2.706 2.655a1 1 0 1 0 .884 1.794c1.073-.529 1.913-1.273 2.558-2.037V22a1 1 0 1 0 2 0V9Z"/>`),
		g.Group(children),
	)
}
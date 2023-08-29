package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleOneThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 16C2 8.268 8.268 2 16 2s14 6.268 14 14s-6.268 14-14 14S2 23.732 2 16Zm16-7a1 1 0 0 0-1.96-.277l-.001.001l-.006.02a7.117 7.117 0 0 1-.164.455c-.126.314-.324.75-.605 1.228c-.57.967-1.445 2.033-2.706 2.655a1 1 0 1 0 .884 1.794c1.073-.529 1.913-1.273 2.558-2.037V22a1 1 0 1 0 2 0V9Z"/>`),
		g.Group(children),
	)
}
package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Import(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.067 2.25c-5.98 0-11.035 3.91-12.778 9.31H5.5c1.602-3.706 5.27-6.302 9.565-6.31a10.45 10.45 0 0 1 10.437 10.437a10.45 10.45 0 0 1-10.437 10.438c-4.294-.007-7.964-2.605-9.566-6.31H2.29c1.743 5.398 6.798 9.31 12.778 9.31c7.42 0 13.437-6.015 13.438-13.437c-.002-7.423-6.02-13.436-13.44-13.438zm-4.15 17.563l7.15-4.126l-7.15-4.13v2.298H-.056v3.66h10.975v2.298z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRtlAcTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M18 4a1 1 0 0 1 1 1v3h1a1 1 0 1 1 0 2h-1v3a1 1 0 1 1-2 0V5a1 1 0 0 1 1-1zm-8 2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1.25c0 1.105-.216 2.506-.958 3.664C14.261 12.136 12.945 13 11 13a1 1 0 1 1 0-2c1.255 0 1.94-.51 2.358-1.164c.458-.717.642-1.69.642-2.586V7h-3a1 1 0 0 1-1-1zm-5.707 8.293a1 1 0 0 1 1.414 1.414L5.414 16H20a1 1 0 1 1 0 2H5.414l.293.293a1 1 0 1 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414l2-2zm1.414-8a1 1 0 0 0-1.414 0l-2 2a1 1 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414-1.414L5.414 10H8a1 1 0 1 0 0-2H5.414l.293-.293a1 1 0 0 0 0-1.414z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
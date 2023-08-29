package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NegativeDynamics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M17 33.95v8.16M9 40v2.056M25 27v15.071m8-23.11v23.127m8-31.118v31.113M7 33L34 6M7 22v11"/>`),
		g.Group(children),
	)
}
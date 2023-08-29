package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBidirectionalUpDownTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.975 9.689a1 1 0 1 1-1.45-1.378l4.75-5a1 1 0 0 1 1.45 0l4.75 5a1 1 0 1 1-1.45 1.378L13 6.505v10.99l3.025-3.184a1 1 0 1 1 1.45 1.378l-4.75 5a1 1 0 0 1-1.45 0l-4.75-5a1 1 0 1 1 1.45-1.378L11 17.495V6.505L7.975 9.689Z"/>`),
		g.Group(children),
	)
}
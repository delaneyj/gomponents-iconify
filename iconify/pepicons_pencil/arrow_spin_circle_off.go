package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSpinCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M16.374 8.038a.5.5 0 0 1-.563.827A5 5 0 1 0 18 13a.5.5 0 0 1 1 0a6 6 0 1 1-2.626-4.962Z"/><path d="M15.769 14.585a.5.5 0 1 1-.539-.842l3.482-2.227a.5.5 0 1 1 .539.842l-3.482 2.227Z"/><path d="M20.947 15.114a.5.5 0 0 1-.913.407l-1.509-3.38a.5.5 0 1 1 .914-.408l1.508 3.381ZM4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
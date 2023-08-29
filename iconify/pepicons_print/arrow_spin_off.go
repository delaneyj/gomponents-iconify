package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSpinOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="11.5" r="6.5" opacity=".2"/><path d="M13.374 5.038a.5.5 0 0 1-.563.827A5 5 0 1 0 15 10a.5.5 0 0 1 1 0a6 6 0 1 1-2.626-4.962Z"/><path d="M12.769 11.585a.5.5 0 1 1-.539-.842l3.482-2.227a.5.5 0 1 1 .539.842l-3.482 2.227Z"/><path d="M17.947 12.114a.5.5 0 0 1-.913.407l-1.509-3.38a.5.5 0 1 1 .914-.408l1.508 3.381ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}
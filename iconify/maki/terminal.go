package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.3 2.1a.5.5 0 0 1 .409-.017l5.926 2.372l3.514-1.457a1.35 1.35 0 1 1 .993 2.513L4.709 9.058a2 2 0 0 1-2.12-.46L.195 6.195a.2.2 0 0 1 .052-.32l.507-.253a.5.5 0 0 1 .48.018L3.5 7l3.048-1.264l-3.804-3.04a.2.2 0 0 1 .036-.336l.52-.26ZM.5 12a.5.5 0 0 0 0 1h14a.5.5 0 0 0 0-1H.5Z"/>`),
		g.Group(children),
	)
}
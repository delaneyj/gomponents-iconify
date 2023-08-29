package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Summit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.394 14.32L24 18.123l3.606-3.803h-7.212ZM24 38.052L4.5 33.367L24 9.948v28.103l19.5-4.685L24 9.949"/>`),
		g.Group(children),
	)
}
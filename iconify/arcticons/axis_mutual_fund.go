package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AxisMutualFund(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.699 16.841L24 26.704l-7.113 12.319H5.5L24 6.977l5.699 9.864z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 39.023H31.113L24 26.704h11.387L42.5 39.023z"/>`),
		g.Group(children),
	)
}
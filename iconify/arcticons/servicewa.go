package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Servicewa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.041 27.504h-3.653m-.911 2.796l2.741-8.408l2.74 8.433m-7.484-8.433l-2.108 8.433l-2.108-8.433l-2.108 8.433l-2.108-8.433"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.913 35.453v-26L31.992 4.5l-8.608 11.143l-12.298 4.952l1.23 8.667l3.074 8.048l-1.844 4.333l7.378 1.857l.527-1.238l15.46-6.809Z"/>`),
		g.Group(children),
	)
}
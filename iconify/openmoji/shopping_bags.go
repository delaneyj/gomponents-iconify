package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#EA5A47" d="M25.747 21.002h10.254"/><path fill="#FCEA2B" d="M36.002 60.96h25l-5-40.867h-40l-5 40.867h25"/><path fill="#EA5A47" d="M46.256 21.002H36.002"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M36.002 60.96h-25l5-39.958h2.569m7.176 0h10.254m.001 39.958h25l-5-39.958h-2.57m-7.176 0H36.002m-13.953 6.66v-6.62c0-7.754 6.286-14.04 14.04-14.04s14.041 6.286 14.041 14.04v6.62"/>`),
		g.Group(children),
	)
}
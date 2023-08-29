package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketballOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBasketballOne0"><g fill="none"><path fill="#555" d="M18 24c0-6.624 5.376-12 12-12s12 5.376 12 12"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M18 24c0-6.624 5.376-12 12-12s12 5.376 12 12M12 4H4v32h8V4Zm0 20h32m-26 0l2 4.5c2 5 1.9 10.4 0 15.5m21-20l-2 4.5c-2 5-1.91 10.4 0 15.5M21.22 32h16.56m-16.09 8h15.63m-7.82 0V24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBasketballOne0)"/>`),
		g.Group(children),
	)
}
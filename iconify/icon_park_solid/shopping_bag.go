package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShoppingBag0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 17h38l-4.2 26H9.2L5 17Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 17c0-6.627-4.925-12-11-12s-11 5.373-11 12"/><circle cx="17" cy="26" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 33s2 3 6 3s6-3 6-3"/><circle cx="31" cy="26" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShoppingBag0)"/>`),
		g.Group(children),
	)
}
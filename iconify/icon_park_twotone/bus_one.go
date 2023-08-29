package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBusOne0"><g fill="none"><rect width="32" height="34" x="8" y="5" fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="3"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 39v4m20-4v4"/><circle cx="34" cy="33" r="2" fill="#fff"/><circle cx="14" cy="33" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 23h32M8 21v4m32-4v4M18 13h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBusOne0)"/>`),
		g.Group(children),
	)
}
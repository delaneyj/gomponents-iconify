package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeLightroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAdobeLightroom0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 15v18h8m6 0V21m0 5c2.25-4 4.629-4 6-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAdobeLightroom0)"/>`),
		g.Group(children),
	)
}
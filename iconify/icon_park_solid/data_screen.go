package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDataScreen0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="30" x="4" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#fff" d="M24 36v7"/><path stroke="#000" d="M32 14L16 28"/><path stroke="#fff" d="M10 43h28"/><circle cx="15" cy="17" r="3" fill="#000" stroke="#000"/><circle cx="33" cy="25" r="3" fill="#000" stroke="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDataScreen0)"/>`),
		g.Group(children),
	)
}
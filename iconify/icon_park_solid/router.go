package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Router(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRouter0"><g fill="none"><rect width="40" height="14" x="4" y="28" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 35h8"/><rect width="4" height="4" x="30" y="33" fill="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 28V8m24 20V8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRouter0)"/>`),
		g.Group(children),
	)
}
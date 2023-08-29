package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBox0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="30" x="6" y="12" fill="#555" rx="2"/><path stroke-linecap="round" d="M17.95 24.008h12M6 13l7-8h22l7 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBox0)"/>`),
		g.Group(children),
	)
}
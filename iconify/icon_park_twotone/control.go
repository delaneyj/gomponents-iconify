package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Control(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTControl0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="30" height="40" x="9" y="4" rx="2"/><circle cx="24" cy="31" r="6" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 12h3m-3 6h3m6 0h3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTControl0)"/>`),
		g.Group(children),
	)
}
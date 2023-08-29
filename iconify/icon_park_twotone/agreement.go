package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Agreement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAgreement0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><rect width="32" height="40" x="8" y="4" stroke-linejoin="round" rx="2"/><path fill="#555" stroke-linejoin="round" d="M16 4h9v16l-4.5-4l-4.5 4V4Z"/><path d="M16 28h10m-10 6h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAgreement0)"/>`),
		g.Group(children),
	)
}
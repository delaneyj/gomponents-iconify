package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarginOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMarginOne0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M6 34h36M6 14h36"/><path stroke="#fff" stroke-linecap="round" d="M6 11v6m0 14v6m36-6v6m0-26v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMarginOne0)"/>`),
		g.Group(children),
	)
}
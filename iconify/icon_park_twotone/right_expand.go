package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRightExpand0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="28" height="36" x="14" y="6" fill="#555" rx="2"/><path d="M6 6v36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRightExpand0)"/>`),
		g.Group(children),
	)
}
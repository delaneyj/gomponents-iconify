package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRightExpand0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="28" height="36" x="14" y="6" fill="#fff" rx="2"/><path d="M6 6v36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRightExpand0)"/>`),
		g.Group(children),
	)
}
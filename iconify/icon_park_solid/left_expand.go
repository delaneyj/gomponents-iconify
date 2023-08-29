package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLeftExpand0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="28" height="36" x="6" y="6" fill="#fff" rx="2"/><path d="M42 6v36"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLeftExpand0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSExpandUp0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M6 16h36M20 32l4-4l4 4"/><path stroke="#fff" stroke-linecap="round" d="M6 10v12m36-12v12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSExpandUp0)"/>`),
		g.Group(children),
	)
}
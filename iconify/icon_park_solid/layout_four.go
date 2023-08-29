package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLayoutFour0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M6 28h18m0-8h18"/><path stroke="#fff" stroke-linecap="round" d="M6 25v6m36-14v6"/><path stroke="#000" stroke-linecap="round" d="M24 42V6"/><path stroke="#fff" stroke-linecap="round" d="M21 6h6m-6 36h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLayoutFour0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLayoutThree0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path stroke-linecap="round" d="M6 22h36m-13 0V6m-3 0h6M6 19v6m36-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLayoutThree0)"/>`),
		g.Group(children),
	)
}
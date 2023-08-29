package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTimer0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="28" r="16" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M28 4h-8m4 0v8m11 4l3-3M24 28v-6m0 6h-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTimer0)"/>`),
		g.Group(children),
	)
}
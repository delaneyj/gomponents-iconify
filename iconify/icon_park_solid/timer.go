package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTimer0"><g fill="none" stroke-width="4"><circle cx="24" cy="28" r="16" fill="#fff" stroke="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M28 4h-8m4 0v8m11 4l3-3"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 28v-6m0 6h-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTimer0)"/>`),
		g.Group(children),
	)
}
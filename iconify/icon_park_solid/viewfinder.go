package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viewfinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSViewfinder0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M16 6H8a2 2 0 0 0-2 2v8m10 26H8a2 2 0 0 1-2-2v-8m26 10h8a2 2 0 0 0 2-2v-8M32 6h8a2 2 0 0 1 2 2v8"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M13 33V18h6l2-3h6l2 3h6v15H13Z"/><path fill="#000" stroke="#000" stroke-miterlimit="10" d="M24 28a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSViewfinder0)"/>`),
		g.Group(children),
	)
}
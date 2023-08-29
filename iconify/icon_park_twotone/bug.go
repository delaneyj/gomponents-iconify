package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBug0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 42c12 0 14-10.468 14-14V14H10v14c0 3.45 2 14 14 14Z"/><path stroke-linecap="round" d="m4 8l6 6m34-6l-6 6M4 27h6m34 0h-6M7 44l6-6m28 6l-6-6m-11 4V14m-9.08 25.04C17.002 40.784 19.924 42 24 42c4.111 0 7.049-1.229 9.134-2.986"/><path fill="#555" d="M32 12.333C32 7.731 28.418 4 24 4s-8 3.731-8 8.333V14h16v-1.667Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBug0)"/>`),
		g.Group(children),
	)
}
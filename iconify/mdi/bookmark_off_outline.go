package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.28 4L2 5.27l3 3V21l7-3l4.78 2.05L18.73 22L20 20.72L3.28 4M7 18v-7.73l6 5.98l-1-.43L7 18M7 5.16L5.5 3.67C5.88 3.26 6.41 3 7 3h10a2 2 0 0 1 2 2v12.16l-2-2V5H7v.16Z"/>`),
		g.Group(children),
	)
}
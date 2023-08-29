package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Build(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M22.5 6.5c0-2.31-.439-3-3-3h-16c-2.56 0-3 .69-3 3v29c0 2.5.62 3 3 3h10v-2l9-8v-22zm-17 2h12v24h-12v-24zm7 28h-3v-2h3v2zm13.561-8.391L18.05 35.4c-1.04 1.108-2.7 2.01-1.16 3.898c1.28 1.57 2.969.12 3.93-.898l7.99-7.41l-2.749-2.881zm12.048-3.048l-2.27 1.959l-1.91-1.97l.871-2.15l-2.14-2.682s-4.16-3.31-9.36-.05c2.341-.109 4.58 1.181 6.2 4.021l-2.189 2.121l2.75 2.629l1.799-1.619l2.201 1.93l-1.761 1.68l2.14 2.16l5.761-5.58l-2.092-2.449z"/>`),
		g.Group(children),
	)
}
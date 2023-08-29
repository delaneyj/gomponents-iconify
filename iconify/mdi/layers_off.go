package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.27 1L2 2.27L6.22 6.5L3 9l1.63 1.27L12 16l2.1-1.63l1.43 1.43L12 18.54l-7.37-5.73L3 14.07l9 7l4.95-3.85L20.73 21L22 19.73L3.27 1m16.09 9.27L21 9l-9-7l-2.91 2.27l7.87 7.88l2.4-1.88m.45 4.73l1.19-.93l-1.43-1.43l-1.19.92L19.81 15Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.809 1.394L20.486 19.07l-1.414 1.414l-3.683-3.681L12.001 21L.691 6.997a18 18 0 0 1 2.95-1.942L1.395 2.808L2.81 1.394Zm.771 5.998L12 17.817l1.967-2.436l-8.835-8.836c-.532.253-1.05.536-1.552.847ZM12.001 3c4.284 0 8.22 1.497 11.31 3.996l-5.407 6.693l-1.422-1.422l3.939-4.876A15.921 15.921 0 0 0 12 5.001c-.873 0-1.735.07-2.58.206L7.727 3.51c1.37-.334 2.802-.51 4.275-.51Z"/>`),
		g.Group(children),
	)
}
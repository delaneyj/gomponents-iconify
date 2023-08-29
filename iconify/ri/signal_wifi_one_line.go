package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiOneLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 3c4.284 0 8.22 1.497 11.31 3.996L12 21L.691 6.997A17.925 17.925 0 0 1 12 3Zm0 12a6.01 6.01 0 0 0-2 .34l2 2.477l2-2.477a6.066 6.066 0 0 0-2-.34Zm0-10c-3.028 0-5.923.842-8.42 2.392l5.108 6.324A7.971 7.971 0 0 1 12.001 13c1.181 0 2.303.256 3.312.716L20.42 7.39A15.921 15.921 0 0 0 12 5Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiErrorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 3c4.284 0 8.22 1.497 11.31 3.996L22.499 8h-4.498v5.571L12 21L.691 6.997A17.925 17.925 0 0 1 12 3Zm10 16v2h-2v-2h2Zm0-9v7h-2v-7h2Z"/>`),
		g.Group(children),
	)
}
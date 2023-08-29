package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 3c4.284 0 8.22 1.497 11.31 3.996L12 21L.691 6.997A17.925 17.925 0 0 1 12 3Z"/>`),
		g.Group(children),
	)
}
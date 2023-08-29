package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 3c4.284 0 8.22 1.497 11.31 3.996L12 21L.691 6.997A17.925 17.925 0 0 1 12 3Zm0 7c-1.898 0-3.683.48-5.241 1.327l5.24 6.49l5.242-6.49A10.95 10.95 0 0 0 12.001 10Zm0-5c-3.028 0-5.923.842-8.42 2.392l1.904 2.357A12.94 12.94 0 0 1 12 8c2.375 0 4.6.637 6.516 1.749l1.904-2.358A15.921 15.921 0 0 0 12 5Z"/>`),
		g.Group(children),
	)
}
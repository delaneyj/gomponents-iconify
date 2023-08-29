package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellZLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M150 144a6 6 0 0 1-6 6h-32a6 6 0 0 1-5-9.33L132.79 102H112a6 6 0 0 1 0-12h32a6 6 0 0 1 5 9.33L123.21 138H144a6 6 0 0 1 6 6Zm70.11 47a13.83 13.83 0 0 1-12.1 7h-42.49a38 38 0 0 1-75 0H48a14 14 0 0 1-12.06-21.06C41.59 167.2 50 139.74 50 104a78 78 0 1 1 156 0c0 35.73 8.42 63.2 14.08 72.94a13.9 13.9 0 0 1 .03 14.06Zm-66.82 7h-50.58a26 26 0 0 0 50.58 0Zm56.41-15c-7.57-13-15.7-43.32-15.7-79a66 66 0 1 0-132 0c0 35.69-8.14 66-15.71 79a2 2 0 0 0 0 2a1.9 1.9 0 0 0 1.7 1H208a1.9 1.9 0 0 0 1.7-1a2 2 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 8h-2.086l1.293-1.293a.999.999 0 1 0-1.414-1.414L10.586 9l3.707 3.707a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L14.414 10H16.5c1.379 0 2.5 1.346 2.5 3s-1.346 3-3 3H8c-1.654 0-3-1.346-3-3s1.346-3 3-3a1 1 0 1 0 0-2c-2.757 0-5 2.243-5 5s2.243 5 5 5h8c2.757 0 5-2.243 5-5s-2.019-5-4.5-5z"/>`),
		g.Group(children),
	)
}
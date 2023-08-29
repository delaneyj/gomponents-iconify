package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpBookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 18l2 1V1H7v2h12v15zM17 5H3v18l7-3l7 3V5z"/>`),
		g.Group(children),
	)
}
package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9.414V15a1 1 0 1 1-2 0V7a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H9.414l8.293 8.293a1 1 0 0 1-1.414 1.414L8 9.414z"/>`),
		g.Group(children),
	)
}
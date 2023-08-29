package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleRightF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.828 10l-3.535 3.536a1 1 0 0 0 1.414 1.414l4.243-4.243a1 1 0 0 0 0-1.414L9.707 5.05a1 1 0 0 0-1.414 1.414L11.828 10zM10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10z"/>`),
		g.Group(children),
	)
}
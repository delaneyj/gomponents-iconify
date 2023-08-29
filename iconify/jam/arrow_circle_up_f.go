package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.707 5.293a.997.997 0 0 0-1.414 0L5.05 9.536a1 1 0 0 0 1.414 1.414L9 8.414V14a1 1 0 0 0 2 0V8.414l2.536 2.536a1 1 0 0 0 1.414-1.414l-4.243-4.243zM10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10z"/>`),
		g.Group(children),
	)
}
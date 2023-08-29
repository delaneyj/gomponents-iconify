package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareUpF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.707 5.293a.997.997 0 0 0-1.414 0L5.05 9.536a1 1 0 0 0 1.414 1.414L9 8.414V14a1 1 0 0 0 2 0V8.414l2.536 2.536a1 1 0 0 0 1.414-1.414l-4.243-4.243zM4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/>`),
		g.Group(children),
	)
}
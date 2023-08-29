package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsSquareRightF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4zm9.212 10.021l-2.121 2.122a1 1 0 0 0 1.414 1.414l2.828-2.829a1 1 0 0 0 0-1.414l-2.828-2.828A1 1 0 1 0 11.09 7.9l2.12 2.121zm-4.99 0l-2.12 2.122a1 1 0 0 0 1.413 1.414l2.829-2.829a1 1 0 0 0 0-1.414L7.515 6.486A1 1 0 0 0 6.101 7.9l2.121 2.121z"/>`),
		g.Group(children),
	)
}
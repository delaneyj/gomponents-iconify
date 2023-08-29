package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsSquareLeftF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.818 10l2.121-2.121a1 1 0 1 0-1.414-1.415L9.697 9.293a1 1 0 0 0 0 1.414l2.828 2.829a1 1 0 0 0 1.414-1.415L11.818 10zm-4.99 0L8.95 7.879a1 1 0 0 0-1.414-1.415L4.707 9.293a1 1 0 0 0 0 1.414l2.829 2.829A1 1 0 1 0 8.95 12.12L6.828 10zM4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/>`),
		g.Group(children),
	)
}
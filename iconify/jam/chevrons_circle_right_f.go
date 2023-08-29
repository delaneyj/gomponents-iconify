package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleRightF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm3.175-9.976l-2.122 2.121a1 1 0 0 0 1.414 1.414l2.829-2.828a1 1 0 0 0 0-1.414l-2.829-2.829a1 1 0 1 0-1.414 1.415l2.122 2.12zm-4.99 0l-2.121 2.121a1 1 0 0 0 1.414 1.414l2.828-2.828a1 1 0 0 0 0-1.414L7.478 6.488a1 1 0 0 0-1.414 1.415l2.121 2.12z"/>`),
		g.Group(children),
	)
}
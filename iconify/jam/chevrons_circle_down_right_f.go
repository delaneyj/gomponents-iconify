package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleDownRightF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9H6a1 1 0 1 0 0 2h4a1 1 0 0 0 1-1V6a1 1 0 0 0-2 0v3zm3 3H9a1 1 0 0 0 0 2h4a1 1 0 0 0 1-1V9a1 1 0 0 0-2 0v3zm-2 8C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10z"/>`),
		g.Group(children),
	)
}
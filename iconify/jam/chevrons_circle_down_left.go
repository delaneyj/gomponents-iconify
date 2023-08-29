package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm-2-6h3a1 1 0 0 1 0 2H7a1 1 0 0 1-1-1V9a1 1 0 1 1 2 0v3zm3-3h3a1 1 0 0 1 0 2h-4a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v3z"/>`),
		g.Group(children),
	)
}
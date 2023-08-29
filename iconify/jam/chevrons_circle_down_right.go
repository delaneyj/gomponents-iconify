package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm2-6V9a1 1 0 0 1 2 0v4a1 1 0 0 1-1 1H9a1 1 0 0 1 0-2h3zM9 9V6a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1H6a1 1 0 0 1 0-2h3z"/>`),
		g.Group(children),
	)
}
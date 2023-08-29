package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm2.818-8l-2.121-2.121a1 1 0 0 1 1.414-1.415l2.828 2.829a1 1 0 0 1 0 1.414l-2.828 2.829a1 1 0 1 1-1.414-1.415L12.818 10zm-4.99 0l-2.12-2.121A1 1 0 0 1 7.12 6.464l2.83 2.829a1 1 0 0 1 0 1.414l-2.83 2.829a1 1 0 1 1-1.414-1.415L7.828 10z"/>`),
		g.Group(children),
	)
}
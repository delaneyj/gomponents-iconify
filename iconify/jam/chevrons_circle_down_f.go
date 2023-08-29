package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleDownF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm.009-7.377l-2.122-2.12a1 1 0 0 0-1.414 1.413l2.828 2.829a1 1 0 0 0 1.415 0l2.828-2.829a1 1 0 1 0-1.414-1.414l-2.121 2.121zm0-5l-2.122-2.12a1 1 0 0 0-1.414 1.413l2.828 2.829a1 1 0 0 0 1.415 0l2.828-2.829a1 1 0 1 0-1.414-1.414l-2.121 2.121z"/>`),
		g.Group(children),
	)
}
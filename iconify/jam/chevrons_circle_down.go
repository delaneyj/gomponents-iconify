package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm0-5.29l2.121-2.12a1 1 0 0 1 1.415 1.414l-2.829 2.828a1 1 0 0 1-1.414 0l-2.829-2.828a1 1 0 1 1 1.415-1.415L10 12.711zm0-5l2.121-2.12a1 1 0 0 1 1.415 1.414l-2.829 2.828a1 1 0 0 1-1.414 0L6.464 7.004A1 1 0 1 1 7.88 5.589L10 7.711z"/>`),
		g.Group(children),
	)
}
package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsCircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm0 2C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10z"/><path d="M10 7.382L7.879 9.503A1 1 0 1 1 6.464 8.09l2.829-2.828a1 1 0 0 1 1.414 0l2.829 2.828a1 1 0 0 1-1.415 1.414L10 7.383z"/><path d="m10 12.382l-2.121 2.121a1 1 0 1 1-1.415-1.414l2.829-2.828a1 1 0 0 1 1.414 0l2.829 2.828a1 1 0 0 1-1.415 1.414L10 12.383z"/></g>`),
		g.Group(children),
	)
}
package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm2.529 13.674H6.947a.698.698 0 0 0 0 1.396h5.582a.698.698 0 1 0 0-1.396ZM5.814 6.28a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79Zm8.372 0a1.395 1.395 0 1 0 0 2.79a1.395 1.395 0 0 0 0-2.79Z"/>`),
		g.Group(children),
	)
}
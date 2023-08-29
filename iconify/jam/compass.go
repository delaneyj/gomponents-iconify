package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.932 9.482a2.003 2.003 0 0 1-2.45 2.45L6.464 14.95a1 1 0 1 1-1.414-1.414l3.018-3.018a2.003 2.003 0 0 1 2.45-2.45l3.018-3.018a1 1 0 0 1 1.414 1.414l-3.018 3.018zM10 20C4.477 20 0 15.523 0 10S4.477 0 10 0s10 4.477 10 10s-4.477 10-10 10zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16z"/>`),
		g.Group(children),
	)
}
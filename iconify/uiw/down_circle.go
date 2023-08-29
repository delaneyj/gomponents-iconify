package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0ZM6.586 7.636a.682.682 0 1 0-.962.967l3.895 3.87c.265.263.69.265.957.005l3.971-3.87a.682.682 0 1 0-.951-.977l-3.491 3.402Z"/>`),
		g.Group(children),
	)
}
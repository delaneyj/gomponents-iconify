package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 6.5V8H7a5 5 0 0 0-5 5v7.5A3.5 3.5 0 0 0 5.5 24H7v1.5a3.5 3.5 0 0 0 3.5 3.5h11a3.5 3.5 0 0 0 3.5-3.5V24h1.5a3.5 3.5 0 0 0 3.5-3.5V13a5 5 0 0 0-5-5h-1V6.5A3.5 3.5 0 0 0 20.5 3h-9A3.5 3.5 0 0 0 8 6.5ZM11.5 5h9A1.5 1.5 0 0 1 22 6.5V8H10V6.5A1.5 1.5 0 0 1 11.5 5ZM9 19.5a1.5 1.5 0 0 1 1.5-1.5h11a1.5 1.5 0 0 1 1.5 1.5v6a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 9 25.5v-6Z"/>`),
		g.Group(children),
	)
}
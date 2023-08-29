package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorHoverFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.75 9.25a3.5 3.5 0 0 0-3.5 3.5v18.5a3.5 3.5 0 0 0 3.5 3.5H18v-13.5a3.25 3.25 0 0 1 5.548-2.298l15.5 15.5c.095.095.183.194.263.297a3.5 3.5 0 0 0 3.439-3.499v-18.5a3.5 3.5 0 0 0-3.5-3.5H8.75Zm13.384 11.116A1.25 1.25 0 0 0 20 21.25v21.5a1.25 1.25 0 0 0 2.285.7l4.961-7.333a.75.75 0 0 1 .808-.306l8.386 2.15a1.25 1.25 0 0 0 1.194-2.095l-15.5-15.5Z"/>`),
		g.Group(children),
	)
}
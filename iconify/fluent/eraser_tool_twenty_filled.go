package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserToolTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2.5a.5.5 0 0 0-1 0v12A3.5 3.5 0 0 0 6.5 18h7a3.5 3.5 0 0 0 3.5-3.5v-12a.5.5 0 0 0-1 0V6H4V2.5ZM4 10V7h12v3H4Z"/>`),
		g.Group(children),
	)
}
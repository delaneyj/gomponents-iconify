package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDensitySixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 1a.5.5 0 0 1 .5.5v13a.5.5 0 0 1-1 0v-13a.5.5 0 0 1 .5-.5ZM7 3H2.5a.5.5 0 0 0 0 1H7V3Zm0 3H2.5a.5.5 0 0 0 0 1H7V6Zm0 3H2.5a.5.5 0 0 0 0 1H7V9Zm0 3H2.5a.5.5 0 0 0 0 1H7v-1Zm5.5 0H10V9h2.5a1.5 1.5 0 0 1 0 3Zm0-5H10V4h2.5a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}
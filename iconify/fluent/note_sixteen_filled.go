package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5A2.5 2.5 0 0 1 4.5 2h7A2.5 2.5 0 0 1 14 4.5V8h-3.5A2.5 2.5 0 0 0 8 10.5V14H4.5A2.5 2.5 0 0 1 2 11.5v-7Zm7 9.359a2.5 2.5 0 0 0 .94-.591l3.328-3.329a2.5 2.5 0 0 0 .59-.939H10.5A1.5 1.5 0 0 0 9 10.5v3.359Z"/>`),
		g.Group(children),
	)
}
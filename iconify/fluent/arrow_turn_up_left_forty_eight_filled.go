package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnUpLeftFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M35 40.5a1.5 1.5 0 0 0 3 0v-18a7.5 7.5 0 0 0-7.5-7.5H13.121l6.44-6.44a1.5 1.5 0 0 0-2.122-2.12l-9 9a1.5 1.5 0 0 0 0 2.12l9 9a1.5 1.5 0 0 0 2.122-2.12L13.12 18H30.5a4.5 4.5 0 0 1 4.5 4.5v18Z"/>`),
		g.Group(children),
	)
}
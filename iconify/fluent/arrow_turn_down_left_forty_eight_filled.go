package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnDownLeftFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M35 7.5a1.5 1.5 0 0 1 3 0v18a7.5 7.5 0 0 1-7.5 7.5H13.121l6.44 6.44a1.5 1.5 0 0 1-2.122 2.12l-9-9a1.5 1.5 0 0 1 0-2.12l9-9a1.5 1.5 0 0 1 2.122 2.12L13.12 30H30.5a4.5 4.5 0 0 0 4.5-4.5v-18Z"/>`),
		g.Group(children),
	)
}
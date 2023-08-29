package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnDownRightFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 7.5a1.5 1.5 0 0 0-3 0v18a7.5 7.5 0 0 0 7.5 7.5h17.379l-6.44 6.44a1.5 1.5 0 0 0 2.122 2.12l9-9a1.5 1.5 0 0 0 0-2.12l-9-9a1.5 1.5 0 0 0-2.122 2.12L34.88 30H17.5a4.5 4.5 0 0 1-4.5-4.5v-18Z"/>`),
		g.Group(children),
	)
}
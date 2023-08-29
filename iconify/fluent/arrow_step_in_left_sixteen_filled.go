package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowStepInLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.707 8.5H14.5a.5.5 0 0 0 0-1H7.707l2.147-2.146a.5.5 0 0 0-.708-.708l-3 3a.5.5 0 0 0 0 .708l3 3a.5.5 0 0 0 .708-.708L7.707 8.5ZM1 8a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/>`),
		g.Group(children),
	)
}
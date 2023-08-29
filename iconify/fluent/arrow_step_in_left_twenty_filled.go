package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowStepInLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm14-2a.5.5 0 0 1-.5.5H9.707l3.147 3.146a.5.5 0 0 1-.708.708l-4-4a.5.5 0 0 1 0-.708l4-4a.5.5 0 0 1 .708.708L9.707 9.5H17.5a.5.5 0 0 1 .5.5Z"/>`),
		g.Group(children),
	)
}
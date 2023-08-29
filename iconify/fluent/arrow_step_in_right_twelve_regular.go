package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowStepInRightTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 6a.5.5 0 0 1 .5-.5h3.793L4.146 4.354a.5.5 0 1 1 .708-.708l2 2a.5.5 0 0 1 0 .708l-2 2a.5.5 0 1 1-.708-.708L5.293 6.5H1.5A.5.5 0 0 1 1 6Zm8.5-1.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm0 1a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}
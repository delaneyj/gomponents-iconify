package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowStepInLeftTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 6a.5.5 0 0 0-.5-.5H6.707l1.147-1.146a.5.5 0 1 0-.708-.708l-2 2a.5.5 0 0 0 0 .708l2 2a.5.5 0 1 0 .708-.708L6.707 6.5H10.5A.5.5 0 0 0 11 6ZM2.5 4.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm0 1a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}
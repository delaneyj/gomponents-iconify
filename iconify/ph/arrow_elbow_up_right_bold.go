package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowUpRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m216.49 88.49l-48 48a12 12 0 0 1-17-17L179 92H76v132a12 12 0 0 1-24 0V80a12 12 0 0 1 12-12h115l-27.49-27.51a12 12 0 1 1 17-17l48 48a12 12 0 0 1-.02 17Z"/>`),
		g.Group(children),
	)
}
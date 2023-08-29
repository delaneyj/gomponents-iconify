package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowDownRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m216.49 184.49l-48 48a12 12 0 0 1-17-17L179 188H64a12 12 0 0 1-12-12V32a12 12 0 0 1 24 0v132h103l-27.52-27.51a12 12 0 1 1 17-17l48 48a12 12 0 0 1 .01 17Z"/>`),
		g.Group(children),
	)
}
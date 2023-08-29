package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184.49 119.51a12 12 0 0 1 0 17l-72 72a12 12 0 0 1-17-17L147 140H32a12 12 0 0 1 0-24h115L95.51 64.49a12 12 0 0 1 17-17ZM216 28a12 12 0 0 0-12 12v176a12 12 0 0 0 24 0V40a12 12 0 0 0-12-12Z"/>`),
		g.Group(children),
	)
}
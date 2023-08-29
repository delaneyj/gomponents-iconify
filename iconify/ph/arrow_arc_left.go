package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowArcLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 184a8 8 0 0 1-16 0a88 88 0 0 0-150.22-62.22L43.4 144H88a8 8 0 0 1 0 16H24a8 8 0 0 1-8-8V88a8 8 0 0 1 16 0v44.77l22.48-22.33A104 104 0 0 1 232 184Z"/>`),
		g.Group(children),
	)
}
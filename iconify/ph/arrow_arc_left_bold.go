package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowArcLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 184a12 12 0 0 1-24 0a84 84 0 0 0-143.4-59.4L53.11 140H88a12 12 0 0 1 0 24H24a12 12 0 0 1-12-12V88a12 12 0 0 1 24 0v35.16l15.66-15.55A108 108 0 0 1 236 184Z"/>`),
		g.Group(children),
	)
}
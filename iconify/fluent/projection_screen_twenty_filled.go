package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectionScreenTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3a1 1 0 0 0 0 2v6a3 3 0 0 0 3 3h3.5v2h-2a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-2v-2H14a3 3 0 0 0 3-3V5a1 1 0 1 0 0-2H3Z"/>`),
		g.Group(children),
	)
}
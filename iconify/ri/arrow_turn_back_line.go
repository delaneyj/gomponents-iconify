package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnBackLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 18.172l-2.535-2.536l-1.414 1.414L18 22l4.95-4.95l-1.415-1.414L19 18.172V11a8 8 0 1 0-16 0v9h2v-9a6 6 0 0 1 12 0v7.172Z"/>`),
		g.Group(children),
	)
}
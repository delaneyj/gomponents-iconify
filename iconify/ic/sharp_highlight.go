package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpHighlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 14l3 3v5h6v-5l3-3V9H6v5zm5-12h2v3h-2V2zM3.5 5.88l1.41-1.41l2.12 2.12L5.62 8L3.5 5.88zm13.46.71l2.12-2.12l1.41 1.41L18.38 8l-1.42-1.41z"/>`),
		g.Group(children),
	)
}
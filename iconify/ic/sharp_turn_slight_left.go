package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpTurnSlightLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.66 6V4H6v5.66h2V7.41l5 5V20h2v-8.41L9.41 6z"/>`),
		g.Group(children),
	)
}
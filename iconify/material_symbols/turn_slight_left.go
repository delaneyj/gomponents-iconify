package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurnSlightLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20v-7.6l-5-5v2.25H6V4h5.65v2H9.4l5.025 5.025q.275.275.425.638t.15.762V20h-2Z"/>`),
		g.Group(children),
	)
}
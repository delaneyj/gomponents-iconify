package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.4 20L13 18.6l2.6-2.6l-2.6-2.6l1.4-1.4l2.6 2.6l2.6-2.6l1.4 1.4l-2.6 2.6l2.6 2.6l-1.4 1.4l-2.6-2.6l-2.6 2.6Zm1.975-9l-3.55-3.55l1.4-1.4l2.125 2.125l4.25-4.25L22 5.35L16.375 11ZM2 17v-2h9v2H2Zm0-8V7h9v2H2Z"/>`),
		g.Group(children),
	)
}
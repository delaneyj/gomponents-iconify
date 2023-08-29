package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneScreenshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2H7m0 3h10v16H7V4m2 2v4h1.5V7.5H13V6H9m4.5 8v2.5H11V18h4v-4h-1.5Z"/>`),
		g.Group(children),
	)
}
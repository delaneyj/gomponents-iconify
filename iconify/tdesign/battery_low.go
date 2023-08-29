package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 5h21v14H0V5Zm2 2v10h17V7H2Zm4 2v6H4V9h2Zm18 0v6h-2V9h2Z"/>`),
		g.Group(children),
	)
}
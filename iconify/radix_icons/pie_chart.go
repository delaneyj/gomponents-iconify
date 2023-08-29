package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.85 7.5a5.65 5.65 0 1 1 11.3 0a5.65 5.65 0 0 1-11.3 0ZM7.5.85a6.65 6.65 0 1 0 0 13.3a6.65 6.65 0 0 0 0-13.3ZM7 8V3.128A4.4 4.4 0 0 1 11.872 8H7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
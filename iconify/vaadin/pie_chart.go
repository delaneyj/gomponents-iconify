package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 1c3.2.2 5.7 2.8 6 6H9V1zm-.5-1H8v8h8v-.5C16 3.4 12.6 0 8.5 0z"/><path fill="currentColor" d="M7 9V1c-3.9.3-7 3.5-7 7.5C0 12.6 3.4 16 7.5 16c4 0 7.2-3.1 7.5-7H7z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckSquareO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 6.2V14H2V2h10.5l1-1H1v14h14V5.2z"/><path fill="currentColor" d="M7.9 10.9L3.7 6.7l1.5-1.4l2.7 2.8l6.7-6.7L16 2.8z"/>`),
		g.Group(children),
	)
}
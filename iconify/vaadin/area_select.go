package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m7.9 7.9l2.1 7.5l1.7-2.6l3.2 3.2l1.1-1.1l-3.3-3.2l2.7-1.6z"/><path fill="currentColor" d="M8 12H1V3h12v5.4l1 .2V2H0v11h8.2z"/>`),
		g.Group(children),
	)
}
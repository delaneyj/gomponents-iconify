package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8zm9 1.6l1.8-1.8l1.4 1.4L8 13.4L3.8 9.2l1.4-1.4L7 9.6V3h2v6.6z"/>`),
		g.Group(children),
	)
}
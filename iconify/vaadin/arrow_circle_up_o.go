package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7zm1 0c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8z"/><path fill="currentColor" d="M7 6.4L5.2 8.2L3.8 6.8L8 2.6l4.2 4.2l-1.4 1.4L9 6.4V13H7V6.4z"/>`),
		g.Group(children),
	)
}
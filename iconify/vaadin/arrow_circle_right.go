package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8zm1.6-9L7.8 5.2l1.4-1.4L13.4 8l-4.2 4.2l-1.4-1.4L9.6 9H3V7h6.6z"/>`),
		g.Group(children),
	)
}
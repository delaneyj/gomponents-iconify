package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8zm4.2 10.8l-1.4 1.4L8 9.4l-2.8 2.8l-1.4-1.4L6.6 8L3.8 5.2l1.4-1.4L8 6.6l2.8-2.8l1.4 1.4L9.4 8l2.8 2.8z"/>`),
		g.Group(children),
	)
}
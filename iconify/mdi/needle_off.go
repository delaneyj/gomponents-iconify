package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeedleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.1 21.5L2.4 1.7L1.1 3L8 9.9l-3.9 3.9l2.1 2.1L3.1 19v2.8l4.5-4.5l2.1 2.1l3.9-3.9l7.2 7.2l1.3-1.2m-10.9-6.3l-1.4 1.4L7 13.8l2.5-2.5l1.4 1.4l-1.1 1.1l1.4 1.4m.7-6.5l-1.4-1.4L14 3.9L16.1 6l1.4-1.4l-1.4-1.4l1.4-1.4L21.8 6l-1.4 1.4L18.9 6l-1.4 1.4l2.1 2.1l-3.4 3.5l-2.8-2.8l.6-.7l1.4 1.4l1.4-1.4L14 6.7l-2.1 2Z"/>`),
		g.Group(children),
	)
}
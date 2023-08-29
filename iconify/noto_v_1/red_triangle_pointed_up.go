package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#db4437" d="M124.81 114.25H3.19c-1.09 0-2.09-.6-2.61-1.55c-.53-.95-.48-2.11.11-3.03L61.5 15.03c1.1-1.7 3.91-1.7 5.01 0l60.81 94.63c.59.92.63 2.08.11 3.03c-.53.96-1.53 1.56-2.62 1.56z"/>`),
		g.Group(children),
	)
}
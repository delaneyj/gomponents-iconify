package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOffTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.854 2.146a.5.5 0 1 0-.708.708l1.271 1.27a4.392 4.392 0 0 0-.156.15c-1.688 1.705-1.68 4.476.016 6.189l6.277 6.34c.26.263.682.263.942 0l2.787-2.813l3.863 3.864a.5.5 0 0 0 .708-.708l-15-15ZM16.74 10.5l-2.05 2.07l-9.438-9.437A4.317 4.317 0 0 1 9.388 4.29l.605.61l.596-.603a4.305 4.305 0 0 1 6.135.015a4.408 4.408 0 0 1 .017 6.187Z"/>`),
		g.Group(children),
	)
}
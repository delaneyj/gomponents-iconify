package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m5.6 5.4l-5.5.8L4 10l-.9 5.5L8 12.9V.4z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m10.5 5.4l5.5.8l-4 3.8l.9 5.5L8 12.9V.4z"/>`),
		g.Group(children),
	)
}
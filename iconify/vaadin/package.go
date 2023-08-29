package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0L0 2v10l8 4l8-4V2L8 0zm0 1l2.1.5l-5.9 1.9l-2.3-.8L8 1zm0 13.9l-7-3.5V3.3l3 1v3.4L5 8V4.7l3 1v9.2zm.5-10.1l-2.7-.9L12 2l2.4.6l-5.9 2.2z"/>`),
		g.Group(children),
	)
}
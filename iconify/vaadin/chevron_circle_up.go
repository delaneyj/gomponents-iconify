package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 8c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8zM4.4 10.8L3 9.4l5-5l5 5l-1.4 1.4L8 7.2l-3.6 3.6z"/>`),
		g.Group(children),
	)
}
package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Collect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#009688"><path d="M22 35h4v11h-4z"/><path d="m24 29.6l7 8.4H17z"/></g><g fill="#009688"><path d="M22 2h4v11h-4z"/><path d="M24 18.4L17 10h14z"/></g><g fill="#009688"><path d="M2 22h11v4H2z"/><path d="M18.4 24L10 31V17z"/></g><g fill="#009688"><path d="M35 22h11v4H35z"/><path d="m29.6 24l8.4-7v14z"/></g><circle cx="24" cy="24" r="3" fill="#F44336"/>`),
		g.Group(children),
	)
}
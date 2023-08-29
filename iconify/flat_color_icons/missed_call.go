package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MissedCall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#F44336"><path d="M30.3 12.9L24 19.2l-8.3-8.3l-2.8 2.8L24 24.8l9.1-9.1z"/><path d="m36 19l-9-9h9z"/></g><path fill="#009688" d="m44.5 30.8l-2.4-2.4c-8.5-8.3-28.9-7.1-36.2 0l-2.4 2.4c-.7.7-.7 1.7 0 2.4l4.8 4.7c.7.7 1.7.7 2.4 0l5.3-5.1l-.4-5.6c1.7-1.7 15.1-1.7 16.8 0l-.3 5.8l5.1 4.9c.7.7 1.7.7 2.4 0l4.8-4.7c.8-.7.8-1.8.1-2.4z"/>`),
		g.Group(children),
	)
}
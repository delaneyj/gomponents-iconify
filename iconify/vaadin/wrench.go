package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.5 13.4L7.7 5.6c.2-.5.3-1 .3-1.6c0-2.2-1.8-4-4-4c-.6 0-1.1.1-1.6.3l2.9 2.9l-2.1 2.1L.3 2.4C.1 2.9 0 3.4 0 4c0 2.2 1.8 4 4 4c.6 0 1.1-.1 1.6-.3l7.8 7.8c.6.6 1.5.6 2.1 0s.6-1.5 0-2.1zM6.8 7.6L5.4 6.2l.9-.9l1.4 1.4l-.9.9zm7.4 7.4c-.4 0-.8-.3-.8-.8c0-.4.3-.8.8-.8s.8.3.8.8s-.3.8-.8.8z"/>`),
		g.Group(children),
	)
}
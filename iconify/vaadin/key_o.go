package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 0L6 6.1C5.7 6 5.4 6 5 6c-2.8 0-5 2.2-5 5s2.3 5 5 5s5-2.2 5-5c0-.3 0-.6-.1-.9L11 9V7h2V5h2l1-1V0h-3zm-1 6h-1.7L12 4.6V6zm3-2.4l-.4.4h-1.9L15 2v1.6zm-7.7 4L8 8l2-1.7v2.3l-.8.8l-.3.4l.1.5c0 .2.1.5.1.7c0 2.2-1.8 4-4 4s-4-1.8-4-4s1.8-4 4-4c.3 0 .5 0 .8.1l.5.1l.4-.3L13.4 1H15L7.3 7.6z"/><path fill="currentColor" d="M6 11.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 6 11.5z"/>`),
		g.Group(children),
	)
}
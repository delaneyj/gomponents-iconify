package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.7 3.1c-.4-.4-1-.4-1.4 0l-2.8 2.8L9 4.5l2.8-2.8c.4-.4.4-1 0-1.4s-1-.4-1.4 0L7.6 3.1L6.2 1.7L4.8 3.1l.7.7l-1.4 1.4c-1.4 1.4-1.5 3.5-.5 5.1C1.9 11.8 1 14.1 1 16h2c0-1.3.4-3.2 2.1-4.4c1.5.8 3.4.5 4.6-.7l1.4-1.4l.7.7l1.4-1.4l-1.4-1.4l2.8-2.8c.5-.5.5-1.1.1-1.5z"/>`),
		g.Group(children),
	)
}
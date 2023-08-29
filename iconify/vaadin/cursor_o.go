package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 2.6L10.75 9H8.29l.63 1.41l1.8 4l-.91.34l-1.88-4.3l-.5-1.11l-1 .71L5 11.07V2.6zM4 0v13l3-2.14L9.26 16l2.8-1l-2.23-5H13L4 0z"/>`),
		g.Group(children),
	)
}
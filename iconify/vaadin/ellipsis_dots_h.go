package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisDotsH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 8a2 2 0 1 1-3.999.001A2 2 0 0 1 4 8zm6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8zm6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 16 8z"/>`),
		g.Group(children),
	)
}
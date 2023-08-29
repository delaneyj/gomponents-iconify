package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisDotsV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2zm0 6a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8zm0 6a2 2 0 1 1-3.999.001A2 2 0 0 1 10 14z"/>`),
		g.Group(children),
	)
}
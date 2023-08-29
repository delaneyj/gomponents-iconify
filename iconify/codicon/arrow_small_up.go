package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSmallUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 6.5L7.5 4h.7l2.5 2.5l-.7.71l-1.65-1.64v5.57h-1V5.57L5.7 7.22L5 6.5z"/>`),
		g.Group(children),
	)
}
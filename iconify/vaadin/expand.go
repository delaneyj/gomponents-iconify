package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 1h-4l1.3 1.3l-4.5 4.5l1.4 1.4l4.5-4.5L15 5zM6.8 7.8l-4.5 4.5L1 11v4h4l-1.3-1.3l4.5-4.5z"/>`),
		g.Group(children),
	)
}
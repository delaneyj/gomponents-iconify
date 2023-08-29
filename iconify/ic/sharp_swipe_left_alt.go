package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSwipeLeftAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.1 13a5 5 0 1 0 0-2H5.83l1.59-1.59L6 8l-4 4l4 4l1.41-1.41L5.83 13h4.27z"/>`),
		g.Group(children),
	)
}
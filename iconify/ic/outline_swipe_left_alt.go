package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineSwipeLeftAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.1 13a5 5 0 1 0 0-2H5.83l1.59-1.59L6 8l-4 4l4 4l1.41-1.41L5.83 13h4.27zm4.9 2c-1.66 0-3-1.34-3-3s1.34-3 3-3s3 1.34 3 3s-1.34 3-3 3z"/>`),
		g.Group(children),
	)
}
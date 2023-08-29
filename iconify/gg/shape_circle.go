package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm0 3a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
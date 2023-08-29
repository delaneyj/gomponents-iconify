package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeRhombus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 6.343L6.343 12L12 17.657L17.657 12L12 6.343ZM2.1 12l9.9 9.9l9.9-9.9L12 2.1L2.1 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
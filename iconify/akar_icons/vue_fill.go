package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VueFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.114 2H15l-3 4.9L9.429 2H0l12 21L24 2h-4.886ZM3 3.75h2.914L12 14.6l6.086-10.85H21L12 19.5L3 3.75Z"/>`),
		g.Group(children),
	)
}
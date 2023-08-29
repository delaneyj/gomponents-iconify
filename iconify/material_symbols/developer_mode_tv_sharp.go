package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperModeTvSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2H2V3h20v16h-6v2H8ZM4.8 11l2.6-2.6L6 7l-4 4l4 4l1.4-1.4L4.8 11Zm14.4 0l-2.6 2.6L18 15l4-4l-4-4l-1.4 1.4l2.6 2.6Z"/>`),
		g.Group(children),
	)
}
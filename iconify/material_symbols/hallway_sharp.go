package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HallwaySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V4h5l4-4l4 4h5v18H3Zm3-4h12l-3.75-5l-3 4L9 14l-3 4Zm4.1-14h3.8L12 2.1L10.1 4Z"/>`),
		g.Group(children),
	)
}
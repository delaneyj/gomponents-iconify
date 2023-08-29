package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventBusySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.7 18.7l-1.4-1.4l2.3-2.3l-2.3-2.3l1.4-1.4l2.3 2.3l2.3-2.3l1.4 1.4l-2.3 2.3l2.3 2.3l-1.4 1.4l-2.3-2.3l-2.3 2.3ZM3 22V4h3V2h2v2h8V2h2v2h3v18H3Zm2-2h14V10H5v10Z"/>`),
		g.Group(children),
	)
}
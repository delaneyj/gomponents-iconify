package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 17.5l-4-4v1.675l-2-2V6H8.825l-2-2H18v6.5l4-4v11Zm-9.55-7.875ZM9.55 12.4Zm11 10.95L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4ZM4 4l2 2H4v12h12v-2l2 2v2H2V4h2Z"/>`),
		g.Group(children),
	)
}
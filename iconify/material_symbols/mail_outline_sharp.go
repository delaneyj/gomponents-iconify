package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm10-7L4 8v10h16V8l-8 5Zm0-2l8-5H4l8 5ZM4 8V6v2Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabCloseRightOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.15 16.25L12 13.4l2.85 2.85l1.4-1.4L13.4 12l2.85-2.85l-1.4-1.4L12 10.6L9.15 7.75l-1.4 1.4L10.6 12l-2.85 2.85l1.4 1.4ZM5 19V5v14Zm-2 2V3h18v10.35q-.475-.175-.975-.262T19 13V5H5v14h8q0 .525.088 1.025t.262.975H3Zm16 2l-1.4-1.4l1.575-1.6H15v-2h4.175L17.6 16.4L19 15l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}
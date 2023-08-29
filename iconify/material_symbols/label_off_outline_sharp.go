package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.45 15.6L17 14.15L18.55 12L15 7H9.85l-2-2h8.2L21 12l-2.55 3.6Zm1.35 7L16.2 19H3V5.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM9.575 12.4Zm3.85-1.825ZM14.2 17L5 7.8V17h9.2Z"/>`),
		g.Group(children),
	)
}
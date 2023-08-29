package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 19.175L6.85 4H9.7l2 2H22v13.175ZM20.475 23.3l-3.3-3.3H2V4h2l2 2H3.2L.7 3.5l1.4-1.4l19.8 19.8l-1.425 1.4Z"/>`),
		g.Group(children),
	)
}
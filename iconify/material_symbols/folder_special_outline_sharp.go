package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSpecialOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.6 16.7l2.3-1.75l2.3 1.75l-.85-2.85l2.3-1.85H15.8l-.9-2.8L14 12h-2.85l2.3 1.85l-.85 2.85ZM2 20V4h8l2 2h10v14H2Zm2-2h16V8h-8.825l-2-2H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
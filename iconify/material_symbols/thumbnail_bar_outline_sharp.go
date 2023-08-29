package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbnailBarOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm8-2h10V6H10v12Zm-2 0V6H4v12h4Zm-4 0V6v12Zm4 0h2h-2ZM8 6h2h-2Z"/>`),
		g.Group(children),
	)
}
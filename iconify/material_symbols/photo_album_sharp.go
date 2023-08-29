package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoAlbumSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v20H4Zm3-4h10l-3.375-4.5L11 17l-1.625-2.175L7 18Zm4-7l2.5-1.5L16 11V4h-5v7Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeCommentOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V2h20v20l-4-4H2Zm2-2h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}
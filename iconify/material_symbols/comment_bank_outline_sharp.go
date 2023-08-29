package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentBankOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 11V4H4v12h16V4h-2.5v7L15 9.5L12.5 11ZM2 22V2h20v16H6l-4 4Zm2-6V4v12Z"/>`),
		g.Group(children),
	)
}
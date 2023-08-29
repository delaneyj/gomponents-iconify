package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentBankSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 11L15 9.5l2.5 1.5V4h-5v7ZM2 22V2h20v16H6l-4 4Z"/>`),
		g.Group(children),
	)
}
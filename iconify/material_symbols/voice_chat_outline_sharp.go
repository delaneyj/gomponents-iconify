package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceChatOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 11h1.5V9H6v2Zm2.5 2H10V7H8.5v6Zm2.75 2h1.5V5h-1.5v10ZM14 13h1.5V7H14v6Zm2.5-2H18V9h-1.5v2ZM2 22V2h20v16H6l-4 4Zm2-6h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}
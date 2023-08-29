package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoChatOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14h8v-3l2 2V7l-2 2V6H7v8Zm-5 8V2h20v16H6l-4 4Zm2-6h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}
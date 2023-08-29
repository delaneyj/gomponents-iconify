package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitscreenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 9h12V4H6v5Zm-2 2V2h16v9H4Zm2 9h12v-5H6v5Zm-2 2v-9h16v9H4ZM6 9V4v5Zm0 11v-5v5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockToBottomOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h14v-3H5v3Zm0-5h14V5H5v9Zm0 2v3v-3Zm-2 5V3h18v18H3Z"/>`),
		g.Group(children),
	)
}
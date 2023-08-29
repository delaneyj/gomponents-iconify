package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitscreenTopOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11V3h18v8H3Zm0 10v-8h18v8H3Zm2-2h14v-4H5v4Zm0-4v4v-4Z"/>`),
		g.Group(children),
	)
}
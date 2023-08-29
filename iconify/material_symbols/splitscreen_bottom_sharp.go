package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitscreenBottomSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-8h18v8H3Zm0-10V3h18v8H3Zm16-6H5v4h14V5Zm0 4H5V5h14v4Z"/>`),
		g.Group(children),
	)
}
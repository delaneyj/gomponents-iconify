package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IframeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 13.5h5v-2h-5v2ZM10 15v-5h8v5h-8Zm-8 5V4h20v16H2Zm2-2h16V8H4v10Z"/>`),
		g.Group(children),
	)
}
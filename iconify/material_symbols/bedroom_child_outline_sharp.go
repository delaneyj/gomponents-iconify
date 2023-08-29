package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedroomChildOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h1.5v-1.5h9V17H18v-6.35h-1.5V7h-9v3.65H6V17Zm3-6.5v-2h6v2H9ZM7.5 14v-2h9v2h-9ZM2 22V2h20v20H2Zm2-2h16V4H4v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}
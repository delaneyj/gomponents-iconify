package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeaturedVideoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 14h9V7H5v7Zm-3 6V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
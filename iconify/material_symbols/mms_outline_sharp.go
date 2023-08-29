package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MmsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h12l-3.75-5l-3 4L9 10l-3 4Zm-4 8V2h20v16H6l-4 4Zm2-6h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}
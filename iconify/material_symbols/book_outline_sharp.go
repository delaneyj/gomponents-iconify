package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v20H4Zm2-2h12V4h-2v7l-2.5-1.5L11 11V4H6v16Zm0 0V4v16Zm5-9l2.5-1.5L16 11l-2.5-1.5L11 11Z"/>`),
		g.Group(children),
	)
}
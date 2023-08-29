package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V6h20v12H2Zm2-2h16V8h-3v4h-2V8h-2v4h-2V8H9v4H7V8H4v8Zm3-4h2h-2Zm4 0h2h-2Zm4 0h2h-2Zm-3 0Z"/>`),
		g.Group(children),
	)
}
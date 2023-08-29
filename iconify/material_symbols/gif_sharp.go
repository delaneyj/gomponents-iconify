package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 15V9H13v6h-1.5ZM5 15V9h5v1.5H6.5v3h2V12H10v3H5Zm9.5 0V9H19v1.5h-3v1h2V13h-2v2h-1.5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IframeOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.45 23.3l-3.3-3.3H2V4h2v2.85L.65 3.5l1.425-1.425l19.8 19.8L20.45 23.3ZM4 18h11.15l-10-10H4v10Zm16-.85V8h-9.15l-4-4H22v15.15l-2-2Zm-2-2l-1.5-1.5V11.5h-2.15l-1.5-1.5H18v5.15Z"/>`),
		g.Group(children),
	)
}
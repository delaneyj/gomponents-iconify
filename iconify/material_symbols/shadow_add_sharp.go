package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShadowAddSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 14v-3h-3V9h3V6h2v3h3v2h-3v3h-2ZM2 22V6h4V2h16v16h-4v4H2Zm6-6h12V4H8v12Z"/>`),
		g.Group(children),
	)
}
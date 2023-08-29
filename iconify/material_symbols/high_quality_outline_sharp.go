package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighQualityOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.75 16.5h1.5V15H18V9h-5v6h1.75v1.5ZM6 15h1.5v-2h2v2H11V9H9.5v2.5h-2V9H6v6Zm8.5-1.5v-3h2v3h-2ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
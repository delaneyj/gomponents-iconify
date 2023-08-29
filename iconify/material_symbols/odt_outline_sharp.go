package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OdtOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15h4V9H5v6Zm1.5-1.5v-3h1v3h-1ZM10 15h3.25l.75-.75v-4.5L13.25 9H10v6Zm1.5-1.5v-3h1v3h-1Zm4.75 1.5h1.5v-4.5H19V9h-4v1.5h1.25V15ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}
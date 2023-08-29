package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFromQueueOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12h8v-2H8v2Zm0 9v-2H2V3h20v16h-6v2H8Zm-4-4h16V5H4v12Zm0 0V5v12Z"/>`),
		g.Group(children),
	)
}
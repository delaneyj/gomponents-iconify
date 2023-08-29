package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAwesomeMosaicOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21H3V3h8v18Zm-2-2V5H5v14h4Zm4-8V3h8v8h-8Zm2-2h4V5h-4v4Zm-2 12v-8h8v8h-8Zm2-2h4v-4h-4v4Zm-6-7Zm6-3Zm0 6Z"/>`),
		g.Group(children),
	)
}
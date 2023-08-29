package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletAndroidSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V1h18v22H3Zm7-3h4v-1h-4v1Zm-5-4h14V6H5v10Z"/>`),
		g.Group(children),
	)
}
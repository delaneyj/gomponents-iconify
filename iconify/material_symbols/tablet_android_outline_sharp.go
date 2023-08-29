package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletAndroidOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V1h18v22H3Zm2-5v3h14v-3H5Zm5 2h4v-1h-4v1Zm-5-4h14V6H5v10ZM5 4h14V3H5v1Zm0 0V3v1Zm0 14v3v-3Z"/>`),
		g.Group(children),
	)
}
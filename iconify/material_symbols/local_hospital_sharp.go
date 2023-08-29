package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalHospitalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 17h3v-3.5H17v-3h-3.5V7h-3v3.5H7v3h3.5V17ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}
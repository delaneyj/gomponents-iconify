package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartnerReportsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13V3h2v10h-2Zm-8 8v-5h2v3h14v-3h2v5H3Zm8-4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}
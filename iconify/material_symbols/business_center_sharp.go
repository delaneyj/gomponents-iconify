package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessCenterSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-6h7v2h6v-2h7v6H2Zm9-6v-2h2v2h-2Zm-9-2V6h6V2h8v4h6v7h-7v-2H9v2H2Zm8-7h4V4h-4v2Z"/>`),
		g.Group(children),
	)
}
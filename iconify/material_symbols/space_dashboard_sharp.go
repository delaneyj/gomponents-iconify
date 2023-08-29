package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceDashboardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h8v18H3Zm10-11V3h8v7h-8Zm8 11h-8v-9h8v9Z"/>`),
		g.Group(children),
	)
}
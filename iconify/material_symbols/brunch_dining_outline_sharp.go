package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrunchDiningOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-2h14v2H2Zm0-4v-2h5v-2h4v2h5v2H2Zm16 4v-6.1q-.9-1.025-1.45-2.025T16 11.45V2h6v9.45q0 1.425-.537 2.438T20 15.9V20h2v2h-4Zm0-14h2V4h-2v4Zm1 6.05q.45-.525.725-1.2t.275-1.4V10h-2v1.45q0 .725.25 1.4t.75 1.2Zm0 0Z"/>`),
		g.Group(children),
	)
}
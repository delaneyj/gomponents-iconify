package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectSmallSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V11h10v10H3ZM3 9V7h2v2H3Zm0-4V3h2v2H3Zm1 14h8l-2.6-3.5L7.5 18l-1.4-1.85L4 19ZM7 5V3h2v2H7Zm4 0V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm0-4V3h2v2h-2Z"/>`),
		g.Group(children),
	)
}
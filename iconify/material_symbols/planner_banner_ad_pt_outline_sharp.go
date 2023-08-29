package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlannerBannerAdPtOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v20H4Zm2-2h12V4H6v16Zm1-2h10l-3.45-4.5l-2.3 3l-1.55-2L7 18Zm-1 2V4v16Z"/>`),
		g.Group(children),
	)
}
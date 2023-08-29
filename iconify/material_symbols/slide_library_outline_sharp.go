package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideLibraryOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V6h7.975v2H4v12h16v-3.975h2V22H2Zm2-2V8v8.025v-.675V20Zm13-5.975l-1.55-3.475L11.975 9l3.475-1.55L17 4l1.55 3.45L22 9l-3.45 1.55L17 14.025Z"/>`),
		g.Group(children),
	)
}
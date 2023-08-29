package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LatticePattern(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 8h4m4 8h4M16 6v4m-8 4v4M22 8h4m4 8h4M32 6v4m-8 4v4M38 8h4m-2 6v4M6 24h4m4 8h4m-2-10v4m-8 4v4m14-10h4m4 8h4m-2-10v4m-8 4v4m14-10h4m-2 6v4M6 40h4m6-2v4m6-2h4m6-2v4m6-2h4"/>`),
		g.Group(children),
	)
}
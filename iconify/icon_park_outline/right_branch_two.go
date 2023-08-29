package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightBranchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 24h10M44 8c-3.284.464-8.955 2.783-14.776 8.348C23.902 21.436 17.433 23.768 14 24m0 0c3.284.464 8.955 2.783 14.776 8.348C34.098 37.436 40.567 39.768 44 40"/>`),
		g.Group(children),
	)
}
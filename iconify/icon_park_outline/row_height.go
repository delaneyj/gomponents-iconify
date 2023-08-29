package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m16 35l-6 6l-6-6m12-22l-6-6l-6 6m6-6v34M44 9H22m14 10H22m22 10H22m14 10H22"/>`),
		g.Group(children),
	)
}
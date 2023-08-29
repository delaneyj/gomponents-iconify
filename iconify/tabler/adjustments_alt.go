package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustmentsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h4v4H4zm2-4v4m0 4v8m4-6h4v4h-4zm2-10v10m0 4v2m4-15h4v4h-4zm2-1v1m0 4v11"/>`),
		g.Group(children),
	)
}
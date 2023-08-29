package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataThresholdingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.725 16L5 17.725V19h.85l3-3H6.725Zm3.95 0l-3 3H9.8l3-3h-2.125Zm3.725 0l-3 3h2.125l3-3H14.4Zm3.75 0l-3 3h2.125L19 17.275V16h-.85Zm-1.8-9.5l-3.675 3.675l-2-2L6.25 12.6l1.4 1.4l3.025-3l2 2l5.075-5.1l-1.4-1.4ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}
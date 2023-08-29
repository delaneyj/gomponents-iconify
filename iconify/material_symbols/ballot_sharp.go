package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallotSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10h5V8h-5v2Zm0 6h5v-2h-5v2Zm-5-5h4V7H7v4Zm0 6h4v-4H7v4Zm-4 4V3h18v18H3Z"/>`),
		g.Group(children),
	)
}
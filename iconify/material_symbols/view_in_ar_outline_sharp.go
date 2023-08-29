package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewInArOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 20.05l-7-4.025v-8.05l7-4.025l7 4.025v8.05l-7 4.025Zm-1-2.875v-4.6L7 10.25v4.625l4 2.3Zm2 0l4-2.3V10.25l-4 2.325v4.6ZM2 7V2h5v2H4v3H2Zm5 15H2v-5h2v3h3v2Zm10 0v-2h3v-3h2v5h-5Zm3-15V4h-3V2h5v5h-2Zm-8 3.85l3.95-2.325L12 6.25L8.05 8.525L12 10.85Zm0 1.125Zm0-1.125Zm1 1.725Zm-2 0Z"/>`),
		g.Group(children),
	)
}
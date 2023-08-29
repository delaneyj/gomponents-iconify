package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewInArSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 20.05l-7-4.025v-8.05l7-4.025l7 4.025v8.05l-7 4.025ZM2 7V2h5v2H4v3H2Zm5 15H2v-5h2v3h3v2Zm10 0v-2h3v-3h2v5h-5Zm3-15V4h-3V2h5v5h-2ZM8.05 8.525l-1.05.6v1.125l4 2.325v4.6l1 .575l1-.575v-4.6l4-2.325V9.125l-1.05-.6L12 10.85L8.05 8.525Z"/>`),
		g.Group(children),
	)
}
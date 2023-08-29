package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RequestQuoteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19h2v-1h2v-5h-4v-1h4v-2h-2V9h-2v1H9v5h4v1H9v2h2v1Zm-7 3V2h11l5 5v15H4Zm2-2h12V8h-4V4H6v16ZM6 4v4v-4v16V4Z"/>`),
		g.Group(children),
	)
}
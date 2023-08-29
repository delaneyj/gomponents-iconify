package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 19c0-2.79 1.91-5.13 4.5-5.8L12 2L1 21h13.35c-.22-.63-.35-1.3-.35-2m-1-1h-2v-2h2v2m0-4h-2v-4h2v4m11 4v2h-8v-2h8Z"/>`),
		g.Group(children),
	)
}
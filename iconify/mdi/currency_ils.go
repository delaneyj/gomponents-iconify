package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyIls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16a5 5 0 0 1-5 5H8V9h2v10h7a3 3 0 0 0 3-3V3h2v13m-6-8v7h-2V8a3 3 0 0 0-3-3H4v16H2V3h9a5 5 0 0 1 5 5Z"/>`),
		g.Group(children),
	)
}
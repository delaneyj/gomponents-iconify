package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRuble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 4h-9v11H7v2h3v3H7v2h3v6h2v-6h9v-2h-9v-3h7a6.007 6.007 0 0 0 6-6v-1a6.007 6.007 0 0 0-6-6Zm4 7a4.005 4.005 0 0 1-4 4h-7V6h7a4.005 4.005 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}
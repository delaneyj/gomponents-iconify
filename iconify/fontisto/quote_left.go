package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M31.2 0H24l-4.8 9.6V24h14.4V9.6h-7.2zM12 0H4.8L0 9.6V24h14.4V9.6H7.2z"/>`),
		g.Group(children),
	)
}
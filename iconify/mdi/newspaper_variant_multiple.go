package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperVariantMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 7v12h15v2H4c-2 0-2-2-2-2V7h2m17.3-4H7.7C6.76 3 6 3.7 6 4.55v10.9c0 .86.76 1.55 1.7 1.55h13.6c.94 0 1.7-.69 1.7-1.55V4.55C23 3.7 22.24 3 21.3 3M8 5h5v6H8V5m13 10H8v-2h13v2m0-4h-6V9h6v2m0-4h-6V5h6v2Z"/>`),
		g.Group(children),
	)
}
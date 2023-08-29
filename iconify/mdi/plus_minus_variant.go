package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinusVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h3V4h2v3h3v2H8v3H6V9H3V7m10 8h8v2h-8v-2m3.04-12h2.31L7.96 21H5.65L16.04 3Z"/>`),
		g.Group(children),
	)
}
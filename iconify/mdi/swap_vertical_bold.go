package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVerticalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 8h-3v6H6V8H3l5.5-6L14 8m1.5 14l5.5-6h-3v-6h-5v6h-3l5.5 6Z"/>`),
		g.Group(children),
	)
}
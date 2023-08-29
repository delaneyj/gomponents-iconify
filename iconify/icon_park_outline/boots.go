package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 4h16l-4 32l1.236.618A5 5 0 0 1 35 41.09V44H10v-2l13-6l-4-32Zm1 8h14"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSubset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 28v7h8m0-7h24v14H18V28ZM6 13.5v-1M6 20v-1M6 7V6m26 7.5v-1m0 7.5v-1m0-12V6m0 14h-1M7 20H6M7 6H6m7 0h-1m7.5 0h-1m1 14h-1M26 6h-1M13 20h-1m14 0h-1m7-14h-1"/>`),
		g.Group(children),
	)
}
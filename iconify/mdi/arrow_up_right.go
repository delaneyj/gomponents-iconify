package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.5 9.5l-1.41 1.42L17 7.83v5.67a6.5 6.5 0 0 1-6.5 6.5H4v-2h6.5c2.5 0 4.5-2 4.5-4.5V7.83l-3.09 3.08L10.5 9.5L16 4l5.5 5.5Z"/>`),
		g.Group(children),
	)
}
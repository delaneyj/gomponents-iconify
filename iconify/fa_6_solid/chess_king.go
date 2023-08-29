package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessKing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M224 0c17.7 0 32 14.3 32 32v16h16c17.7 0 32 14.3 32 32s-14.3 32-32 32h-16v48h152c22.1 0 40 17.9 40 40c0 5.3-1 10.5-3.1 15.4L368 400H80L3.1 215.4C1 210.5 0 205.3 0 200c0-22.1 17.9-40 40-40h152v-48h-16c-17.7 0-32-14.3-32-32s14.3-32 32-32h16V32c0-17.7 14.3-32 32-32zM38.6 473.4L80 432h288l41.4 41.4c4.2 4.2 6.6 10 6.6 16c0 12.5-10.1 22.6-22.6 22.6H54.6C42.1 512 32 501.9 32 489.4c0-6 2.4-11.8 6.6-16z"/>`),
		g.Group(children),
	)
}
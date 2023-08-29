package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 20H12v24h24V20ZM26 36h10m-10-8h10M8 20h32m-28-6h8.4V7.6C20.4 6.398 21.6 4 24 4c2.4 0 3.6 2.398 3.6 3.6V14H36"/>`),
		g.Group(children),
	)
}
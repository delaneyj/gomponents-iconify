package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryArrowUpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.54 22H7.33C6.6 22 6 21.4 6 20.67V5.33C6 4.6 6.6 4 7.33 4H9V2h6v2h1.67C17.4 4 18 4.6 18 5.33V12c-.7 0-1.37.13-2 .35V6H8v14h4.35c.26.75.65 1.42 1.19 2m7.4-4.5l-3-3l-3 3h2v4h2v-4h2"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 4v17"/><path stroke-linecap="round" d="M24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20c0-4.626-1.57-8.885-4.207-12.273"/><path stroke-linecap="round" d="M24 13c-6.075 0-11 4.925-11 11s4.925 11 11 11s11-4.925 11-11c0-2.544-.864-4.887-2.314-6.75"/><circle cx="24" cy="24" r="3"/></g>`),
		g.Group(children),
	)
}
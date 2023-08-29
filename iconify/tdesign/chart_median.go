package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartMedian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm12 2.404c-1.075 0-2 .909-2 2.096v2.998h.004v2.004H14v3.07C14 16.887 12.33 19 10 19s-4-2.112-4-4.429v-1h2v1C8 16.042 9.017 17 10 17c.983 0 2-.958 2-2.429V6.5c0-2.233 1.762-4.096 4-4.096s4 1.863 4 4.096v1h-2v-1c0-1.187-.925-2.096-2-2.096ZM6 9.496h2.004V11.5H6V9.496Zm2.996.002H11v2.004H8.996V9.498Zm6 0H17v2.004h-2.004V9.498Zm3.004 0h2.004v2.004H18V9.498Z"/>`),
		g.Group(children),
	)
}
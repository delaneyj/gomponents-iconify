package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelAdjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m11.114 18l3.763-14.044l27.046 7.247L40.102 18M7.898 30l-1.82 6.797l27.045 7.247L36.886 30"/><path stroke-dasharray="2 6" d="M4 24h40"/></g>`),
		g.Group(children),
	)
}
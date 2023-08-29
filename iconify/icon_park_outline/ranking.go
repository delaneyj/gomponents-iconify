package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ranking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M17 18H4v24h13V18Z"/><path d="M30 6H17v36h13V6Z"/><path stroke-linecap="round" d="M43 26H30v16h13V26Z"/></g>`),
		g.Group(children),
	)
}
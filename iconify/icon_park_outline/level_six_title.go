package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelSixTitle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M6 8v32M24 8v32M7 24h16"/><path d="M36.5 40a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M41.596 24.74C40.778 22.545 38.804 21 36.5 21c-3.038 0-5.5 2.686-5.5 6v7"/></g>`),
		g.Group(children),
	)
}
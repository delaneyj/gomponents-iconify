package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m17 14l27 10v20H17V14Z" clip-rule="evenodd"/><path d="M17 14L4 24v20h13m18 0V32l-9-3v15m18 0H17"/></g>`),
		g.Group(children),
	)
}
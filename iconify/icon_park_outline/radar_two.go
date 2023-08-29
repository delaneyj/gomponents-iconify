package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 20V4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20"/><path d="M24 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}
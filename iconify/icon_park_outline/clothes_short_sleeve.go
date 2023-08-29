package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesShortSleeve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 4H4v8h6v32h28V12h6V4ZM10 32h28m-28-8h28"/><path d="M30 4a6 6 0 0 1-12 0"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortSkirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m12 9l5-5h14l5 5l7 26s-3 9-19 9s-19-9-19-9l7-26Zm1 33l4-16"/><path d="M5 35s3 9 19 9"/></g>`),
		g.Group(children),
	)
}
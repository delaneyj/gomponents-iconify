package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func STurnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M30 24H13c-4 0-7-3-7-7s3-7 7-7h19M8 38h27c4 0 7-3 7-7s-3-7-7-7h-5"/><path stroke-linecap="round" stroke-linejoin="round" d="m13 43l-5-5l5-5"/><circle cx="37.176" cy="10" r="5" transform="rotate(-180 37.176 10)"/></g>`),
		g.Group(children),
	)
}
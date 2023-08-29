package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CycleArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M17 16h-4a9 9 0 1 0 6.345 15.383C20.985 29.753 24 25 24 25s3.006-4.732 4.632-6.36A9 9 0 1 1 35 34h-4"/><path d="m35 30l-4 4l4 4M13 12l4 4l-4 4"/></g>`),
		g.Group(children),
	)
}
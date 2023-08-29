package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecyclingPool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m20 16l4 4l4-4m4 13l4 4l4-4M8 29l4 4l4-4m8-9V4"/><path d="M36 32.867v-13.23a4 4 0 0 1 4-4h4m-32 17.23v-13.23a4 4 0 0 0-4-4H4M4 36v4a4 4 0 0 0 4 4h32a4 4 0 0 0 4-4v-4"/></g>`),
		g.Group(children),
	)
}
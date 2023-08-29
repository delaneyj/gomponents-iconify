package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CycleMovement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m20 15l4 4l4-4"/><path d="M24 19V8a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v11M28 33l-4-4l-4 4"/><path d="M24 29v11a4 4 0 0 1-4 4H8a4 4 0 0 1-4-4V29m29-9l-4 4l4 4"/><path d="M29 24h11a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H29M15 28l4-4l-4-4"/><path d="M19 24H8a4 4 0 0 1-4-4V8a4 4 0 0 1 4-4h11"/></g>`),
		g.Group(children),
	)
}
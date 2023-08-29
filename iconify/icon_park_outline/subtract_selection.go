package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="27" height="27" x="16" y="16" rx="2"/><path d="M16 32H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h23a2 2 0 0 1 2 2v9m-3 0L16 30m22-14L16 40m27-19L23 43m20-11L33 43"/></g>`),
		g.Group(children),
	)
}
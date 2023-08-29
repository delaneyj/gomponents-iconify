package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentBottomCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linejoin="round" rx="3"/><path d="M18 32v4m6-12v12m6-8v8"/></g>`),
		g.Group(children),
	)
}
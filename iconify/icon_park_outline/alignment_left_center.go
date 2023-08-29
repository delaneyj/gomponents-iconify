package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentLeftCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" stroke-linejoin="round" rx="3"/><path d="M12 30h4m-4-6h12m-12-6h8"/></g>`),
		g.Group(children),
	)
}
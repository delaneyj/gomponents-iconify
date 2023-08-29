package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="36" x="4" y="6" rx="3"/><path d="M4 14H44"/><path d="M20 24H36"/><path d="M20 32H36"/><path d="M12 24H14"/><path d="M12 32H14"/></g>`),
		g.Group(children),
	)
}
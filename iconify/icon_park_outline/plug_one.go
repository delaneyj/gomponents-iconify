package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="24" height="24" x="12" y="12" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 36v5a3 3 0 0 1-3 3H8m12-32V4m8 8V4m-6 20h4"/></g>`),
		g.Group(children),
	)
}
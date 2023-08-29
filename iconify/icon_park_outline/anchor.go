package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M14 8h20M14 8h20M14 40h20"/><rect width="8" height="8" x="36" y="4" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="4" y="4" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="36" y="36" stroke-linejoin="round" rx="2"/><rect width="8" height="8" x="4" y="36" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" d="M40 14v20M8 14v20"/></g>`),
		g.Group(children),
	)
}
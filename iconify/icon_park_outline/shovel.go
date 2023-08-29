package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shovel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M13 4h22l-1 9l-8.5 6h-3L14 13l-1-9Zm11 15v11"/><rect width="6" height="14" x="21" y="30" rx="3"/></g>`),
		g.Group(children),
	)
}
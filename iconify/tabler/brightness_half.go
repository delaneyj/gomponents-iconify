package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 9a3 3 0 0 0 0 6V9z"/><path d="M6 6h3.5L12 3.5L14.5 6H18v3.5l2.5 2.5l-2.5 2.5V18h-3.5L12 20.5L9.5 18H6v-3.5L3.5 12L6 9.5z"/></g>`),
		g.Group(children),
	)
}
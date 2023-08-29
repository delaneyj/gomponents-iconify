package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecycleBin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="m16 18l8-13.5L32 18m6 10.5L46 42H30m-12.309-.322l-15.69.178L10 28"/><path d="M17 29h13.822"/></g>`),
		g.Group(children),
	)
}
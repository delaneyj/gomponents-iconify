package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecycleBin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M16 18L24 4.5L32 18"/><path stroke-linejoin="round" d="M38 28.5L46 42H30"/><path stroke-linejoin="round" d="M17.6914 41.6782L2.00006 41.8564L10.0001 28"/><path d="M17 29H30.8217"/></g>`),
		g.Group(children),
	)
}
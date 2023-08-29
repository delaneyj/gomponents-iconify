package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterRateTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 44c8.837 0 16-7.164 16-16C40 15 24 4 24 4S8 15 8 28c0 8.837 7.163 16 16 16Z" clip-rule="evenodd"/><path d="m24 20l-4 8h8l-4 8"/></g>`),
		g.Group(children),
	)
}
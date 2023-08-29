package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchingDone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M29.633 4h-7.64v23h22v-6.037"/><path stroke-linejoin="round" d="m29.002 13.003l4.563 4.442L45 6"/><path d="M30.005 44H17.658c-1.702 0-3.742-.568-5.11-2.387c-.925-1.23-1.543-3.03-1.543-5.613V25"/><path stroke-linejoin="round" d="m3 33l8.005-8l8.009 8"/></g>`),
		g.Group(children),
	)
}
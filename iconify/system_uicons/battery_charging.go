package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12.5 6.5l2-.001c1.105-.002 2 .893 2.001 1.997l-.001 3.001a2 2 0 0 1-2 2l-2 .003m-5 0H4.487a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2H6.5"/><path d="M10.5 9.5H13L9.4 16l.1-5.5h-3l4-6zm8-1v3"/></g>`),
		g.Group(children),
	)
}
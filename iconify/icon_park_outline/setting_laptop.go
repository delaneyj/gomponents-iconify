package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M21 9H11a3 3 0 0 0-3 3v21h32v-7"/><path stroke-linejoin="round" d="M4 33h40v2a6 6 0 0 1-6 6H10a6 6 0 0 1-6-6v-2Z"/><circle cx="35" cy="14" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M35 21v-4m0-6V7m-6.062 10.5l3.464-2m5.196-3l3.464-2m-12.124 0l3.464 2m5.196 3l3.464 2"/></g>`),
		g.Group(children),
	)
}
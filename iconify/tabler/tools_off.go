package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 12l4-4a2.828 2.828 0 1 0-4-4l-4 4m-2 2l-7 7v4h4l7-7m.5-8.5l4 4"/><path d="M12 8L7 3M5 5L3 7l5 5M7 8L5.5 9.5M16 12l5 5m-2 2l-2 2l-5-5m4 1l-1.5 1.5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
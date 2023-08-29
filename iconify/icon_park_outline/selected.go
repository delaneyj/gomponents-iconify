package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Selected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M12 4H4v8h8V4Zm32 32h-8v8h8v-8Zm-32 0H4v8h8v-8ZM44 4h-8v8h8V4Z"/><path stroke-linecap="round" d="M8 36V12m32 24V12M12 8h24M12 40h24"/></g>`),
		g.Group(children),
	)
}
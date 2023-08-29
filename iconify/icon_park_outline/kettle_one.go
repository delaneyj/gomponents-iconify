package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KettleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M13 14h17s8 4.148 8 13.8c0 9.65-8 14.2-8 14.2H13s-7-6.022-7-14c0-7.979 7-14 7-14Z"/><path d="M38 28c-13-3-19 6-32 0m25-14h9s4 4 4 16M9 6l23 2.667L31 14H13L9 6Z"/></g>`),
		g.Group(children),
	)
}
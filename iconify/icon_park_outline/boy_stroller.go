package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyStroller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 24H12l2.5 10H36l4-10Zm-28 0l-4-9H3.5"/><circle cx="20" cy="41" r="3"/><circle cx="31" cy="41" r="3"/><path d="m23 8l9 16h8l4-11c-2.333-3-7-9-10-9c-4 0-8.333 2.667-11 4Zm6-3l4 7"/></g>`),
		g.Group(children),
	)
}
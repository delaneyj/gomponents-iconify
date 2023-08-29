package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><rect width="20" height="20" x="14" y="14" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" d="M34 23L23 34"/><path stroke="#fff" d="M25 14L14 25"/><path stroke="#fff" d="M34 14L14 34"/><path stroke="#fff" d="M14 22V34H26"/><path stroke="#fff" d="M22 14H34V26"/></g>`),
		g.Group(children),
	)
}
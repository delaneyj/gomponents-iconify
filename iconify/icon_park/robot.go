package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Robot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="30" height="26" x="9" y="17" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 9L28 17"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 9L20 17"/><circle cx="34" cy="7" r="2"/><circle cx="14" cy="7" r="2"/><rect width="16" height="8" x="16" y="24" fill="#2F88FF" rx="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 24H4V34H9"/><path stroke-linecap="round" stroke-linejoin="round" d="M39 24H44V34H39"/></g>`),
		g.Group(children),
	)
}
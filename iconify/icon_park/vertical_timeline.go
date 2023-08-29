package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalTimeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M3.99999 5H44"/><path stroke-linecap="round" d="M3.99999 43H44"/><path stroke-linecap="round" d="M7.99999 36V43"/><path fill="#2F88FF" stroke-linejoin="round" d="M12 28H3.99999V36H12V28Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M28 20H20V28H28V20Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M44 12H36V20H44V12Z"/><path stroke-linecap="round" d="M40 20V43"/><path stroke-linecap="round" d="M7.99999 12V13"/><path stroke-linecap="round" d="M7.99999 20V21"/><path stroke-linecap="round" d="M23 12V13"/><path stroke-linecap="round" d="M24 28V43"/></g>`),
		g.Group(children),
	)
}
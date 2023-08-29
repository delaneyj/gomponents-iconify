package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="38" height="28" x="5" y="14" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 22H36"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 28H36"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 34H36"/><circle cx="18" cy="28" r="7" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 14V6H38V14"/></g>`),
		g.Group(children),
	)
}
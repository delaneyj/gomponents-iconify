package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirConditioning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="20" x="4" y="8" rx="2"/><rect width="24" height="8" x="12" y="20" fill="#2F88FF"/><path d="M32 14H36"/><path d="M24 34V40"/><path d="M16 36V38"/><path d="M32 36V38"/></g>`),
		g.Group(children),
	)
}
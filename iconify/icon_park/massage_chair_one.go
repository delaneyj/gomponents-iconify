package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageChairOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M15 28V15.6522C15 13 18 10 24 10C30 10 33 13 33 15.6522V28"/><path fill="#2F88FF" d="M12 34V28H36V34H12Z"/><path d="M20 4H28"/><path d="M8 16V28H40V16"/><path d="M17 43L31 43"/><path d="M24 34V43"/><path d="M24 4V10"/></g>`),
		g.Group(children),
	)
}
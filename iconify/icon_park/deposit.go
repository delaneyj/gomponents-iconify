package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deposit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 16V44H36V16"/><path d="M19 20L24 26L29 20"/><path d="M18 32H30"/><path d="M18 26H30"/><path d="M24 26V38"/><path d="M14 10L23.6154 10L34 10"/><path d="M36 18H42V4H6V18H12"/></g>`),
		g.Group(children),
	)
}
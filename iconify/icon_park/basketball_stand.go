package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketballStand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="28" x="4" y="5" rx="2"/><path d="M31 22V15H17V22"/><path d="M18 39H30"/><path d="M17 33H31"/><path d="M32 27L30 39L31 43"/><path d="M16 27L18 39L17 43"/><path d="M24 27L24 43"/><path d="M34 27L14 27"/></g>`),
		g.Group(children),
	)
}
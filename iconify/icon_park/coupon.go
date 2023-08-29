package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coupon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 19.313V9H44V19.313C41.8815 20.068 40.3636 22.1053 40.3636 24.5C40.3636 26.8947 41.8815 28.932 44 29.687V40H4V29.687C6.11853 28.932 7.63636 26.8947 7.63636 24.5C7.63636 22.1053 6.11853 20.068 4 19.313Z"/><path stroke="#fff" stroke-linecap="round" d="M19 16L24 21L29 16"/><path stroke="#fff" stroke-linecap="round" d="M18 22H30"/><path stroke="#fff" stroke-linecap="round" d="M18 28.1667H30"/><path stroke="#fff" stroke-linecap="round" d="M24 22V34"/></g>`),
		g.Group(children),
	)
}
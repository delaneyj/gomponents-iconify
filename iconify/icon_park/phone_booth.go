package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneBooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="20" height="9" x="14" y="34" fill="#2F88FF"/><rect width="20" height="6" x="14" y="4" fill="#2F88FF"/><path d="M14 10V34"/><path d="M20 10V34"/><path d="M22 16V20"/><path d="M32 26L14 26"/><path d="M34 10V34"/><path d="M4 44H44"/></g>`),
		g.Group(children),
	)
}
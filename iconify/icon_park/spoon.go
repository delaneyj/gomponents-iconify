package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="20" height="20" x="14" y="4" fill="#2F88FF" rx="10"/><path d="M24 24L24 36"/><rect width="6" height="8" x="21" y="36" fill="#2F88FF" rx="3"/></g>`),
		g.Group(children),
	)
}
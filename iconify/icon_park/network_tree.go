package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="8" height="8" x="4" y="34" fill="#2F88FF" stroke="#000"/><rect width="32" height="12" x="8" y="6" fill="#2F88FF" stroke="#000"/><path stroke="#000" d="M24 34V18"/><path stroke="#000" d="M8 34V26H40V34"/><rect width="8" height="8" x="36" y="34" fill="#2F88FF" stroke="#000"/><rect width="8" height="8" x="20" y="34" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M14 12H16"/></g>`),
		g.Group(children),
	)
}
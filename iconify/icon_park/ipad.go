package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="30" height="38" x="9" y="5" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M22 36L26 36"/></g>`),
		g.Group(children),
	)
}
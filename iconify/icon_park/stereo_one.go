package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StereoOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="30" height="38" x="9" y="5" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 18H39"/><circle cx="24" cy="30" r="6" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}
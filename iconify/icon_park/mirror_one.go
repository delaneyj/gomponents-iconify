package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="20" r="16" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 36V44"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 44H34"/></g>`),
		g.Group(children),
	)
}
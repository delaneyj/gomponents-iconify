package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListFail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 10H44"/><path d="M20 24H44"/><path d="M20 38H44"/><path d="M5 7L11 13M11 7L5 13"/><circle cx="8" cy="24" r="4" fill="#2F88FF"/><circle cx="8" cy="38" r="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}
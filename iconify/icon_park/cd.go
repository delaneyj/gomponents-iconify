package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="24" r="18"/><path stroke-linecap="round" d="M13 24C13 17.9249 17.9249 13 24 13"/><circle cx="24" cy="24" r="5" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}
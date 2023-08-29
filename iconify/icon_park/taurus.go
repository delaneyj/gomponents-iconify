package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taurus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="24" cy="31" r="9" fill="#2F88FF"/><path stroke-linecap="round" d="M44 8C44 15.732 35.0457 22 24 22C12.9543 22 4 15.732 4 8"/></g>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><circle cx="14" cy="24" r="3" fill="#fff"/><circle cx="24" cy="24" r="3" fill="#fff"/><circle cx="34" cy="24" r="3" fill="#fff"/></g>`),
		g.Group(children),
	)
}
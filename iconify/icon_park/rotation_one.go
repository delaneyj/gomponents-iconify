package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotationOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 4V40H44"/><path fill="#2F88FF" d="M28 40C28 28.9543 19.0457 20 8 20V40H28Z"/></g>`),
		g.Group(children),
	)
}
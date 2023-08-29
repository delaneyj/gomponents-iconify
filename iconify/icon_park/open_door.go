package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M20 4V44L42 38V10L20 4Z"/><path stroke="#000" stroke-linecap="round" d="M6 10H20V38H6V10Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M28 22V26"/></g>`),
		g.Group(children),
	)
}
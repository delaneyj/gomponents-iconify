package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreatWall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 9V40H44V9H36V16H28V9H20V16H12V9H4Z"/><path d="M4 24H44"/><path d="M4 32H44"/><path d="M24 24V32"/><path d="M16 32V40"/><path d="M16 16V24"/><path d="M32 32V40"/><path d="M32 16V24"/></g>`),
		g.Group(children),
	)
}
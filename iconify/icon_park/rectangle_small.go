package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-width="4" d="M36 14H12C10.8954 14 10 14.8954 10 16V32C10 33.1046 10.8954 34 12 34H36C37.1046 34 38 33.1046 38 32V16C38 14.8954 37.1046 14 36 14Z"/>`),
		g.Group(children),
	)
}
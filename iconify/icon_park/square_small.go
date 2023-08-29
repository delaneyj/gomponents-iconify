package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-width="4" d="M34 12H14C12.8954 12 12 12.8954 12 14V34C12 35.1046 12.8954 36 14 36H34C35.1046 36 36 35.1046 36 34V14C36 12.8954 35.1046 12 34 12Z"/>`),
		g.Group(children),
	)
}
package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KagiMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 42H8V24H15V42H19.9406V6H27.9792V33.0659H34.0203V16.125H40.0687V42H44"/>`),
		g.Group(children),
	)
}
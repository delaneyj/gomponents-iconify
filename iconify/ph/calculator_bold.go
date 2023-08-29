package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 20H56a20 20 0 0 0-20 20v176a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20V40a20 20 0 0 0-20-20Zm-4 192H60V44h136ZM80 76a12 12 0 0 1 12-12h72a12 12 0 0 1 0 24H92a12 12 0 0 1-12-12Zm40 52a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm48 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm-48 48a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm48 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectanglesTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M208 132H48a20.023 20.023 0 0 0-20 20v48a20.023 20.023 0 0 0 20 20h160a20.023 20.023 0 0 0 20-20v-48a20.023 20.023 0 0 0-20-20zm-4 64H52v-40h152zm4-160H48a20.023 20.023 0 0 0-20 20v48a20.023 20.023 0 0 0 20 20h160a20.023 20.023 0 0 0 20-20V56a20.023 20.023 0 0 0-20-20zm-4 64H52V60h152z" fill="currentColor"/>`),
		g.Group(children),
	)
}
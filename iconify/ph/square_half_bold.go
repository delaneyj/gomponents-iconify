package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareHalfBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 36H56a20 20 0 0 0-20 20v144a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-60 64h56v16h-56Zm0 40h56v16h-56Zm56-64h-56V60h56ZM60 60h56v136H60Zm80 136v-16h56v16Z"/>`),
		g.Group(children),
	)
}
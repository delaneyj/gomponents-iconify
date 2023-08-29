package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M108 92a20 20 0 1 1 20 20a20 20 0 0 1-20-20Zm20 36a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm88-88v176a16 16 0 0 1-16 16H56a16 16 0 0 1-16-16V40a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Zm-44 116a44 44 0 0 0-20.23-37a36 36 0 1 0-47.54 0A44 44 0 1 0 172 156Z"/>`),
		g.Group(children),
	)
}
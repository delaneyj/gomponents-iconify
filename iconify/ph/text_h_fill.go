package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-24 144a8 8 0 0 1-16 0v-40H88v40a8 8 0 0 1-16 0V80a8 8 0 0 1 16 0v40h80V80a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}
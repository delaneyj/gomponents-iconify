package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Zm-24 64a8 8 0 0 1-16 0v-8h-32v88h16a8 8 0 0 1 0 16h-48a8 8 0 0 1 0-16h16V88H88v8a8 8 0 0 1-16 0V80a8 8 0 0 1 8-8h96a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 32H48a16 16 0 0 0-16 16v160a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16ZM56 96a8 8 0 0 1 8-8h16V72a8 8 0 0 1 16 0v16h16a8 8 0 0 1 0 16H96v16a8 8 0 0 1-16 0v-16H64a8 8 0 0 1-8-8Zm24 96a8 8 0 0 1-5.66-13.66l96-96a8 8 0 0 1 11.32 11.32l-96 96A8 8 0 0 1 80 192Zm112-8h-48a8 8 0 0 1 0-16h48a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}
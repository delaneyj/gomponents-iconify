package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradientFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 192a8 8 0 0 1-8 8H32a8 8 0 0 1 0-16h40a8 8 0 0 1 8 8Zm144-8h-40a8 8 0 0 0 0 16h40a8 8 0 0 0 0-16Zm-72 0h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16ZM32 168h80a8 8 0 0 0 0-16H32a8 8 0 0 0 0 16Zm192-16h-80a8 8 0 0 0 0 16h80a8 8 0 0 0 0-16Zm0-96H32a8 8 0 0 0-8 8v24a8 8 0 0 0 8 8h192a8 8 0 0 0 8-8V64a8 8 0 0 0-8-8Zm0 56H32a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h192a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}
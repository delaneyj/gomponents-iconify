package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueueFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M32 64a8 8 0 0 1 8-8h176a8 8 0 0 1 0 16H40a8 8 0 0 1-8-8Zm104 56H40a8 8 0 0 0 0 16h96a8 8 0 0 0 0-16Zm0 64H40a8 8 0 0 0 0 16h96a8 8 0 0 0 0-16Zm108.24-30.78l-64-40A8 8 0 0 0 168 120v80a8 8 0 0 0 12.24 6.78l64-40a8 8 0 0 0 0-13.56Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 112v32a8 8 0 0 1-8 8H48v16h88a8 8 0 0 1 8 8v24a8 8 0 0 1-8 8H48v8a8 8 0 0 1-16 0V40a8 8 0 0 1 16 0v8h120a8 8 0 0 1 8 8v24a8 8 0 0 1-8 8H48v16h168a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}
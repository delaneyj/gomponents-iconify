package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlurTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 12a9 9 0 0 1 13.977-7.5H12v1h6.225a9.05 9.05 0 0 1 1.26 1.5H12v1h8.064c.238.477.434.979.584 1.5H12v1h8.876c.081.488.124.989.124 1.5h-9v1h8.945a8.963 8.963 0 0 1-.297 1.5H12v1h8.294a8.98 8.98 0 0 1-.81 1.5H12v1h6.708A9 9 0 0 1 3 12Z"/>`),
		g.Group(children),
	)
}
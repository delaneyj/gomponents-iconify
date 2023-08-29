package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.075 23.52C1.264 13.642 0 12.629 0 9c0-4.971 4.029-9 9-9s9 4.029 9 9c0 3.629-1.264 4.64-8.075 14.516a1.126 1.126 0 0 1-1.847.004l-.002-.004zM9 12.75a3.75 3.75 0 1 0 0-7.5a3.75 3.75 0 0 0 0 7.5z"/>`),
		g.Group(children),
	)
}
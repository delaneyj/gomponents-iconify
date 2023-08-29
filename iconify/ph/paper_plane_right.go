package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m223.87 114l-168-95.89a16 16 0 0 0-22.94 19.21l31 90.47a.42.42 0 0 0 0 .1a.3.3 0 0 0 0 .1l-31 90.67A16 16 0 0 0 48 240a16.14 16.14 0 0 0 7.92-2.1l167.91-96.05a16 16 0 0 0 .05-27.89ZM48 224v-.09L78.14 136H136a8 8 0 0 0 0-16H78.22L48.06 32.12L48 32l168 95.83Z"/>`),
		g.Group(children),
	)
}
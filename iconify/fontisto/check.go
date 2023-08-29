package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 10.909l4.364-4.364l8.727 8.727L28.364-.001l4.364 4.364l-19.636 19.636z"/>`),
		g.Group(children),
	)
}
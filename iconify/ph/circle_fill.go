package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 128A104 104 0 1 1 128 24a104.13 104.13 0 0 1 104 104Z"/>`),
		g.Group(children),
	)
}
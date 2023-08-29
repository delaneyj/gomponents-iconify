package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowDownRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m213.66 181.66l-48 48A8 8 0 0 1 152 224v-40H64a8 8 0 0 1-8-8V32a8 8 0 0 1 16 0v136h80v-40a8 8 0 0 1 13.66-5.66l48 48a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}
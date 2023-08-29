package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashStraightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M112 112h32v32h-32Zm112-64v160a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Zm-64 96v-32h32a8 8 0 0 0 0-16h-32V64a8 8 0 0 0-16 0v32h-32V64a8 8 0 0 0-16 0v32H64a8 8 0 0 0 0 16h32v32H64a8 8 0 0 0 0 16h32v32a8 8 0 0 0 16 0v-32h32v32a8 8 0 0 0 16 0v-32h32a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}
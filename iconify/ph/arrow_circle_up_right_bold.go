package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm44-116v48a12 12 0 0 1-24 0v-19l-43.51 43.52a12 12 0 0 1-17-17L131 108h-19a12 12 0 0 1 0-24h48a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}
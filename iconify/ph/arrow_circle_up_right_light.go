package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpRightLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm38-122v48a6 6 0 0 1-12 0v-33.51l-53.76 53.75a6 6 0 0 1-8.48-8.48L145.51 102H112a6 6 0 0 1 0-12h48a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}
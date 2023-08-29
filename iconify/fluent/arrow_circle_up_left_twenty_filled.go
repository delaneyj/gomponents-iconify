package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpLeftTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16ZM7.149 7.144A.498.498 0 0 1 7.5 7h5a.5.5 0 0 1 0 1H8.707l4.147 4.146a.5.5 0 0 1-.708.708L8 8.707V12.5a.5.5 0 0 1-1 0v-5a.498.498 0 0 1 .144-.351l.005-.005Z"/>`),
		g.Group(children),
	)
}
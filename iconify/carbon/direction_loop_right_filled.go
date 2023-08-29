package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionLoopRightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M16 26h-2v-8h2Zm5-18l-1.414 1.414L24.172 14H11a3 3 0 1 1 3-3v1h2v-1a5 5 0 1 0-5 5h13.172l-4.586 4.586L21 22l7-7Z"/><path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM16 26h-2v-8h2Zm5-4l-1.414-1.414L24.172 16H11a5 5 0 1 1 5-5v1h-2v-1a3 3 0 1 0-3 3h13.172l-4.586-4.586L21 8l7 7Z"/>`),
		g.Group(children),
	)
}
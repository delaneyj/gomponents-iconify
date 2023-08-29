package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleTapSwipeUpTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.707 2.293a1 1 0 0 0-1.414 0l-3 3a1 1 0 0 0 1.414 1.414L11 5.414V15a1 1 0 1 0 2 0V5.414l1.293 1.293a1 1 0 1 0 1.414-1.414l-3-3ZM4.5 15A7.503 7.503 0 0 1 10 7.77v2.105A5.502 5.502 0 0 0 12 20.5a5.5 5.5 0 0 0 2-10.625V7.77A7.5 7.5 0 1 1 4.5 15Zm3 0a4.5 4.5 0 0 1 2.5-4.032V13.5a2.5 2.5 0 1 0 4 0v-2.532A4.5 4.5 0 1 1 7.5 15Z"/>`),
		g.Group(children),
	)
}
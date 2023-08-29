package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleTapSwipeDownTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.53 17.78a.75.75 0 0 1-1.06 0l-2.5-2.5a.75.75 0 1 1 1.06-1.06l1.22 1.22V7.75a.75.75 0 1 1 1.5 0v7.69l1.22-1.22a.75.75 0 1 1 1.06 1.06l-2.5 2.5ZM4.25 7.5a5.752 5.752 0 0 0 4 5.479v-1.605a4.25 4.25 0 1 1 3.5 0v1.605A5.75 5.75 0 1 0 4.25 7.5Zm2.5 0a3.25 3.25 0 0 0 1.5 2.74V7.75c0-.042.001-.084.004-.125a1.785 1.785 0 0 1 3.496.125v2.49a3.25 3.25 0 1 0-5-2.74Z"/>`),
		g.Group(children),
	)
}
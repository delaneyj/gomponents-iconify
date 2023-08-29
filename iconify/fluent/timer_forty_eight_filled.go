package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.25 5a1.25 1.25 0 1 0 0 2.5h9.5a1.25 1.25 0 1 0 0-2.5h-9.5ZM7 27c0-9.389 7.611-17 17-17s17 7.611 17 17s-7.611 17-17 17S7 36.389 7 27Zm15.75-.25a1.25 1.25 0 1 0 2.5 0v-9.5a1.25 1.25 0 1 0-2.5 0v9.5Zm13.616-16.384a1.25 1.25 0 0 1 1.768 0l2.5 2.5a1.25 1.25 0 0 1-1.768 1.768l-2.5-2.5a1.25 1.25 0 0 1 0-1.768Z"/>`),
		g.Group(children),
	)
}
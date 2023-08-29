package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.75 3a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5h-6.5ZM14 9.5a.75.75 0 0 1 .75.75v6a.75.75 0 0 1-1.5 0v-6A.75.75 0 0 1 14 9.5ZM14 6C8.477 6 4 10.477 4 16s4.477 10 10 10s10-4.477 10-10S19.523 6 14 6ZM5.5 16a8.5 8.5 0 1 1 17 0a8.5 8.5 0 0 1-17 0Zm16.78-9.78a.75.75 0 1 0-1.06 1.06l1.5 1.5a.75.75 0 1 0 1.06-1.06l-1.5-1.5Z"/>`),
		g.Group(children),
	)
}
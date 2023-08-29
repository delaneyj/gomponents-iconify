package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextItalicTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3.25a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-3.235L8.592 15.5h2.658a.75.75 0 0 1 0 1.5h-7.5a.75.75 0 0 1 0-1.5h3.235L11.408 4H8.75A.75.75 0 0 1 8 3.25Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraDomeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 3.75C2 2.784 2.784 2 3.75 2h16.5a1.75 1.75 0 1 1 0 3.5H3.75A1.75 1.75 0 0 1 2 3.75ZM12 17.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm2.5-4a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM3 7h18v6a9 9 0 1 1-18 0V7Zm9 12a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Z"/>`),
		g.Group(children),
	)
}
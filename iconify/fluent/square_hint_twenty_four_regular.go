package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareHintTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3.75a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75Zm7.75.75a.75.75 0 0 1 0-1.5A3.25 3.25 0 0 1 21 6.25a.75.75 0 0 1-1.5 0a1.75 1.75 0 0 0-1.75-1.75ZM6.25 3a.75.75 0 0 1 0 1.5A1.75 1.75 0 0 0 4.5 6.25a.75.75 0 0 1-1.5 0A3.25 3.25 0 0 1 6.25 3Zm-2.5 14a.75.75 0 0 0-.75.75A3.25 3.25 0 0 0 6.25 21a.75.75 0 0 0 0-1.5a1.75 1.75 0 0 1-1.75-1.75a.75.75 0 0 0-.75-.75Zm7 2.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 0-1.5h-2.5Zm7 0a.75.75 0 0 0 0 1.5A3.25 3.25 0 0 0 21 17.75a.75.75 0 0 0-1.5 0a1.75 1.75 0 0 1-1.75 1.75Zm2.5-9.5a.75.75 0 0 1 .75.75v2.5a.75.75 0 0 1-1.5 0v-2.5a.75.75 0 0 1 .75-.75ZM3 13.25a.75.75 0 0 0 1.5 0v-2.5a.75.75 0 0 0-1.5 0v2.5Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookExclamationMarkTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4.5v15A2.5 2.5 0 0 0 6.5 22h13.25a.75.75 0 0 0 0-1.5H6.5a1 1 0 0 1-1-1h14.25a.75.75 0 0 0 .75-.75V4.5A2.5 2.5 0 0 0 18 2H6.5A2.5 2.5 0 0 0 4 4.5Zm7.5 7.25v-5a.75.75 0 0 1 1.5 0v5a.75.75 0 0 1-1.5 0Zm1.75 2.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}
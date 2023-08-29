package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamesTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 5.5a8.5 8.5 0 0 0 0 17h7a8.5 8.5 0 1 0 0-17h-7Zm7.5 11a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm1.5-2.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-14-.25a.75.75 0 0 1 .75-.75H8v-1.75a.75.75 0 0 1 1.5 0V13h1.75a.75.75 0 0 1 0 1.5H9.5v1.75a.75.75 0 1 1-1.5 0V14.5H6.25a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}
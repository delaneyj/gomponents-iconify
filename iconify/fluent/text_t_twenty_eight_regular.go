package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextTTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 4.75A.75.75 0 0 1 5.75 4h16.5a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0V5.5h-6.75v17h1.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1 0-1.5h1.5v-17H6.5v2.75a.75.75 0 0 1-1.5 0v-3.5Z"/>`),
		g.Group(children),
	)
}
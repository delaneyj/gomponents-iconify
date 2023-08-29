package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterHorizontalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 3a.75.75 0 0 1 .75.75v16.5a.75.75 0 0 1-1.5 0V3.75A.75.75 0 0 1 3.75 3Zm16.5 0a.75.75 0 0 1 .75.75v16.5a.75.75 0 0 1-1.5 0V3.75a.75.75 0 0 1 .75-.75Zm-10 2A2.25 2.25 0 0 0 8 7.25v9.5A2.25 2.25 0 0 0 10.25 19h3.5A2.25 2.25 0 0 0 16 16.75v-9.5A2.25 2.25 0 0 0 13.75 5h-3.5Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.972 12.75H2.027C2.411 17.922 6.73 22 12 22c5.27 0 9.588-4.078 9.972-9.25Zm0-1.5H2.027C2.411 6.077 6.73 2 12 2c5.27 0 9.588 4.077 9.972 9.25Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.543 2.16A1 1 0 0 1 24 3v9a1 1 0 0 1-1 1H3a1 1 0 0 1-.41-1.912l20-9a1 1 0 0 1 .953.072ZM7.66 11H22V4.547L7.66 11ZM24 25.25a.75.75 0 0 1-1.065.68l-20.5-9.5A.75.75 0 0 1 2.75 15h20.5a.75.75 0 0 1 .75.75v9.5Z"/>`),
		g.Group(children),
	)
}
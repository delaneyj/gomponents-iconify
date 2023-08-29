package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassNorthwestSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14Zm2.438-5.236a.5.5 0 0 1-.674.674l-1.922-1a3 3 0 0 1-1.279-1.281l-.989-1.909a.5.5 0 0 1 .674-.674l1.91.99a3 3 0 0 1 1.28 1.278l1 1.922Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a5 5 0 1 1 0 10A5 5 0 0 1 6 1Zm0 1a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-.5 1.5a.5.5 0 0 1 .492.41L6 4v2h1.5a.5.5 0 0 1 .09.992L7.5 7h-2a.5.5 0 0 1-.492-.41L5 6.5V4a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}
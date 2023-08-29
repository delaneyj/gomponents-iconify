package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M8.334 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0z" fill="currentColor"/><path d="M14.334 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0z" fill="currentColor"/><path d="M18.334 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
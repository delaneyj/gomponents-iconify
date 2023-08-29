package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M8.584 14a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0z" fill="currentColor"/><path d="M16.584 14a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0z" fill="currentColor"/><path d="M22.334 16.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
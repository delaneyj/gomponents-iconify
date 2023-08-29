package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M8.751 14a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0z" fill="currentColor"/><path d="M16.751 14a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0z" fill="currentColor"/><path d="M22.001 16.75a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
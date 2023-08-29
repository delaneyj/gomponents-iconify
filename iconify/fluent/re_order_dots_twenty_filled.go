package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReOrderDotsTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M6.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3z" fill="currentColor"/><path d="M6.5 11a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3z" fill="currentColor"/><path d="M8 14.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z" fill="currentColor"/><path d="M13.5 6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3z" fill="currentColor"/><path d="M15 9.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z" fill="currentColor"/><path d="M13.5 16a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
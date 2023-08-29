package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M14 4a1 1 0 0 1-1 1h-2a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1z" fill="currentColor"/><path d="M5 11a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2z" fill="currentColor"/><path d="M19 11a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0v-2z" fill="currentColor"/><path d="M13 21a1 1 0 1 0 0-2h-2a1 1 0 1 0 0 2h2z" fill="currentColor"/><path d="M7 4a1 1 0 0 0-1-1a3 3 0 0 0-3 3a1 1 0 0 0 2 0a1 1 0 0 1 1-1a1 1 0 0 0 1-1z" fill="currentColor"/><path d="M18 3a1 1 0 1 0 0 2a1 1 0 0 1 1 1a1 1 0 1 0 2 0a3 3 0 0 0-3-3z" fill="currentColor"/><path d="M7 20a1 1 0 0 1-1 1a3 3 0 0 1-3-3a1 1 0 1 1 2 0a1 1 0 0 0 1 1a1 1 0 0 1 1 1z" fill="currentColor"/><path d="M18 21a1 1 0 1 1 0-2a1 1 0 0 0 1-1a1 1 0 1 1 2 0a3 3 0 0 1-3 3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
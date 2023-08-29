package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodUnamused(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m8 4l4-1.5"/><path d="M10 10c-.5-1-2.5-1-3 0m10 0c-.5-1-2.5-1-3 0"/></g>`),
		g.Group(children),
	)
}
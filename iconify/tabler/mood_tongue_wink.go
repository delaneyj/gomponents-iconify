package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodTongueWink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18z"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m6-2h.01"/><path d="M10 14v2a2 2 0 0 0 4 0v-2m1.5 0h-7m8.5-4c-.5-1-2.5-1-3 0"/></g>`),
		g.Group(children),
	)
}
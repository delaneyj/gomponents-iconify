package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodAnnoyedTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18z"/><path d="M15 14c-2 0-3 1-3.5 2.05M10 9.25c-.5 1-2.5 1-3 0m10 0c-.5 1-2.5 1-3 0"/></g>`),
		g.Group(children),
	)
}
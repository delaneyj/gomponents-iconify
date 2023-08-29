package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodXd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18z"/><path d="M9 14h6a3 3 0 1 1-6 0zm0-6l6 3m-6 0l6-3"/></g>`),
		g.Group(children),
	)
}
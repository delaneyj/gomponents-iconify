package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodSmileBeam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18z"/><path d="M10 10c-.5-1-2.5-1-3 0m10 0c-.5-1-2.5-1-3 0m.5 5a3.5 3.5 0 0 1-5 0"/></g>`),
		g.Group(children),
	)
}
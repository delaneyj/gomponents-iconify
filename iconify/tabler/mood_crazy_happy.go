package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodCrazyHappy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m4-3.5l3 3m-3 0l3-3m4 0l3 3m-3 0l3-3"/><path d="M9.5 15a3.5 3.5 0 0 0 5 0"/></g>`),
		g.Group(children),
	)
}
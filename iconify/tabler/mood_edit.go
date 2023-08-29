package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.955 11.104a9 9 0 1 0-9.895 9.847M9 10h.01M15 10h.01"/><path d="M9.5 15c.658.672 1.56 1 2.5 1c.126 0 .251-.006.376-.018m6.044-.372a2.1 2.1 0 0 1 2.97 2.97L18 22h-3v-3l3.42-3.39z"/></g>`),
		g.Group(children),
	)
}
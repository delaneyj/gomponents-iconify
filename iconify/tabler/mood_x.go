package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.983 12.556a9 9 0 1 0-8.433 8.427M9 10h.01M15 10h.01"/><path d="M9.5 15c.658.64 1.56 1 2.5 1c.194 0 .386-.015.574-.045M21.5 21.5l-5-5m0 5l5-5"/></g>`),
		g.Group(children),
	)
}
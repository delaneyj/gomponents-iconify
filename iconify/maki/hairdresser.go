package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hairdresser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M15 3s-2-.6-3.5.5l-4.3 3c-.8-.6-2-1.3-3.2-1.7V4c0-1.1-.9-2-2-2s-2 .9-2 2v1.5c0 .5.5.5.5.5h2C4.5 6 6 7.5 6 7.5S4.5 9 2.5 9h-2S0 9 0 9.5V11c0 1.1.9 2 2 2s2-.9 2-2v-.8c1.2-.4 2.4-1.1 3.2-1.7l4.3 3c1.5 1.1 3.5.5 3.5.5L8.5 7.5L15 3zM3 5H1V4c0-.6.4-1 1-1s1 .4 1 1v1zm0 6c0 .6-.4 1-1 1s-1-.4-1-1v-1h2v1zm4.25-3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1z"/>`),
		g.Group(children),
	)
}
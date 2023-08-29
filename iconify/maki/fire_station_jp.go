package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireStationJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8.5 8.149v3.601a1 1 0 0 1-2 0V8.149a5.008 5.008 0 0 1-4-4.899a1 1 0 0 1 2 0a3 3 0 0 0 6 0a1 1 0 0 1 2 0a5.008 5.008 0 0 1-4 4.899Z"/>`),
		g.Group(children),
	)
}
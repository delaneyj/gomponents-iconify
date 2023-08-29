package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleStrokedEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 0a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zm0 1.222a4.278 4.278 0 1 0 0 8.556a4.278 4.278 0 0 0 0-8.556z" fill="currentColor"/>`),
		g.Group(children),
	)
}
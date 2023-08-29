package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayBackwards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h3v10H2V7Zm4 5l7.002-5v10L6 12Zm15.002-5L14 12l7.002 5V7Z"/>`),
		g.Group(children),
	)
}
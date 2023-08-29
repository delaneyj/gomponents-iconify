package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarStrokedEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 2.69l.59 1.47l.25.63h1.81l-1 .9l-.5.44l.18.63l.56 1.69l-1.32-.92l-.57-.41l-.57.4l-1.31.92l.55-1.68l.21-.63l-.5-.43l-1-.9h1.78l.25-.63l.59-1.48M5.5 0L4 3.79H.19l3 2.66L1.71 11L5.5 8.34L9.29 11L7.78 6.45l3-2.66H7L5.5 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}
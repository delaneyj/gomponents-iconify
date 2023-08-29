package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M327.71 130.93L184 39L32 144v336l152.29-98.93L328 473l152-105V32ZM312 421l-112-72V91l112 72Z"/>`),
		g.Group(children),
	)
}
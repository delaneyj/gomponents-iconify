package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.83 13.41l-1.41 1.41l3.13 3.13L14.5 20H20v-5.5l-2.04 2.04l-3.13-3.13M14.5 4l2.04 2.04L4 18.59L5.41 20L17.96 7.46L20 9.5V4m-9.41 5.17L5.41 4L4 5.41l5.17 5.17l1.42-1.41Z"/>`),
		g.Group(children),
	)
}